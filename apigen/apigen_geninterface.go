package main

import (
	"fmt"
	"sort"
	"strings"

	j "github.com/dave/jennifer/jen"

	"github.com/13k/geyser/schema"
)

func (g *APIGen) GenerateInterfaceFile() (string, EGenerated, error) {
	etest, err := g.interfaceFile.Update(g.genInterfaceFile)
	return g.interfaceFile.Filename(), etest, err
}

func (g *APIGen) RemoveInterfaceFile() (string, EGenerated, error) {
	etest, err := g.interfaceFile.Remove()
	return g.interfaceFile.Filename(), etest, err
}

func (g *APIGen) genInterfaceFile() (*j.File, error) {
	f := jFile(g.pkgPath, g.pkgName)

	f.HeaderComment(commentDisclaimer)
	f.HeaderComment(fmt.Sprintf(commentfDisclaimerInterface, g.baseName))

	schemaDecl, err := g.genSchemaDecl()

	if err != nil {
		return nil, err
	}

	f.Add(schemaDecl)
	f.Add(g.genStructDef())
	f.Add(g.genStructCtor())
	f.Add(g.genStructGetter())
	f.Add(g.genMethods())

	return f, nil
}

func (g *APIGen) genSchemaDecl() (j.Code, error) {
	appIDs := g.appIDs

	if len(appIDs) == 0 {
		appIDs = []uint32{0}
	}

	siDecls := make([]j.Code, len(appIDs)+1)
	siDecls[len(siDecls)-1] = j.Line()

	for i, appID := range appIDs {
		key := schema.SchemaInterfaceKey{Name: g.baseName, AppID: appID}
		si := g.interfaces[key]
		siDecl, err := g.genSIDecl(si)

		if err != nil {
			return nil, err
		}

		siDecls[i] = j.Line().Add(siDecl)
	}

	comment := fmt.Sprintf(commentfSchemaVar, g.schemaVarName, g.baseName)

	schemaDecl := j.
		Comment(comment).
		Line().
		Var().
		Id(g.schemaVarName).
		Op("=").
		Add(jSchemaInterfacesCtorID()).
		Call(siDecls...)

	return schemaDecl, nil
}

func (g *APIGen) genSIDecl(si *schema.SchemaInterface) (j.Code, error) {
	methodDecls := make([]j.Code, len(si.Methods)+1)
	methodDecls[len(methodDecls)-1] = j.Line()

	for i, sm := range si.Methods {
		methodDecl, err := g.genSMDecl(si, sm)

		if err != nil {
			return nil, err
		}

		methodDecls[i] = j.Line().Add(methodDecl)
	}

	methodsDecl := jSchemaMethodsCtorID().Call(methodDecls...)

	code := jSchemaInterfaceAddr().Values(j.Dict{
		j.Id("Name"):         j.Lit(si.Name),
		j.Id("Methods"):      methodsDecl,
		j.Id("Undocumented"): j.Lit(si.Undocumented),
	})

	return code, nil
}

func (g *APIGen) genSMDecl(si *schema.SchemaInterface, sm *schema.SchemaMethod) (j.Code, error) {
	httpMethod := jHTTPMethod(sm.HTTPMethod)

	if httpMethod == nil {
		err := fmt.Errorf(errfUnknownHTTPMethod, sm.HTTPMethod, si.Name, sm.Name)
		return nil, err
	}

	paramDecls := make([]j.Code, len(sm.Params)+1)
	paramDecls[len(paramDecls)-1] = j.Line()

	for i, smp := range sm.Params {
		paramDecl := g.genSMPDecl(si, sm, smp)
		paramDecls[i] = j.Line().Add(paramDecl)
	}

	paramsSliceDecl := jSchemaMethodParamsCtor().Call(paramDecls...)

	code := jSchemaMethodAddr().Values(j.Dict{
		j.Id("Name"):         j.Lit(sm.Name),
		j.Id("Version"):      j.Lit(sm.Version),
		j.Id("HTTPMethod"):   httpMethod,
		j.Id("Params"):       paramsSliceDecl,
		j.Id("Undocumented"): j.Lit(sm.Undocumented),
	})

	return code, nil
}

func (g *APIGen) genSMPDecl(
	_ *schema.SchemaInterface,
	_ *schema.SchemaMethod,
	smp *schema.SchemaMethodParam,
) j.Code {
	return jSchemaMethodParamAddr().Values(j.Dict{
		j.Id("Name"):        j.Lit(smp.Name),
		j.Id("Type"):        j.Lit(smp.Type),
		j.Id("Optional"):    j.Lit(smp.Optional),
		j.Id("Description"): j.Lit(smp.Description),
	})
}

func (g *APIGen) genStructDef() j.Code {
	comments := []string{
		fmt.Sprintf(commentfStruct, g.structName, g.baseName),
	}

	if g.requiredAppID {
		appIDsComment := stringJoinUint32(g.appIDs, ", ")
		comments = append(comments, fmt.Sprintf(commentfSupportedAppIDs, appIDsComment))
	}

	if g.undoc {
		comments = append(comments, commentStructUndoc)
	}

	fields := []j.Code{
		j.Id("Client").Add(jClientPtr(g.pkgPath)),
		j.Id("Interface").Add(jSchemaInterfacePtr()),
	}

	code := j.Null()

	for i, c := range comments {
		code.Comment(c).Line()

		if i != len(comments)-1 {
			code.Comment("").Line()
		}
	}

	return code.Type().Id(g.structName).Struct(fields...)
}

func (g *APIGen) genStructCtor() j.Code {
	comments := []string{
		fmt.Sprintf(commentfStructCtor, g.structCtorName, g.structName),
	}

	params := []j.Code{
		j.Id("c").Add(jClientPtr(g.pkgPath)),
	}

	if g.requiredAppID {
		appIDsComment := stringJoinUint32(g.appIDs, ", ")
		params = append(params, j.Id("appID").Uint32())
		comments = append(comments, fmt.Sprintf(commentfSupportedAppIDs, appIDsComment))
	}

	retTypes := []j.Code{
		j.Op("*").Id(g.structName),
		j.Error(),
	}

	var appID j.Code

	if g.requiredAppID {
		appID = j.Id("appID")
	}

	getArgs := []j.Code{jSchemaInterfaceKey(g.baseName, appID)}

	body := []j.Code{
		j.List(j.Id("si"), j.Err()).Op(":=").Id(g.schemaVarName).Dot("Get").Call(getArgs...),
		j.Line(),
		jIfErrRetNilErr(),
		j.Line(),
		j.Id("s").Op(":=").Op("&").Id(g.structName).Values(j.Dict{
			j.Id("Client"):    j.Id("c"),
			j.Id("Interface"): j.Id("si"),
		}),
		j.Line(),
		j.Return(j.Id("s"), j.Nil()),
	}

	code := j.Null()

	for i, c := range comments {
		code.Comment(c).Line()

		if i != len(comments)-1 {
			code.Comment("").Line()
		}
	}

	return code.
		Func().
		Id(g.structCtorName).
		Params(params...).
		Params(retTypes...).
		Block(body...)
}

func (g *APIGen) genStructGetter() j.Code {
	comments := []string{
		fmt.Sprintf(commentfStructGetter, g.structName, g.structName),
	}

	getterReceiver := j.Id("c").Add(jClientPtr(g.pkgPath))
	getterParams := []j.Code{}
	ctorParams := []j.Code{j.Id("c")}

	if g.requiredAppID {
		appIDsComment := stringJoinUint32(g.appIDs, ", ")
		getterParams = append(getterParams, j.Id("appID").Uint32())
		ctorParams = append(ctorParams, j.Id("appID"))
		comments = append(comments, fmt.Sprintf(commentfSupportedAppIDs, appIDsComment))
	}

	retTypes := []j.Code{
		j.Op("*").Id(g.structName),
		j.Error(),
	}

	fnBody := []j.Code{
		j.Return(j.Id(g.structCtorName).Call(ctorParams...)),
	}

	code := j.Null()

	for i, c := range comments {
		code.Comment(c).Line()

		if i != len(comments)-1 {
			code.Comment("").Line()
		}
	}

	return code.
		Func().
		Parens(getterReceiver).
		Id(g.structName).
		Params(getterParams...).
		Params(retTypes...).
		Block(fnBody...)
}

func (g *APIGen) genMethods() j.Code {
	code := j.Null()

	for _, name := range g.methodsNames {
		code.Line().Add(g.genMethod(name, g.groupedMethods[name]))
	}

	return code
}

func (g *APIGen) genMethod(methodName string, group schema.SchemaMethodsGroup) j.Code {
	versions := group.Versions()
	requiredVersion := len(versions) > 1
	funcName := g.methodFuncName(methodName)
	// 	resultStructName := g.methodResultStructName(methodName)

	sort.Ints(versions)

	undoc := true

	for _, sm := range group {
		if !sm.Undocumented {
			undoc = false
			break
		}
	}

	comments := []string{
		fmt.Sprintf(commentfStructMethodHeader, funcName, methodName),
	}

	if requiredVersion {
		versionsComment := stringJoinInt(versions, ", ")
		comments = append(comments, "")
		comments = append(comments, fmt.Sprintf(commentfSupportedVersions, versionsComment))
	}

	if undoc {
		comments = append(comments, "")
		comments = append(comments, commentStructMethodUndoc)
	}

	for _, version := range versions {
		key := schema.SchemaMethodKey{Name: methodName, Version: version}
		sm := group[key]

		if len(sm.Params) == 0 {
			continue
		}

		var title string

		if requiredVersion {
			title = fmt.Sprintf(commentfParamsVersioned, sm.Version)
		} else {
			title = commentParams
		}

		comments = append(comments, "", title, "")

		for _, param := range sm.Params {
			comments = append(comments, g.paramComment(param))
		}
	}

	structReceiver := j.Id("i").Op("*").Id(g.structName)

	params := []j.Code{}

	if requiredVersion {
		params = append(params, j.Id("version").Int())
	}

	retTypes := j.List(
		jRequestPtr(),
		j.Error(),
	)

	var version j.Code

	if requiredVersion {
		version = j.Id("version")
	} else {
		version = j.Lit(versions[0])
	}

	getArgs := []j.Code{jSchemaMethodKey(methodName, version)}

	reqInst := j.Id("req").Op(":=").Add(jRequestAddr()).Values(j.Dict{
		j.Id("Client"):    j.Id("i").Dot("Client"),
		j.Id("Interface"): j.Id("i").Dot("Interface"),
		j.Id("Method"):    j.Id("sm"),
		// j.Id("Result"):    j.Op("&").Id(resultStructName).Values(),
	})

	body := []j.Code{
		j.List(j.Id("sm"), j.Err()).Op(":=").Id("i").Dot("Interface").Dot("Methods").Dot("Get").Call(getArgs...),
		j.Line(),
		jIfErrRetNilErr(),
		j.Line(),
		reqInst,
		j.Line(),
		j.Return(j.Id("req"), j.Nil()),
	}

	code := j.Null().
		Comment(strings.Join(comments, "\n")).Line().
		Func().Parens(structReceiver).Id(funcName).Params(params...).Parens(retTypes).Block(body...)

	return code
}

func (g *APIGen) paramComment(param *schema.SchemaMethodParam) string {
	var b strings.Builder

	b.WriteString("  * ")
	b.WriteString(param.Name)
	b.WriteString(" [")
	b.WriteString(param.Type)
	b.WriteRune(']')

	if !param.Optional {
		b.WriteString(" (required)")
	}

	if param.Description != "" {
		b.WriteString(": ")
		b.WriteString(param.Description)
	}

	return b.String()
}
