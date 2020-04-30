package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type SandboxCommand struct {
	Directory string
	Log       logrus.FieldLogger
}

func (cmd *SandboxCommand) listFiles() ([]*GeneratedFile, bool, error) {
	var files []*GeneratedFile

	gomod := false

	err := filepath.Walk(cwd, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && i.Name() == "go.mod" {
			gomod = true
			return nil
		}

		if i.IsDir() && i.Name() == "apigen" {
			return filepath.SkipDir
		}

		if !i.IsDir() && strings.HasSuffix(p, ".go") {
			files = append(files, NewGeneratedFile(p))
		}

		return nil
	})

	return files, gomod, err
}

func (cmd *SandboxCommand) createDir() (string, error) {
	var err error

	dir := cmd.Directory

	if dir == "" {
		dir, err = ioutil.TempDir(os.TempDir(), "geyser.")
	} else {
		err = os.MkdirAll(dir, 0750)
	}

	if err != nil {
		return "", err
	}

	return dir, nil
}

func (cmd *SandboxCommand) Run() error {
	files, gomod, err := cmd.listFiles()

	if err != nil {
		return err
	}

	if !gomod {
		return fmt.Errorf("could not find go.mod file in current directory %s", cwd)
	}

	sandboxDir, err := cmd.createDir()

	if err != nil {
		return err
	}

	for _, f := range files {
		etest, err := f.Test()

		if err != nil {
			return err
		}

		if etest == EGeneratedModified {
			rp, err := filepath.Rel(cwd, f.Filename())

			if err != nil {
				return err
			}

			dest := filepath.Join(sandboxDir, rp)
			destDir := filepath.Dir(dest)

			if err := os.MkdirAll(destDir, 0750); err != nil {
				return err
			}

			if err := os.Symlink(f.Filename(), dest); err != nil {
				return err
			}
		}
	}

	fmt.Println(sandboxDir)

	return nil
}
