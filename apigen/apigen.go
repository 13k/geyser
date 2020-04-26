package main

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/iancoleman/strcase"

	"github.com/13k/geyser/schema"
)

const (
	pkgPathNetHTTP        = "net/http"
	pkgPathTesting        = "testing"
	pkgPathGeyser         = "github.com/13k/geyser"
	pkgPathGeyserSchema   = "github.com/13k/geyser/schema"
	pkgPathGeyserSteam    = "github.com/13k/geyser/steam"
	pkgPathGeyserDota2    = "github.com/13k/geyser/dota2"
	pkgPathTestifyAssert  = "github.com/stretchr/testify/assert"
	pkgPathTestifyRequire = "github.com/stretchr/testify/require"

	pkgNameGeyser         = "geyser"
	pkgNameGeyserSchema   = "schema"
	pkgNameGeyserSteam    = "steam"
	pkgNameGeyserDota2    = "dota2"
	pkgNameTestifyAssert  = "assert"
	pkgNameTestifyRequire = "require"
	fmtPkgNameTest        = "%s_test"

	fnamefInterface = "z_%s.go"
	fnamefResults   = "z_%s_results.go"
	fnamefTests     = "z_%s_test.go"

	srcSchemaInterfacesCtor        = "MustNewSchemaInterfaces"
	srcSchemaInterface             = "SchemaInterface"
	srcSchemaInterfaceKey          = "SchemaInterfaceKey"
	srcSchemaMethodsCtor           = "MustNewSchemaMethods"
	srcSchemaMethod                = "SchemaMethod"
	srcSchemaMethodKey             = "SchemaMethodKey"
	srcSchemaMethodParamsCtor      = "NewSchemaMethodParams"
	srcSchemaMethodParam           = "SchemaMethodParam"
	srcClient                      = "Client"
	srcRequest                     = "Request"
	srcfStructCtorName             = "New%s"
	srcfSchemaVarName              = "Schema%s"
	srcIntefaceMethodNotFoundError = "InterfaceMethodNotFoundError"

	commentDisclaimer           = "Code generated by geyser. DO NOT EDIT."
	commentfDisclaimerInterface = "API interface: %s."
	commentfSchemaVar           = "%s stores the SchemaInterfaces for interface %s."
	commentfStruct              = "%s represents interface %s."
	commentfSupportedAppIDs     = "Supported AppIDs: %s."
	commentfSupportedVersions   = "Supported versions: %s."
	commentStructUndoc          = "This is an undocumented interface."
	commentfStructCtor          = "%s creates a new %s interface."
	commentfStructGetter        = "%s creates a new %s interface."
	commentfStructMethodHeader  = "%s creates a Request for interface method %s."
	commentStructMethodUndoc    = "This is an undocumented method."
	commentfParamsVersioned     = "Parameters (v%d)"
	commentParams               = "Parameters"
	commentfStructResult        = "%s holds the result of the method %s/%s."

	errfUnknownInterfaceFilename = "Unknown filename for interface %q"
	errfUnknownHTTPMethod        = "Unknown HTTP method %q of interface method %q/%q"

	testfInvalidResultType = "invalid result type %T"
	testfInvalidErrorType  = "invalid error type %T"
)

var (
	schemaSpecs = map[string]*struct {
		Name      string
		PkgPath   string
		PkgName   string
		Filenames map[string]string
		RelPath   string
	}{
		"steam": {
			PkgPath:   pkgPathGeyserSteam,
			PkgName:   pkgNameGeyserSteam,
			Filenames: filenamesSteam,
			RelPath:   "steam",
		},
		"dota2": {
			PkgPath:   pkgPathGeyserDota2,
			PkgName:   pkgNameGeyserDota2,
			Filenames: filenamesDota,
			RelPath:   "dota2",
		},
	}
)

type APIGen struct {
	baseName       string
	interfaces     schema.SchemaInterfacesGroup
	groupedMethods map[string]schema.SchemaMethodsGroup
	methodsNames   []string
	appIDs         []uint32
	requiredAppID  bool
	undoc          bool
	pkgPath        string
	pkgName        string
	testPkgName    string
	schemaVarName  string
	structName     string
	structCtorName string
	interfaceFile  *GeneratedFile
	resultsFile    *GeneratedFile
	testsFile      *GeneratedFile
}

func NewAPIGen(interfaces schema.SchemaInterfacesGroup, pkgPath, pkgName, outputDir, basefile string) (*APIGen, error) {
	baseName := interfaces.Name()
	appIDs := interfaces.AppIDs()
	groupedMethods, err := interfaces.GroupMethods()

	if err != nil {
		return nil, err
	}

	testPkgName := fmt.Sprintf(fmtPkgNameTest, pkgName)
	structName := baseName
	structCtorName := fmt.Sprintf(srcfStructCtorName, structName)
	schemaVarName := fmt.Sprintf(srcfSchemaVarName, structName)
	interfaceFname := fmt.Sprintf(fnamefInterface, basefile)
	resultsFname := fmt.Sprintf(fnamefResults, basefile)
	testsFname := fmt.Sprintf(fnamefTests, basefile)
	interfaceFile := NewGeneratedFile(filepath.Join(outputDir, interfaceFname))
	resultsFile := NewGeneratedFile(filepath.Join(outputDir, resultsFname))
	testsFile := NewGeneratedFile(filepath.Join(outputDir, testsFname))
	methodsNames := make([]string, 0, len(groupedMethods))

	for name := range groupedMethods {
		methodsNames = append(methodsNames, name)
	}

	sort.Slice(appIDs, func(i, j int) bool {
		return appIDs[i] < appIDs[j]
	})

	sort.Strings(methodsNames)

	undoc := true
	// if at least 1 interface in the group is documented, consider the struct documented
	for _, si := range interfaces {
		if !si.Undocumented {
			undoc = false
			break
		}
	}

	g := &APIGen{
		baseName:       baseName,
		interfaces:     interfaces,
		groupedMethods: groupedMethods,
		methodsNames:   methodsNames,
		appIDs:         appIDs,
		requiredAppID:  len(appIDs) > 0,
		undoc:          undoc,
		pkgPath:        pkgPath,
		pkgName:        pkgName,
		testPkgName:    testPkgName,
		structName:     structName,
		structCtorName: structCtorName,
		schemaVarName:  schemaVarName,
		interfaceFile:  interfaceFile,
		resultsFile:    resultsFile,
		testsFile:      testsFile,
	}

	return g, nil
}

func (g *APIGen) methodFuncName(methodName string) string {
	if methodName[0] >= 'a' && methodName[0] <= 'z' {
		return strcase.ToCamel(methodName)
	}

	return methodName
}

func (g *APIGen) methodResultStructName(methodName string) string {
	return g.structName + g.methodFuncName(methodName)
}
