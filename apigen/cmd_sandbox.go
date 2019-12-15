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

func (cmd *SandboxCommand) Run() error {
	var files []*GeneratedFile

	goMod := false

	err := filepath.Walk(cwd, func(p string, i os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !i.IsDir() && i.Name() == "go.mod" {
			goMod = true
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

	if err != nil {
		return err
	}

	if !goMod {
		return fmt.Errorf("could not find go.mod file in current directory %s", cwd)
	}

	sandboxDir := cmd.Directory

	if sandboxDir == "" {
		if sandboxDir, err = ioutil.TempDir(os.TempDir(), "geyser."); err != nil {
			return err
		}
	} else {
		if err := os.MkdirAll(sandboxDir, 0755); err != nil {
			return err
		}
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

			if err := os.MkdirAll(destDir, 0755); err != nil {
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
