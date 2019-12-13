package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	fmtUsage = `
Usage: apigen [options] <command>

Commands:

filenames [options]  Print a list of interfaces and respective filenames
generate             Generate files
clean                Remove all generated files


Options:

%s

If any of the cached schema option is used and the file doesn't exist, the
schema will be fetched remotely and the file will be created. If the file
already exists, the schema will be loaded from it.

`
)

var (
	apiKey          string
	steamSchemaFile string
	dotaSchemaFile  string
	outputDir       string
)

func init() {
	cwd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	flag.StringVar(&apiKey, "key", "", "use API key when fetching API schema")
	flag.StringVar(&steamSchemaFile, "steam", "", "Use file as cached Steam API schema")
	flag.StringVar(&dotaSchemaFile, "dota", "", "Use file as cached Dota2 API schema")
	flag.StringVar(&outputDir, "output", cwd, "Output directory")
	flag.StringVar(&outputDir, "o", cwd, "Output directory")

	flag.Usage = usage
}

func usage() {
	var optsBuf bytes.Buffer

	originalOutput := flag.CommandLine.Output()

	flag.CommandLine.SetOutput(&optsBuf)
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(originalOutput)

	fmt.Fprintf(flag.CommandLine.Output(), fmtUsage, optsBuf.String())
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}

	var cmd Command
	cmdStr := flag.Arg(0)

	switch cmdStr {
	case "clean":
		cmd = &CleanCommand{OutputDir: outputDir}
	case "filenames":
		flags := flag.NewFlagSet("filenames", flag.ExitOnError)
		onlyMissing := flags.Bool("m", false, "print only missing filenames")

		if err := flags.Parse(flag.Args()[1:]); err != nil {
			log.Fatal(err)
		}

		cmd = &FilenamesCommand{OnlyMissing: *onlyMissing}
	case "generate":
		cmd = &GenerateCommand{OutputDir: outputDir}
	case "help":
		usage()
		os.Exit(0)
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "Invalid command %q\n", cmdStr)
		os.Exit(1)
	}

	steamSchema, err := NewSteamSchema(steamSchemaFile, apiKey)

	if err != nil {
		log.Fatal(err)
	}

	dotaSchema, err := NewDotaSchema(dotaSchemaFile)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Run(steamSchema, dotaSchema); err != nil {
		log.Fatal(err)
	}
}
