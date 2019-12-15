package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/13k/geyser"
	"github.com/iancoleman/strcase"
)

var _ Command = (*FilenamesCommand)(nil)

type FilenamesCommand struct {
	OnlyMissing bool
}

func (cmd *FilenamesCommand) Run(schemas ...*Schema) error {
	sep := strings.Repeat("-", 80)

	fmt.Println(sep)

	for _, schema := range schemas {
		fmt.Printf("%s\n%s\n", path.Join("geyser", schema.relPath), sep)

		err := schema.eachSortedInterfaceGroup(func(baseName string, group geyser.SchemaInterfacesGroup) error {
			var comment string
			var missing bool

			filename := schema.Filename(group)

			if filename == "" {
				missing = true
				filename = strcase.ToSnake(strings.TrimPrefix(baseName, "I"))
				comment = " // suggested"
			}

			if !cmd.OnlyMissing || missing {
				fmt.Printf("%q: %q,%s\n", baseName, filename, comment)
			}

			return nil
		})

		if err != nil {
			return err
		}

		fmt.Println(sep)
	}

	return nil
}
