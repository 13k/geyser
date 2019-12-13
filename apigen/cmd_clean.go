package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/13k/geyser"
)

type CleanCommand struct {
	OutputDir string
}

func (cmd *CleanCommand) Run(schemas ...*Schema) error {
	for _, schema := range schemas {
		err := schema.eachSortedInterfaceGroup(func(groupName string, group *geyser.SchemaInterfaces) error {
			outputDir := filepath.Join(cmd.OutputDir, schema.relPath)

			g, err := NewAPIGen(
				group,
				schema.pkgName,
				outputDir,
				schema.filenames,
			)

			if err != nil {
				return err
			}

			log.Printf("* %s", groupName)

			if err := cmd.clean(g.RemoveInterfaceFile); err != nil {
				return err
			}

			if err := cmd.clean(g.RemoveResultsFile); err != nil {
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

func (cmd *CleanCommand) clean(clean cleanFunc) error {
	filename, etest, err := clean()

	if err != nil {
		return err
	}

	switch etest {
	case EGeneratedDoesNotExist:
	case EGeneratedModified:
		log.Printf("  [=] %s (modified file unchanged)\n", filename)
	case EGeneratedGenerated:
		log.Printf("  [-] %s\n", filename)
	default:
		return fmt.Errorf("unknown clean result: %d", etest)
	}

	return nil
}
