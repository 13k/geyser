package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/13k/geyser"
	j "github.com/dave/jennifer/jen"
	"github.com/huandu/xstrings"
)

var (
	reGeneratedCode = regexp.MustCompile(`^//.+DO NOT EDIT`)
)

func writeJenFile(jf *j.File, output string) error {
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

func isGeneratedCode(r io.Reader) bool {
	buf := bufio.NewReader(r)
	return reGeneratedCode.MatchReader(buf)
}

type eTestGenerated int

const (
	testGeneratedError eTestGenerated = iota
	testGeneratedDoesNotExist
	testGeneratedModified
	testGeneratedGenerated
)

func testGeneratedFile(filename string) (eTestGenerated, error) {
	file, err := os.Open(filename)

	if os.IsNotExist(err) {
		return testGeneratedDoesNotExist, nil
	}

	if err != nil {
		return testGeneratedError, err
	}

	defer file.Close()

	if isGeneratedCode(file) {
		return testGeneratedGenerated, nil
	}

	return testGeneratedModified, nil
}

func updateGeneratedFile(filename string, gen func() (*j.File, error)) (eTestGenerated, error) {
	rtest, err := testGeneratedFile(filename)

	if err == nil && rtest != testGeneratedModified {
		var jf *j.File

		jf, err = gen()

		if err != nil {
			return rtest, err
		}

		err = writeJenFile(jf, filename)
	}

	return rtest, err
}

func removeGeneratedFile(filename string) (eTestGenerated, error) {
	rtest, err := testGeneratedFile(filename)

	if err == nil && rtest == testGeneratedGenerated {
		err = os.Remove(filename)
	}

	return rtest, err
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

func (a *action) absPath(filename string) string {
	return filepath.Join(a.OutputDir, filename)
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

	generation := map[string]func() (*j.File, error){
		g.Filename:        g.InterfaceFile,
		g.ResultsFilename: g.ResultsFile,
	}

	log.Printf("* %s\n", g.BaseName)

	for filename, generator := range generation {
		absFilename := a.absPath(filename)
		rtest, err := updateGeneratedFile(absFilename, generator)

		if err = a.logGenerate(filename, rtest, err); err != nil {
			return err
		}
	}

	return nil
}

func (a *action) logGenerate(filename string, rtest eTestGenerated, err error) error {
	if err != nil {
		return err
	}

	switch rtest {
	case testGeneratedDoesNotExist:
		log.Printf("  [+] %s\n", filename)
	case testGeneratedModified:
		log.Printf("  [=] %s (modified file unchanged)\n", filename)
	case testGeneratedGenerated:
		log.Printf("  [#] %s (re-generated)\n", filename)
	default:
		return fmt.Errorf("unhandled testGeneratedFile result: %d", rtest)
	}

	return nil
}

func (a *action) clean() error {
	g, err := a.gen()

	if err != nil {
		return err
	}

	cleanup := []string{
		g.Filename,
		g.ResultsFilename,
	}

	log.Printf("* %s\n", g.BaseName)

	for _, filename := range cleanup {
		absFilename := a.absPath(filename)
		rtest, err := removeGeneratedFile(absFilename)

		if err = a.logClean(filename, rtest, err); err != nil {
			return err
		}
	}

	return nil
}

func (a *action) logClean(filename string, rtest eTestGenerated, err error) error {
	if err != nil {
		return err
	}

	switch rtest {
	case testGeneratedDoesNotExist:
	case testGeneratedModified:
		log.Printf("  [=] %s (modified file unchanged)\n", filename)
	case testGeneratedGenerated:
		log.Printf("  [-] %s\n", filename)
	default:
		return fmt.Errorf("unhandled testGeneratedFile result: %d", rtest)
	}

	return nil
}
