package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/13k/geyser"
	"github.com/sirupsen/logrus"
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

Filenames options:

%s

`
)

var (
	apiKey          string
	steamSchemaFile string
	dotaSchemaFile  string
	outputDir       string

	filenamesFlags       *flag.FlagSet
	filenamesOnlyMissing bool
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

	filenamesFlags = flag.NewFlagSet("filenames", flag.ExitOnError)
	filenamesFlags.BoolVar(&filenamesOnlyMissing, "m", false, "print only missing filenames")

	flag.Usage = usage
}

func usage() {
	var optionsBuf bytes.Buffer

	flag.CommandLine.SetOutput(&optionsBuf)
	flag.PrintDefaults()
	flag.CommandLine.SetOutput(os.Stderr)

	var filenamesBuf bytes.Buffer

	filenamesFlags.SetOutput(&filenamesBuf)
	filenamesFlags.PrintDefaults()
	filenamesFlags.SetOutput(os.Stderr)

	fmt.Fprintf(os.Stderr, fmtUsage, optionsBuf.String(), filenamesBuf.String())
}

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		os.Exit(1)
	}

	log := logrus.New()
	log.SetOutput(os.Stderr)

	var cmd Command
	cmdStr := flag.Arg(0)

	switch cmdStr {
	case "clean":
		cmd = &CleanCommand{
			OutputDir: outputDir,
			Log:       log,
		}
	case "filenames":
		if err := filenamesFlags.Parse(flag.Args()[1:]); err != nil {
			fatal(log, err)
		}

		cmd = &FilenamesCommand{
			OnlyMissing: filenamesOnlyMissing,
		}
	case "generate":
		cmd = &GenerateCommand{
			OutputDir: outputDir,
			Log:       log,
		}
	case "help":
		usage()
		os.Exit(0)
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "Invalid command %q\n", cmdStr)
		os.Exit(1)
	}

	steamSchema, err := NewSteamSchema(steamSchemaFile, apiKey)

	if err != nil {
		fatal(log, err)
	}

	dotaSchema, err := NewDotaSchema(dotaSchemaFile)

	if err != nil {
		fatal(log, err)
	}

	if err := cmd.Run(steamSchema, dotaSchema); err != nil {
		fatal(log, err)
	}
}

func fatal(log logrus.FieldLogger, err error) {
	l := log.WithError(err)

	switch e := err.(type) {
	case *geyser.InterfaceNotFoundError:
		l.WithField("key", e.Key).Fatal()
	case *geyser.InvalidInterfaceNameError:
		l.WithFields(logrus.Fields{
			"interface": e.Interface.Name,
			"err":       e.Err,
		}).Fatal()
	case *geyser.InterfaceMethodNotFoundError:
		l.WithField("key", e.Key).Fatal()
	case *geyser.InvalidMethodNameError:
		l.WithFields(logrus.Fields{
			"method":  e.Method.Name,
			"version": e.Method.Version,
		}).Fatal()
	case *geyser.InvalidMethodVersionError:
		l.WithFields(logrus.Fields{
			"method":  e.Method.Name,
			"version": e.Method.Version,
		}).Fatal()
	case *geyser.InvalidMethodHTTPMethodError:
		l.WithFields(logrus.Fields{
			"method":      e.Method.Name,
			"version":     e.Method.Version,
			"http_method": e.Method.HTTPMethod,
		}).Fatal()
	default:
		l.Fatal()
	}
}
