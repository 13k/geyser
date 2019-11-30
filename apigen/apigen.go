package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/13k/geyser"
	j "github.com/dave/jennifer/jen"
)

const (
	srcPackage           = "geyser"
	srcSchemaInterface   = "SchemaInterface"
	srcSchemaMethod      = "SchemaMethod"
	srcSchemaMethodParam = "SchemaMethodParam"
	srcClient            = "Client"
	srcRequest           = "Request"

	commentDisclaimer           = "Code generated by github.com/13k/geyser/apigen. DO NOT EDIT."
	commentfDisclaimerInterface = "API interface: %s"
	commentfSchemaVar           = "%s stores the SchemaInterfaces for interface %s"
	commentfStruct1             = "%s represents interface %s"
	commentfStruct2             = "Supported AppIDs: %v"
	commentfStructCtor          = "%s creates a new %s interface"
	commentfStructGetter        = "%s creates a new %s interface"
	commentfStructMethod1       = "%s creates a Request for interface method %s"
	commentfStructMethod2       = "Supported versions: %v"
	commentfStructResult        = "%s holds the result of the method %s/%s"
)

var (
	interfaceFilenames = map[string]string{
		"IBroadcastService":              "broadcast_service",
		"ICheatReportingService":         "cheat_reporting_service",
		"IClientStats":                   "client_stats",
		"IContentServerConfigService":    "content_server_config_service",
		"IContentServerDirectoryService": "content_server_directory_service",
		"ICSGOPlayers":                   "csgo_players",
		"ICSGOServers":                   "csgo_servers",
		"ICSGOTournaments":               "csgo_tournaments",
		"IDOTA2Fantasy":                  "dota2_fantasy",
		"IDOTA2Match":                    "dota2_match",
		"IDOTA2MatchStats":               "dota2_match_stats",
		"IDOTA2StreamSystem":             "dota2_stream_system",
		"IDOTA2Ticket":                   "dota2_ticket",
		"IEconDOTA2":                     "econ_dota2",
		"IEconItems":                     "econ_items",
		"IEconService":                   "econ_service",
		"IGameNotificationsService":      "game_notifications_service",
		"IGameServersService":            "game_servers_service",
		"IGCVersion":                     "gc_version",
		"IInventoryService":              "inventory_service",
		"IPlayerService":                 "player_service",
		"IPortal2Leaderboards":           "portal2_leaderboards",
		"IPublishedFileService":          "published_file_service",
		"ISteamApps":                     "steam_apps",
		"ISteamBroadcast":                "steam_broadcast",
		"ISteamCDN":                      "steam_cdn",
		"ISteamDirectory":                "steam_directory",
		"ISteamEconomy":                  "steam_economy",
		"ISteamEnvoy":                    "steam_envoy",
		"ISteamNews":                     "steam_news",
		"ISteamRemoteStorage":            "steam_remote_storage",
		"ISteamUserAuth":                 "steam_user_auth",
		"ISteamUserOAuth":                "steam_user_oauth",
		"ISteamUserStats":                "steam_user_stats",
		"ISteamUser":                     "steam_user",
		"ISteamWebAPIUtil":               "steam_web_api_util",
		"ISteamWebUserPresenceOAuth":     "steam_web_user_presence_oauth",
		"IStoreService":                  "store_service",
		"ITFItems":                       "tf_items",
		"ITFPromos":                      "tf_promos",
		"ITFSystem":                      "tf_system",
	}

	jenImportNames = map[string]string{}
)

func jenFile() *j.File {
	f := j.NewFile(srcPackage)

	f.ImportNames(jenImportNames)

	return f
}

func jenIfErrRet(values ...j.Code) *j.Statement {
	return j.If(j.Err().Op("!=").Nil()).Block(j.Return(values...))
}

func jenIfErrRetNilErr() *j.Statement {
	return jenIfErrRet(j.Nil(), j.Err())
}

func jenHTTPMethod(method string) j.Code {
	method = strings.ToUpper(method)

	switch method {
	case http.MethodGet:
		return j.Qual("net/http", "MethodGet")
	case http.MethodPost:
		return j.Qual("net/http", "MethodPost")
	case http.MethodPut:
		return j.Qual("net/http", "MethodPut")
	default:
		return nil
	}
}

func jenSchemaInterfaceAddr() *j.Statement   { return j.Op("&").Id(srcSchemaInterface) }
func jenSchemaInterfacePtr() *j.Statement    { return j.Op("*").Id(srcSchemaInterface) }
func jenSchemaMethodAddr() *j.Statement      { return j.Op("&").Id(srcSchemaMethod) }
func jenSchemaMethodParamAddr() *j.Statement { return j.Op("&").Id(srcSchemaMethodParam) }
func jenClientPtr() *j.Statement             { return j.Op("*").Id(srcClient) }
func jenRequestAddr() *j.Statement           { return j.Op("&").Id(srcRequest) }
func jenRequestPtr() *j.Statement            { return j.Op("*").Id(srcRequest) }

type apiGen struct {
	BaseName              string
	Interfaces            *geyser.SchemaInterfaces
	AppIDs                []uint32
	RequiredAppID         bool
	Filename              string
	ResultsFilename       string
	SchemaVarName         string
	StructName            string
	StructConstructorName string
}

func newAPIGen(sis *geyser.SchemaInterfaces) (*apiGen, error) {
	baseName, err := sis.GroupName()

	if err != nil {
		return nil, err
	}

	baseFilename := interfaceFilenames[baseName]

	if baseFilename == "" {
		err = fmt.Errorf("Unknown filename for interface %q", baseName)
		return nil, err
	}

	structName := strings.TrimPrefix(baseName, "I")
	appIDs, err := sis.AppIDs()

	if err != nil {
		return nil, err
	}

	g := &apiGen{
		BaseName:              baseName,
		Interfaces:            sis,
		AppIDs:                appIDs,
		RequiredAppID:         len(appIDs) > 0,
		Filename:              baseFilename + ".go",
		ResultsFilename:       baseFilename + "_results.go",
		SchemaVarName:         "Schema" + structName,
		StructName:            structName,
		StructConstructorName: "New" + structName,
	}

	return g, nil
}

func (g *apiGen) genSIDecl(si *geyser.SchemaInterface) (j.Code, error) {
	var methodDecls []j.Code

	for _, sm := range si.Methods.Methods {
		methodDecl, err := g.genSMDecl(si, sm)

		if err != nil {
			return nil, err
		}

		methodDecls = append(methodDecls, j.Line().Add(methodDecl))
	}

	methodDecls = append(methodDecls, j.Line())
	methodsDecl := j.Id("NewSchemaMethods").Call(methodDecls...)

	code := jenSchemaInterfaceAddr().Values(j.Dict{
		j.Id("Name"):    j.Lit(si.Name),
		j.Id("Methods"): methodsDecl,
	})

	return code, nil
}

func (g *apiGen) genSMDecl(si *geyser.SchemaInterface, sm *geyser.SchemaMethod) (j.Code, error) {
	httpMethod := jenHTTPMethod(sm.HTTPMethod)

	if httpMethod == nil {
		err := fmt.Errorf("Unknown HTTP method %q in interface method %q.%q", sm.HTTPMethod, si.Name, sm.Name)
		return nil, err
	}

	var paramDecls []j.Code

	for _, smp := range sm.Params.Params {
		paramDecl := g.genSMPDecl(si, sm, smp)
		paramDecls = append(paramDecls, j.Line().Add(paramDecl))
	}

	paramDecls = append(paramDecls, j.Line())
	paramsSliceDecl := j.Id("NewSchemaMethodParams").Call(paramDecls...)

	code := jenSchemaMethodAddr().Values(j.Dict{
		j.Id("Name"):       j.Lit(sm.Name),
		j.Id("Version"):    j.Lit(sm.Version),
		j.Id("HTTPMethod"): httpMethod,
		j.Id("Params"):     paramsSliceDecl,
	})

	return code, nil
}

func (g *apiGen) genSMPDecl(
	_ *geyser.SchemaInterface,
	_ *geyser.SchemaMethod,
	smp *geyser.SchemaMethodParam,
) j.Code {
	return jenSchemaMethodParamAddr().Values(j.Dict{
		j.Id("Name"):        j.Lit(smp.Name),
		j.Id("Type"):        j.Lit(smp.Type),
		j.Id("Optional"):    j.Lit(smp.Optional),
		j.Id("Description"): j.Lit(smp.Description),
	})
}

func (g *apiGen) genSchemaDecl() (j.Code, error) {
	var siDecls []j.Code

	for _, si := range g.Interfaces.Interfaces {
		siDecl, err := g.genSIDecl(si)

		if err != nil {
			return nil, err
		}

		siDecls = append(siDecls, j.Line().Add(siDecl))
	}

	siDecls = append(siDecls, j.Line())
	comment := fmt.Sprintf(commentfSchemaVar, g.SchemaVarName, g.BaseName)

	schemaDecl := j.
		Comment(comment).
		Line().
		Var().
		Id(g.SchemaVarName).
		Op("=").
		Id("MustNewSchemaInterfaces").
		Call(siDecls...)

	return schemaDecl, nil
}

func (g *apiGen) genStructDef() j.Code {
	comments := []string{
		fmt.Sprintf(commentfStruct1, g.StructName, g.BaseName),
	}

	if g.RequiredAppID {
		comments = append(comments, fmt.Sprintf(commentfStruct2, g.AppIDs))
	}

	fields := []j.Code{
		j.Id("Client").Add(jenClientPtr()),
		j.Id("Interface").Add(jenSchemaInterfacePtr()),
	}

	code := j.Null()

	for _, c := range comments {
		code.Comment(c).Line()
	}

	return code.Type().Id(g.StructName).Struct(fields...)
}

func (g *apiGen) genStructCtor() j.Code {
	comment := fmt.Sprintf(commentfStructCtor, g.StructConstructorName, g.StructName)

	params := []j.Code{
		j.Id("c").Add(jenClientPtr()),
	}

	if g.RequiredAppID {
		params = append(params, j.Id("appID").Uint32())
	}

	retTypes := []j.Code{
		j.Op("*").Id(g.StructName),
		j.Error(),
	}

	getArgs := []j.Code{
		j.Lit(g.BaseName),
	}

	if g.RequiredAppID {
		getArgs = append(getArgs, j.Id("appID"))
	} else {
		getArgs = append(getArgs, j.Lit(0))
	}

	body := []j.Code{
		j.List(j.Id("si"), j.Err()).Op(":=").Id(g.SchemaVarName).Dot("Get").Call(getArgs...),
		j.Line(),
		jenIfErrRetNilErr(),
		j.Line(),
		j.Id("s").Op(":=").Op("&").Id(g.StructName).Values(j.Dict{
			j.Id("Client"):    j.Id("c"),
			j.Id("Interface"): j.Id("si"),
		}),
		j.Line(),
		j.Return(j.Id("s"), j.Nil()),
	}

	return j.
		Comment(comment).Line().
		Func().
		Id(g.StructConstructorName).
		Params(params...).
		Params(retTypes...).
		Block(body...)
}

func (g *apiGen) genStructGetter() j.Code {
	comment := fmt.Sprintf(commentfStructGetter, g.StructName, g.StructName)
	getterReceiver := j.Id("c").Add(jenClientPtr())
	getterParams := []j.Code{}
	ctorParams := []j.Code{j.Id("c")}

	if g.RequiredAppID {
		getterParams = append(getterParams, j.Id("appID").Uint32())
		ctorParams = append(ctorParams, j.Id("appID"))
	}

	retTypes := []j.Code{
		j.Op("*").Id(g.StructName),
		j.Error(),
	}

	fnBody := []j.Code{
		j.Return(j.Id(g.StructConstructorName).Call(ctorParams...)),
	}

	return j.
		Comment(comment).Line().
		Func().
		Parens(getterReceiver).
		Id(g.StructName).
		Params(getterParams...).
		Params(retTypes...).
		Block(fnBody...)
}

func (g *apiGen) genMethod(name string, sms *geyser.SchemaMethods) (j.Code, error) {
	versions, err := sms.Versions()

	if err != nil {
		return nil, err
	}

	requiredVersion := len(versions) > 1

	comments := []string{
		fmt.Sprintf(commentfStructMethod1, name, name),
	}

	if requiredVersion {
		comments = append(comments, fmt.Sprintf(commentfStructMethod2, versions))
	}

	structReceiver := j.Id("i").Op("*").Id(g.StructName)

	params := []j.Code{}

	if requiredVersion {
		params = append(params, j.Id("version").Int())
	}

	retTypes := j.List(
		jenRequestPtr(),
		j.Error(),
	)

	var getParamsLast j.Code

	if requiredVersion {
		getParamsLast = j.Id("version")
	} else {
		getParamsLast = j.Lit(versions[0])
	}

	getParams := []j.Code{j.Lit(name), getParamsLast}

	resultStructName := fmt.Sprintf("%s%s", g.StructName, name)

	reqInst := j.Id("req").Op(":=").Add(jenRequestAddr()).Values(j.Dict{
		j.Id("Client"):    j.Id("i").Dot("Client"),
		j.Id("Interface"): j.Id("i").Dot("Interface"),
		j.Id("Method"):    j.Id("sm"),
		j.Id("Result"):    j.Op("&").Id(resultStructName).Values(),
	})

	body := []j.Code{
		j.List(j.Id("sm"), j.Err()).Op(":=").Id("i").Dot("Interface").Dot("Methods").Dot("Get").Call(getParams...),
		j.Line(),
		jenIfErrRetNilErr(),
		j.Line(),
		reqInst,
		j.Line(),
		j.Return(j.Id("req"), j.Nil()),
	}

	code := j.Null()

	for _, c := range comments {
		code.Comment(c).Line()
	}

	code.Func().Parens(structReceiver).Id(name).Params(params...).Parens(retTypes).Block(body...)

	return code, nil
}

func (g *apiGen) genMethods() (j.Code, error) {
	grouped, err := g.Interfaces.GroupMethods()

	if err != nil {
		return nil, err
	}

	code := j.Null()

	for name, group := range grouped {
		method, err := g.genMethod(name, group)

		if err != nil {
			return nil, err
		}

		code.Line().Add(method)
	}

	return code, nil
}

func (g *apiGen) genResult(name string) j.Code {
	resultStructName := fmt.Sprintf("%s%s", g.StructName, name)
	comment := fmt.Sprintf(commentfStructResult, resultStructName, g.BaseName, name)

	return j.
		Comment(comment).
		Line().
		Type().
		Id(resultStructName).
		Struct()
}

func (g *apiGen) genResults() (j.Code, error) {
	grouped, err := g.Interfaces.GroupMethods()

	if err != nil {
		return nil, err
	}

	code := j.Null()

	for name := range grouped {
		result := g.genResult(name)
		code.Line().Add(result)
	}

	return code, nil
}

func (g *apiGen) InterfaceFile() (*j.File, error) {
	f := jenFile()

	f.HeaderComment(commentDisclaimer)
	f.HeaderComment(fmt.Sprintf(commentfDisclaimerInterface, g.BaseName))

	schemaDecl, err := g.genSchemaDecl()

	if err != nil {
		return nil, err
	}

	f.Add(schemaDecl)
	f.Add(g.genStructDef())
	f.Add(g.genStructCtor())
	f.Add(g.genStructGetter())

	methods, err := g.genMethods()

	if err != nil {
		return nil, err
	}

	f.Add(methods)

	return f, nil
}

func (g *apiGen) ResultsFile() (*j.File, error) {
	f := jenFile()

	f.HeaderComment(commentDisclaimer)

	results, err := g.genResults()

	if err != nil {
		return nil, err
	}

	f.Add(results)

	return f, nil
}
