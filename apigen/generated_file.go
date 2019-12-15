package main

import (
	"bufio"
	"os"
	"regexp"

	j "github.com/dave/jennifer/jen"
)

var (
	reGeneratedCode = regexp.MustCompile(`(?m:^// Code generated .* DO NOT EDIT\.$)`)
)

type EGenerated int

const (
	EGeneratedError EGenerated = iota
	EGeneratedDoesNotExist
	EGeneratedModified
	EGeneratedGenerated
)

type GeneratorFunc func() (*j.File, error)

type GeneratedFile struct {
	filename string
	file     *os.File
}

func NewGeneratedFile(filename string) *GeneratedFile {
	return &GeneratedFile{filename: filename}
}

func (f *GeneratedFile) create() error {
	fd, err := os.Create(f.filename)

	if err != nil {
		return err
	}

	f.file = fd

	return nil
}

func (f *GeneratedFile) open() error {
	fd, err := os.Open(f.filename)

	if err != nil {
		return err
	}

	f.file = fd

	return nil
}

func (f *GeneratedFile) close() error {
	return f.file.Close()
}

func (f *GeneratedFile) isGeneratedCode() bool {
	buf := bufio.NewReader(f.file)
	return reGeneratedCode.MatchReader(buf)
}

func (f *GeneratedFile) Filename() string {
	return f.filename
}

func (f *GeneratedFile) Test() (EGenerated, error) {
	err := f.open()

	if os.IsNotExist(err) {
		return EGeneratedDoesNotExist, nil
	}

	if err != nil {
		return EGeneratedError, err
	}

	defer f.close()

	if f.isGeneratedCode() {
		return EGeneratedGenerated, nil
	}

	return EGeneratedModified, nil
}

func (f *GeneratedFile) Update(gen GeneratorFunc) (etest EGenerated, err error) {
	etest, err = f.Test()

	if err == nil && etest != EGeneratedModified {
		var jf *j.File

		jf, err = gen()

		if err != nil {
			return etest, err
		}

		err = f.WriteJenFile(jf)
	}

	return etest, err
}

func (f *GeneratedFile) Remove() (EGenerated, error) {
	etest, err := f.Test()

	if err == nil && etest == EGeneratedGenerated {
		err = os.Remove(f.filename)
	}

	return etest, err
}

func (f *GeneratedFile) WriteJenFile(jf *j.File) error {
	if err := f.create(); err != nil {
		return err
	}

	defer f.file.Close()

	return jf.Render(f.file)
}
