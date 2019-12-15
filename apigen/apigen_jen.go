package main

import (
	j "github.com/dave/jennifer/jen"
)

func (g *APIGen) jQual(id string) *j.Statement {
	return j.Qual(g.pkgPath, id)
}

func (g *APIGen) jQualOp(op, id string) *j.Statement {
	return j.Op(op).Add(g.jQual(id))
}

func (g *APIGen) jQualPtr(id string) *j.Statement {
	return g.jQualOp("*", id)
}
