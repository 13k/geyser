package main

import (
	"fmt"

	j "github.com/dave/jennifer/jen"
)

func (g *APIGen) GenerateResultsFile() (string, EGenerated, error) {
	etest, err := g.resultsFile.Update(g.genResultsFile)
	return g.resultsFile.Filename(), etest, err
}

func (g *APIGen) RemoveResultsFile() (string, EGenerated, error) {
	etest, err := g.resultsFile.Remove()
	return g.resultsFile.Filename(), etest, err
}

func (g *APIGen) genResultsFile() (*j.File, error) {
	f := jFile(g.pkgPath, g.pkgName)

	f.HeaderComment(commentDisclaimer)
	f.Add(g.genResults())

	return f, nil
}

func (g *APIGen) genResults() j.Code {
	code := j.Null()

	for _, name := range g.methodsNames {
		result := g.genResult(name)
		code.Line().Add(result)
	}

	return code
}

func (g *APIGen) genResult(methodName string) j.Code {
	resultStructName := g.methodResultStructName(methodName)
	comment := fmt.Sprintf(commentfStructResult, resultStructName, g.baseName, methodName)

	return j.
		Comment(comment).
		Line().
		Type().
		Id(resultStructName).
		Struct()
}
