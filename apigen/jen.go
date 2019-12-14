package main

import (
	"net/http"
	"strings"

	j "github.com/dave/jennifer/jen"
)

var (
	jenImportNames = map[string]string{
		pkgPathGeyser: srcPkgGeyser,
	}
)

func jenFile(pkgName string) *j.File {
	f := j.NewFile(pkgName)

	f.ImportNames(jenImportNames)

	return f
}

func jenIfErrRet(values ...j.Code) *j.Statement {
	return j.If(j.Err().Op("!=").Nil()).Block(j.Return(values...))
}

func jenIfErrRetNilErr() *j.Statement {
	return jenIfErrRet(j.Nil(), j.Err())
}

func jenHTTPMethod(method string) j.Code {
	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet:
		return j.Qual(srcPkgNetHTTP, "MethodGet")
	case http.MethodPost:
		return j.Qual(srcPkgNetHTTP, "MethodPost")
	case http.MethodPut:
		return j.Qual(srcPkgNetHTTP, "MethodPut")
	default:
		return nil
	}
}

func (g *APIGen) jenGeyserID(id string) *j.Statement {
	if g.externalPkg {
		return j.Qual(pkgPathGeyser, id)
	}

	return j.Id(id)
}

func (g *APIGen) jenGeyserTypeOp(op, id string) *j.Statement {
	return j.Op(op).Add(g.jenGeyserID(id))
}

func (g *APIGen) jenGeyserTypeAddr(id string) *j.Statement {
	return g.jenGeyserTypeOp("&", id)
}

func (g *APIGen) jenGeyserTypePtr(id string) *j.Statement {
	return g.jenGeyserTypeOp("*", id)
}

func (g *APIGen) jenSchemaInterfaceAddr() *j.Statement {
	return g.jenGeyserTypeAddr(srcSchemaInterface)
}

func (g *APIGen) jenSchemaInterfacePtr() *j.Statement {
	return g.jenGeyserTypePtr(srcSchemaInterface)
}

func (g *APIGen) jenSchemaMethodAddr() *j.Statement {
	return g.jenGeyserTypeAddr(srcSchemaMethod)
}

func (g *APIGen) jenSchemaMethodParamAddr() *j.Statement {
	return g.jenGeyserTypeAddr(srcSchemaMethodParam)
}

func (g *APIGen) jenUnqualiClientPtr() *j.Statement {
	return j.Op("*").Id(srcClient)
}

func (g *APIGen) jenRequestAddr() *j.Statement {
	return g.jenGeyserTypeAddr(srcRequest)
}

func (g *APIGen) jenRequestPtr() *j.Statement {
	return g.jenGeyserTypePtr(srcRequest)
}

func (g *APIGen) jenSchemaInterfacesCtorID() *j.Statement {
	return g.jenGeyserID(srcSchemaInterfacesCtor)
}

func (g *APIGen) jenSchemaMethodsCtorID() *j.Statement {
	return g.jenGeyserID(srcSchemaMethodsCtor)
}

func (g *APIGen) jenSchemaMethodParamsCtor() *j.Statement {
	return g.jenGeyserID(srcSchemaMethodParamsCtor)
}

func (g *APIGen) jenSchemaInterfaceKey(name string, appID j.Code) *j.Statement {
	keyValues := j.Dict{
		j.Id("Name"): j.Lit(name),
	}

	if appID != nil {
		keyValues[j.Id("AppID")] = appID
	}

	return g.jenGeyserID(srcSchemaInterfaceKey).Values(keyValues)
}

func (g *APIGen) jenSchemaMethodKey(name string, version j.Code) *j.Statement {
	return g.jenGeyserID(srcSchemaMethodKey).Values(j.Dict{
		j.Id("Name"):    j.Lit(name),
		j.Id("Version"): version,
	})
}
