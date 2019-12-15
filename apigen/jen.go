package main

import (
	"net/http"
	"strings"

	j "github.com/dave/jennifer/jen"
)

var (
	jenImportNames = map[string]string{
		pkgPathGeyser:         srcPkgGeyser,
		pkgPathGeyserDota2:    srcPkgGeyserDota2,
		pkgPathTestifyAssert:  srcPkgTestifyAssert,
		pkgPathTestifyRequire: srcPkgTestifyRequire,
	}
)

func jFile(pkgPath, pkgName string) *j.File {
	f := j.NewFilePathName(pkgPath, pkgName)

	f.ImportNames(jenImportNames)

	return f
}

func jVarErrError() *j.Statement {
	return j.Var().Err().Error()
}

func jIfErrRet(values ...j.Code) *j.Statement {
	return j.If(j.Err().Op("!=").Nil()).Block(j.Return(values...))
}

func jIfErrRetNilErr() *j.Statement {
	return jIfErrRet(j.Nil(), j.Err())
}

func jHTTPMethod(method string) *j.Statement {
	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet:
		return j.Qual(pkgPathNetHTTP, "MethodGet")
	case http.MethodPost:
		return j.Qual(pkgPathNetHTTP, "MethodPost")
	case http.MethodPut:
		return j.Qual(pkgPathNetHTTP, "MethodPut")
	default:
		return nil
	}
}

func jTestingTIdPtr(id string) *j.Statement {
	return j.Id(id).Op("*").Qual(pkgPathTesting, "T")
}

func jTestifyAssert(assertion, tID string, args ...j.Code) *j.Statement {
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyAssert, assertion).Call(args...)
}

func jTestifyRequire(assertion, tID string, args ...j.Code) *j.Statement {
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyRequire, assertion).Call(args...)
}

func jGeyserID(id string) *j.Statement {
	return j.Qual(pkgPathGeyser, id)
}

func jGeyserTypeOp(op, id string) *j.Statement {
	return j.Op(op).Add(jGeyserID(id))
}

func jGeyserTypeAddr(id string) *j.Statement {
	return jGeyserTypeOp("&", id)
}

func jGeyserTypePtr(id string) *j.Statement {
	return jGeyserTypeOp("*", id)
}

func jSchemaInterfaceAddr() *j.Statement {
	return jGeyserTypeAddr(srcSchemaInterface)
}

func jSchemaInterfacePtr() *j.Statement {
	return jGeyserTypePtr(srcSchemaInterface)
}

func jSchemaMethodAddr() *j.Statement {
	return jGeyserTypeAddr(srcSchemaMethod)
}

func jSchemaMethodParamAddr() *j.Statement {
	return jGeyserTypeAddr(srcSchemaMethodParam)
}

func jRequestAddr() *j.Statement {
	return jGeyserTypeAddr(srcRequest)
}

func jRequestPtr() *j.Statement {
	return jGeyserTypePtr(srcRequest)
}

func jSchemaInterfacesCtorID() *j.Statement {
	return jGeyserID(srcSchemaInterfacesCtor)
}

func jSchemaMethodsCtorID() *j.Statement {
	return jGeyserID(srcSchemaMethodsCtor)
}

func jSchemaMethodParamsCtor() *j.Statement {
	return jGeyserID(srcSchemaMethodParamsCtor)
}

func jSchemaInterfaceKey(name string, appID j.Code) *j.Statement {
	keyValues := j.Dict{
		j.Id("Name"): j.Lit(name),
	}

	if appID != nil {
		keyValues[j.Id("AppID")] = appID
	}

	return jGeyserID(srcSchemaInterfaceKey).Values(keyValues)
}

func jSchemaMethodKey(name string, version j.Code) *j.Statement {
	return jGeyserID(srcSchemaMethodKey).Values(j.Dict{
		j.Id("Name"):    j.Lit(name),
		j.Id("Version"): version,
	})
}

func jInterfaceMethodNotFoundErrorPtr() *j.Statement {
	return jGeyserTypePtr(srcIntefaceMethodNotFoundError)
}

func jClientID(pkgPath string) *j.Statement {
	if pkgPath != "" {
		return j.Qual(pkgPath, srcClient)
	}

	return j.Id(srcClient)
}

func jClientOp(op, pkgPath string) *j.Statement {
	return j.Op(op).Add(jClientID(pkgPath))
}

func jClientAddr(pkgPath string) *j.Statement {
	return jClientOp("&", pkgPath)
}

func jClientPtr(pkgPath string) *j.Statement {
	return jClientOp("*", pkgPath)
}

func (g *APIGen) jQual(id string) *j.Statement {
	return j.Qual(g.pkgPath, id)
}

func (g *APIGen) jQualOp(op, id string) *j.Statement {
	return j.Op(op).Add(g.jQual(id))
}

func (g *APIGen) jQualPtr(id string) *j.Statement {
	return g.jQualOp("*", id)
}
