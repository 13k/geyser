package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/13k/geyser"
	"github.com/go-resty/resty/v2"
)

const (
	usage = `
Usage: apigen <command>

Environment variables:

DEBUG=1           Enable debugging
API_KEY=key       Use API_KEY when fetching API schema
SCHEMA=file.json  Use file.json as (locally cached) schema

Commands:

filename  Print a list of interfaces and respective filenames
generate  Generate files
clean     Remove all generated files
	`

	schemaURL = "https://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1"
)

var (
	debug      bool
	apiKey     string
	schemaFile string
)

func init() {
	debug = os.Getenv("DEBUG") != ""
	apiKey = os.Getenv("API_KEY")
	schemaFile = os.Getenv("SCHEMA")
}

type schema struct {
	API *schemaAPI `json:"apilist"`
}

type schemaAPI struct {
	Interfaces *geyser.SchemaInterfaces `json:"interfaces"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "clean", "filename", "generate":
	default:
		log.Fatalf("Invalid command %q", cmd)
	}

	var result *schema

	if schemaFile != "" {
		result = localSchema(schemaFile)
	} else {
		result = remoteSchema()
	}

	interfaces := result.API.Interfaces
	outputDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	for _, group := range interfaces.GroupByName() {
		a := &action{
			Interfaces: group,
			OutputDir:  outputDir,
		}

		if err := a.Execute(cmd); err != nil {
			log.Fatal(err)
		}
	}
}

func remoteSchema() *schema {
	if apiKey == "" {
		log.Fatal("API_KEY environment var is required when fetching remote schema")
	}

	result := &schema{}

	client := resty.New().
		SetDebug(debug).
		SetQueryParam("key", apiKey)

	resp, err := client.R().
		SetResult(result).
		Get(schemaURL)

	if err != nil {
		log.Fatal(err)
	}

	if !resp.IsSuccess() {
		log.Fatalf("Response: %s", resp.Status())
	}

	return result
}

func localSchema(filename string) *schema {
	schemaData, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	result := &schema{}

	if err = json.Unmarshal(schemaData, result); err != nil {
		log.Fatal(err)
	}

	return result
}
