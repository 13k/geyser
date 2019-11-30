package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"strings"

	"github.com/13k/geyser"
	j "github.com/dave/jennifer/jen"
	"github.com/huandu/xstrings"
)

func writeFile(jf *j.File, output string) error {
	buf := bytes.Buffer{}

	if err := jf.Render(&buf); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())

	if err != nil {
		return err
	}

	f, err := os.Create(output)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(formatted)

	return err
}

type action struct {
	Interfaces *geyser.SchemaInterfaces
	OutputDir  string
}

func (a *action) Execute(cmd string) error {
	switch cmd {
	case "filename":
		return a.filename()
	case "generate":
		return a.generate()
	case "clean":
		return a.clean()
	default:
		return fmt.Errorf("Invalid command %q", cmd)
	}
}

func (a *action) gen() (*apiGen, error) {
	return newAPIGen(a.Interfaces)
}

func (a *action) filename() error {
	baseName, err := a.Interfaces.GroupName()

	if err != nil {
		return err
	}

	if filename, ok := interfaceFilenames[baseName]; ok {
		fmt.Printf("%q: %q,\n", baseName, filename)
	} else {
		filename = xstrings.ToSnakeCase(strings.TrimPrefix(baseName, "I"))
		fmt.Printf("%q: %q, // suggested\n", baseName, filename)
	}

	return nil
}

func (a *action) generate() error {
	g, err := a.gen()

	if err != nil {
		return err
	}

	log.Printf("* %s\n", g.BaseName)

	outputFile := path.Join(a.OutputDir, g.Filename)
	f, err := g.InterfaceFile()

	if err != nil {
		return err
	}

	if err = writeFile(f, outputFile); err != nil {
		return err
	}

	log.Printf("  [+] %s\n", outputFile)

	resultsFile := path.Join(a.OutputDir, g.ResultsFilename)

	if _, err = os.Stat(resultsFile); err != nil && os.IsNotExist(err) {
		rf, err := g.ResultsFile()

		if err != nil {
			return err
		}

		if err = writeFile(rf, resultsFile); err != nil {
			return err
		}

		log.Printf("  [+] %s\n", resultsFile)
	}

	return nil
}

func (a *action) clean() error {
	g, err := a.gen()

	if err != nil {
		return err
	}

	log.Printf("* %s\n", g.BaseName)

	outputFile := path.Join(a.OutputDir, g.Filename)

	if err := os.Remove(outputFile); err != nil && !os.IsNotExist(err) {
		return err
	}

	log.Printf("  [-] %s\n", outputFile)

	return nil
}
