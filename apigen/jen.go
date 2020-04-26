package main

import (
	"net/http"
	"strings"

	j "github.com/dave/jennifer/jen"
)

var (
	jenImportNames = map[string]string{
		pkgPathGeyser:         pkgNameGeyser,
		pkgPathGeyserSchema:   pkgNameGeyserSchema,
		pkgPathGeyserSteam:    pkgNameGeyserSteam,
		pkgPathGeyserDota2:    pkgNameGeyserDota2,
		pkgPathTestifyAssert:  pkgNameTestifyAssert,
		pkgPathTestifyRequire: pkgNameTestifyRequire,
	}
)

func jFile(pkgPath, pkgName string) *j.File {
	f := j.NewFilePathName(pkgPath, pkgName)

	f.ImportNames(jenImportNames)

	return f
}

//// Go statements

func jVarErrError() *j.Statement {
	return j.Var().Err().Error()
}

func jIfErrRet(values ...j.Code) *j.Statement {
	return j.If(j.Err().Op("!=").Nil()).Block(j.Return(values...))
}

func jIfErrRetNilErr() *j.Statement {
	return jIfErrRet(j.Nil(), j.Err())
}

type jBagItem struct {
	ID   string
	Stmt j.Code
}

type jBag []jBagItem

func (bag *jBag) Declare(id string, stmt j.Code) bool {
	for _, item := range *bag {
		if item.ID == id {
			return false
		}
	}

	*bag = append(*bag, jBagItem{ID: id, Stmt: stmt})

	return true
}

func (bag jBag) Codes() []j.Code {
	if bag == nil {
		return nil
	}

	s := make([]j.Code, len(bag))

	for i, item := range bag {
		s[i] = item.Stmt
	}

	return s
}

//// package http statements

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

//// package testing statements

func jTestingTIdPtr(id string) *j.Statement {
	return j.Id(id).Op("*").Qual(pkgPathTesting, "T")
}

func jTestifyAssert(assertion, tID string, args ...j.Code) *j.Statement { //nolint:unparam
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyAssert, assertion).Call(args...)
}

func jTestifyRequire(assertion, tID string, args ...j.Code) *j.Statement { //nolint:unparam
	args = append([]j.Code{j.Id(tID)}, args...)
	return j.Qual(pkgPathTestifyRequire, assertion).Call(args...)
}

//// package geyser statements

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

func jRequestAddr() *j.Statement {
	return jGeyserTypeAddr(srcRequest)
}

func jRequestPtr() *j.Statement {
	return jGeyserTypePtr(srcRequest)
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

//// package schema statements

func jGeyserSchemaID(id string) *j.Statement {
	return j.Qual(pkgPathGeyserSchema, id)
}

func jGeyserSchemaTypeOp(op, id string) *j.Statement {
	return j.Op(op).Add(jGeyserSchemaID(id))
}

func jGeyserSchemaTypeAddr(id string) *j.Statement {
	return jGeyserSchemaTypeOp("&", id)
}

func jGeyserSchemaTypePtr(id string) *j.Statement {
	return jGeyserSchemaTypeOp("*", id)
}

func jSchemaInterfaceAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaInterface)
}

func jSchemaInterfacePtr() *j.Statement {
	return jGeyserSchemaTypePtr(srcSchemaInterface)
}

func jSchemaMethodAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaMethod)
}

func jSchemaMethodParamAddr() *j.Statement {
	return jGeyserSchemaTypeAddr(srcSchemaMethodParam)
}

func jSchemaInterfacesCtorID() *j.Statement {
	return jGeyserSchemaID(srcSchemaInterfacesCtor)
}

func jSchemaMethodsCtorID() *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodsCtor)
}

func jSchemaMethodParamsCtor() *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodParamsCtor)
}

func jSchemaInterfaceKey(name string, appID j.Code) *j.Statement {
	keyValues := j.Dict{
		j.Id("Name"): j.Lit(name),
	}

	if appID != nil {
		keyValues[j.Id("AppID")] = appID
	}

	return jGeyserSchemaID(srcSchemaInterfaceKey).Values(keyValues)
}

func jSchemaMethodKey(name string, version j.Code) *j.Statement {
	return jGeyserSchemaID(srcSchemaMethodKey).Values(j.Dict{
		j.Id("Name"):    j.Lit(name),
		j.Id("Version"): version,
	})
}

func jInterfaceMethodNotFoundErrorPtr() *j.Statement {
	return jGeyserSchemaTypePtr(srcIntefaceMethodNotFoundError)
}
