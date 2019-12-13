package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/13k/geyser"
	"github.com/iancoleman/strcase"
)

type FilenamesCommand struct {
	OnlyMissing bool
}

func (cmd *FilenamesCommand) Run(schemas ...*Schema) error {
	sep := strings.Repeat("-", 80)

	for _, schema := range schemas {
		fmt.Printf("%s\n%s\n", path.Join("geyser", schema.relPath), sep)

		err := schema.eachSortedInterfaceGroup(func(groupName string, group *geyser.SchemaInterfaces) error {
			var comment string

			filename, err := schema.Filename(group)

			if err != nil {
				return err
			}

			var missing bool

			if filename == "" {
				missing = true
				filename = strcase.ToSnake(strings.TrimPrefix(groupName, "I"))
				comment = " // suggested"
			}

			if !cmd.OnlyMissing || missing {
				fmt.Printf("%q: %q,%s\n", groupName, filename, comment)
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
