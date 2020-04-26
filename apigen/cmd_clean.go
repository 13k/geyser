package main

import (
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/13k/geyser/schema"
)

var _ Command = (*CleanCommand)(nil)

type CleanCommand struct {
	OutputDir string
	Log       logrus.FieldLogger
}

func (cmd *CleanCommand) Run(schemas ...*Schema) error {
	for _, s := range schemas {
		err := s.eachSortedInterfaceGroup(func(baseName string, group schema.SchemaInterfacesGroup) error {
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

			if err := cmd.clean(baseName, g.RemoveInterfaceFile); err != nil {
				return err
			}

			if err := cmd.clean(baseName, g.RemoveResultsFile); err != nil {
				return err
			}

			if err := cmd.clean(baseName, g.RemoveTestsFile); err != nil {
				return err
			}

			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}

type cleanFunc func() (string, EGenerated, error)

func (cmd *CleanCommand) clean(baseName string, clean cleanFunc) error {
	filename, etest, err := clean()

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
	case EGeneratedModified:
		l.Warn("manually modified file unchanged")
	case EGeneratedGenerated:
		l.Info("removed")
	default:
		return fmt.Errorf("unknown clean result: %d", etest)
	}

	return nil
}
