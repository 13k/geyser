package main

import (
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/13k/geyser/v2/schema"
)

var _ Command = (*GenerateCommand)(nil)

type GenerateCommand struct {
	OutputDir string
	Log       logrus.FieldLogger
}

func (cmd *GenerateCommand) Run(schemas ...*Schema) error {
	for _, s := range schemas {
		err := s.eachSortedInterfaceGroup(cmd.generator(s))

		if err != nil {
			return err
		}
	}

	return nil
}

func (cmd *GenerateCommand) generator(s *Schema) interfaceGroupIterator {
	return func(baseName string, group schema.InterfacesGroup) error {
		baseFilename := s.Filename(group)

		if baseFilename == "" {
			return fmt.Errorf(errfUnknownInterfaceFilename, baseName)
		}

		outputDir := filepath.Join(cmd.OutputDir, s.relPath)

		g, err := NewAPIGen(
			group,
			s.pkgPath,
			s.pkgName,
			outputDir,
			baseFilename,
		)

		if err != nil {
			return err
		}

		if err := cmd.gen(baseName, g.GenerateInterfaceFile); err != nil {
			return err
		}

		// if err := cmd.gen(baseName, g.GenerateResultsFile); err != nil {
		// 	return err
		// }

		if err := cmd.gen(baseName, g.GenerateTestsFile); err != nil {
			return err
		}

		return nil
	}
}

type genFunc func() (string, EGenerated, error)

func (cmd *GenerateCommand) gen(baseName string, gen genFunc) error {
	filename, etest, err := gen()

	if err != nil {
		return err
	}

	relpath, err := filepath.Rel(cmd.OutputDir, filename)

	if err != nil {
		panic(err)
	}

	l := cmd.Log.WithFields(logrus.Fields{
		"interface": baseName,
		"file":      relpath,
	})

	switch etest {
	case EGeneratedDoesNotExist:
		l.Info("generated")
	case EGeneratedModified:
		l.Warn("manually modified file unchanged")
	case EGeneratedGenerated:
		l.Info("re-generated")
	default:
		return fmt.Errorf("unknown generation result: %d", etest)
	}

	return nil
}
