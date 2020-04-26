package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/13k/geyser/schema"
)

var _ Command = (*FilenamesCommand)(nil)

type FilenamesCommand struct {
	OnlyMissing bool
}

func (cmd *FilenamesCommand) Run(schemas ...*Schema) error {
	sep := strings.Repeat("-", 80)

	fmt.Println(sep)

	for _, s := range schemas {
		fmt.Printf("%s\n%s\n", path.Join("geyser", s.relPath), sep)

		err := s.eachSortedInterfaceGroup(func(baseName string, group schema.InterfacesGroup) error {
			var comment string
			var missing bool

			filename := s.Filename(group)

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
