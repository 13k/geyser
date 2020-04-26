package main

import (
	"fmt"
	"sort"

	j "github.com/dave/jennifer/jen"

	"github.com/13k/geyser/schema"
)

func (g *APIGen) GenerateTestsFile() (string, EGenerated, error) {
	etest, err := g.testsFile.Update(g.genTestsFile)
	return g.testsFile.Filename(), etest, err
}

func (g *APIGen) RemoveTestsFile() (string, EGenerated, error) {
	etest, err := g.testsFile.Remove()
	return g.testsFile.Filename(), etest, err
}

func (g *APIGen) genTestsFile() (*j.File, error) {
	f := jFile(g.pkgPath+"_test", g.testPkgName)

	f.HeaderComment(commentDisclaimer)
	f.Add(g.genTests())

	return f, nil
}

func (g *APIGen) genTests() j.Code {
	code := j.Null()

	code.Add(g.genTestCtor())
	code.Add(g.genTestMethods())

	return code
}

func (g *APIGen) genTestCtor() j.Code {
	testName := fmt.Sprintf("Test%s", g.structCtorName)

	ctorArgs := []j.Code{
		j.Id("client"),
	}

	if g.requiredAppID {
		ctorArgs = append(ctorArgs, j.Id("appID"))
	}

	testCaseBlock := []j.Code{
		j.List(j.Id("i"), j.Err()).Op(":=").Add(g.jQual(g.structCtorName)).Call(ctorArgs...),
		j.Line(),
		jTestifyRequire("NoError", "t", j.Err()),
		jTestifyRequire("NotNil", "t", j.Id("i")),
		j.Line(),
		jTestifyAssert("Same", "t", j.Id("client"), j.Id("i").Dot("Client")),
		jTestifyAssert("NotNil", "t", j.Id("i").Dot("Interface")),
	}

	if g.requiredAppID {
		forCond := j.List(j.Id("_"), j.Id("appID")).Op(":=").Range().Id("appIDs")

		testCaseBlock = []j.Code{
			j.Line(),
			j.For(forCond).Block(testCaseBlock...),
		}
	}

	body := []j.Code{
		j.Id("client").Op(":=").Add(jClientAddr(g.pkgPath)).Values(),
	}

	if g.requiredAppID {
		appIDs := make([]j.Code, len(g.appIDs))

		for i, appID := range g.appIDs {
			appIDs[i] = j.Lit(int(appID))
		}

		appIDsDecl := j.Id("appIDs").Op(":=").Index().Uint32().Values(appIDs...)
		body = append(body, appIDsDecl)
	}

	body = append(body, testCaseBlock...)

	return j.Func().Id(testName).Call(jTestingTIdPtr("t")).Block(body...)
}

func (g *APIGen) genTestMethods() j.Code {
	code := j.Null()

	appIDs := g.appIDs

	if len(appIDs) == 0 {
		appIDs = []uint32{0}
	}

	for _, methodName := range g.methodsNames {
		methodsGroup := g.groupedMethods[methodName]
		code.Line().Add(g.genTestMethod(appIDs, methodName, methodsGroup))
	}

	return code
}

func (g *APIGen) genTestMethod(appIDs []uint32, name string, group schema.SchemaMethodsGroup) j.Code {
	structFuncName := g.methodFuncName(name)
	testName := fmt.Sprintf("Test%s_%s", g.structName, structFuncName)
	resultStructName := g.methodResultStructName(name)
	versions := group.Versions()
	requiredVersion := len(versions) > 1

	sort.Ints(versions)

	body := []j.Code{
		j.Var().Id("i").Add(g.jQualPtr(g.structName)),
		jVarErrError(),
		j.Var().Id("req").Add(jRequestPtr()),
		j.Var().Id("result").Add(g.jQualPtr(resultStructName)),
		j.Var().Id("ok").Bool(),
		j.Line(),
		j.Id("client").Op(":=").Add(jClientAddr(g.pkgPath)).Values(),
	}

	for _, appID := range appIDs {
		siKey := schema.SchemaInterfaceKey{Name: g.baseName, AppID: appID}
		si := g.interfaces[siKey]

		for _, version := range versions {
			smKey := schema.SchemaMethodKey{Name: name, Version: version}
			_, err := si.Methods.Get(smKey)

			if err != nil {
				if _, ok := err.(*schema.InterfaceMethodNotFoundError); !ok {
					panic(err)
				}
			}

			ctorArgs := []j.Code{j.Id("client")}

			if appID != 0 {
				ctorArgs = append(ctorArgs, j.Lit(int(appID)))
			}

			var methodArgs []j.Code

			if requiredVersion {
				methodArgs = append(methodArgs, j.Lit(version))
			}

			body = append(
				body,
				j.Line(),
				j.List(j.Id("i"), j.Err()).Op("=").Add(g.jQual(g.structCtorName)).Call(ctorArgs...),
				j.Line(),
				jTestifyRequire("NoError", "t", j.Err()),
				jTestifyRequire("NotNil", "t", j.Id("i")),
				j.Line(),
				j.List(j.Id("req"), j.Err()).Op("=").Id("i").Dot(structFuncName).Call(methodArgs...),
				j.Line(),
			)

			if err == nil {
				body = append(
					body,
					jTestifyRequire("NoError", "t", j.Err()),
					jTestifyRequire("NotNil", "t", j.Id("req")),
					j.Line(),
					jTestifyAssert("Same", "t", j.Id("client"), j.Id("req").Dot("Client")),
					jTestifyAssert("Same", "t", j.Id("i").Dot("Interface"), j.Id("req").Dot("Interface")),
					j.Line(),
					j.If(jTestifyAssert("NotNil", "t", j.Id("req").Dot("Method"))).Block(
						jTestifyAssert("Equal", "t", j.Lit(name), j.Id("req").Dot("Method").Dot("Name")),
						jTestifyAssert("Equal", "t", j.Lit(version), j.Id("req").Dot("Method").Dot("Version")),
					),
					j.Line(),
					j.List(j.Id("result"), j.Id("ok")).Op("=").Id("req").Dot("Result").Assert(g.jQualPtr(resultStructName)),
					j.Line(),
					j.If(jTestifyAssert("Truef", "t", j.Id("ok"), j.Lit(testfInvalidResultType), j.Id("result"))).Block(
						jTestifyAssert("NotNil", "t", j.Id("result")),
					),
				)
			} else {
				body = append(
					body,
					jTestifyRequire("Error", "t", j.Err()),
					j.Line(),
					j.List(j.Id("_"), j.Id("ok")).Op("=").Err().Assert(jInterfaceMethodNotFoundErrorPtr()),
					j.Line(),
					jTestifyAssert("Truef", "t", j.Id("ok"), j.Lit(testfInvalidErrorType), j.Id("err")),
				)
			}
		}
	}

	return j.Line().Func().Id(testName).Call(jTestingTIdPtr("t")).Block(body...)
}
