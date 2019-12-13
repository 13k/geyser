package main

import (
	"fmt"
	"log"
	"path/filepath"
)

type GenerateCommand struct {
	OutputDir string
}

func (cmd *GenerateCommand) Run(schemas ...*Schema) error {
	for _, schema := range schemas {
		for groupName, group := range schema.API.Interfaces.GroupByName() {
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

			if err := cmd.gen(g.GenerateInterfaceFile); err != nil {
				return err
			}

			if err := cmd.gen(g.GenerateResultsFile); err != nil {
				return err
			}
		}
	}

	return nil
}

type genFunc func() (string, EGenerated, error)

func (cmd *GenerateCommand) gen(gen genFunc) error {
	filename, etest, err := gen()

	if err != nil {
		return err
	}

	switch etest {
	case EGeneratedDoesNotExist:
		log.Printf("  [+] %s\n", filename)
	case EGeneratedModified:
		log.Printf("  [=] %s (modified file unchanged)\n", filename)
	case EGeneratedGenerated:
		log.Printf("  [#] %s (re-generated)\n", filename)
	default:
		return fmt.Errorf("unknown generation result: %d", etest)
	}

	return nil
}
