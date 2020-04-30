package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/13k/geyser/v2/schema"
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

		err := s.eachSortedInterfaceGroup(cmd.filenamePrinter(s))

		if err != nil {
			return err
		}

		fmt.Println(sep)
	}

	return nil
}

func (cmd *FilenamesCommand) filenamePrinter(s *Schema) interfaceGroupIterator {
	return func(baseName string, group schema.InterfacesGroup) error {
		var (
			comment string
			missing bool
		)

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
	}
}
