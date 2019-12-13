/*
Package geyser implements an HTTP client for the Steam Web API.

API interfaces and methods are generated automatically by sub-package "apigen" and are structured
as a tree.

NOTE: A non-partner api key is used to fetch the API schema, so only non-partner interfaces are
generated.

Structure

Interfaces are grouped by their base name, with the leading "I" and the AppID removed. For
example: all "IGCVersion_<ID>" interfaces are grouped into a single interface struct named
"GCVersion" that is instantiated by passing the AppID as parameter ("GCVersion(version)").

Interfaces that don't have an AppID do not require an AppID parameter (e.g.: "ISteamApps" ->
"SteamApps()"). Interfaces that have a single AppID still require an AppID parameter.

The same grouping is done for interface methods. All methods with the same name but different
versions are grouped by name and the version is parameterized ("Method(version)").

Methods that have a single version do not require a version parameter ("Method()").

The workflow for requesting an interface method is:

  Client
  |-> Interface([appID])
      |-> Method([version])
          |-> Request
              |-> Configure Request (optional)
              |-> Request execution
                  |-> HTTP response
                      |-> Check HTTP response status
                      |-> Parse response body or use the configured result value

Client is the root element. Interfaces are instantiated with methods in the Client struct
(possibly requiring an AppID parameter). Interface methods are struct methods of the interface
struct (possibly requiring a version parameter).

Parameters to the request must be set in the Request configuration step (see Request.SetOptions).
Parameters are validated automatically conforming the interface method schema and the request
execution will return an error if the validation failed.

Results of requests can be obtained in two ways: parsing the response body manually, or
configuring the result object in the Request before executing.

Requests are always sent using "format=json", so response bodies (seem to) always be in JSON
format. For manual parsing, checking "Content-Type" response header is recommended. When setting
the result object in the Request (see Request.SetResult), a "JSON-unmarshable" object is
expected.

All requests created by generated methods are pre-configured with a corresponding result struct.

Result structs are also (initially) automatically generated, they are named by concatenating the
interface struct name and the method name (e.g.: "ISteamApps/GetAppList" ->
"SteamAppsGetAppList").

Given that each response has a specific format and cannot be automatically generated, most of
these result structs are empty (so they are useless, but response bodies are still available for
manual parsing). When a response format is known, the result struct will be updated and won't be
automatically (re)generated.

Look for non-empty result structs for each interface method to see if there's any available at
the time. Contributions to implement proper result structs are welcome!

NOTE: Since not all similarly named interfaces with different AppIDs are necessarily identical,
this grouping can result in generated interface struct methods that are not present in a given
interface. For example, given an API schema of:

  {
    IFace_1[GetValue(v1), GetValue(v2)],
    IFace_2[GetValue(v1), GetValue(v2), GetSpecificAppID2Value(v1)]
  }

where IFace_2 is a superset of IFace_1, the resulting structure would be:

  {
    Face(appID)[GetValue(version), GetSpecificAppID2Value()]
  }

GetValue and GetSpecificAppID2Value are generated struct methods of Face, regardless of AppID and
version.

Accessing Face(1).GetValue(1) is valid, so is accessing Face(1).GetValue(2). But accessing
Face(1).GetSpecificAppID2Value() returns an error.

Schema

API schema is defined by "Schema*" types that compose the tree-like structure of the generated
API and provide basic methods for navigating and manipulation of the schema.

A schema starts with a root Schema node that contains API information and a SchemaInterfaces
collection of interfaces.

SchemaInterfaces is a collection of SchemaInterface interfaces. It also provides helper methods
to group interfaces by name, extract AppIDs and so on.

SchemaInterface holds the specificiation of the interface and a SchemaMethods collection of
methods.

SchemaMethods is a collection of SchemaMethod and provides helpers methods to group methods by
name, extract versions and so on.

SchemaMethod represents the specification of an interface method. It also contains a collection
of parameters SchemaMethodParams.

SchemaMethodParams is a collection of SchemaMethodParam. It also provides helpers for parameter
validation.

SchemaMethodParam holds the specification of an interface method parameter.

The specification for each interface is exposed through "Schema<InterfaceName>" public variables.
These can also be used direcly by users for any other purpose, including instantiating individual
interface structs directly.

All of the collection types provide JSON encoding methods that help in
serialization/deserialization to/from JSON format. These types can be used directly when
deserializing a JSON schema.

Official Documentation

Steam Web API documentation can be found at:

https://partner.steamgames.com/doc/webapi_overview

https://developer.valvesoftware.com/wiki/Steam_Web_API

https://wiki.teamfortress.com/wiki/WebAPI

Undocumented API

There are several interfaces and methods not present in the official API schema (returned by
"ISteamWebAPIUtil/GetSupportedAPIList"), including game-specific schemas.

These undocumented interfaces are fetched from third-party sources and are also generated along
with the official one.

Specifications of undocumented methods don't define any parameters, so no validation is performed
or any documentation is generated. More importantly, they also don't define the HTTP method to be
used. For now, these methods default to GET HTTP method.

When an interface or method originates from an undocumented source, they'll have a comment
indicating it.

A comprehensive list of interfaces and methods, documented and undocumented, can be found at
https://steamapi.xpaw.me

Other APIs

APIs that are available on different hosts and have different schemas are generated under
sub-packages.

These sub-packages will contain their own "Client" struct that are thin wrappers around the base
Client.

Schema structure and code generation rules are the same as for the main package.

Currently, only Dota 2 is known to geyser and is implemented in sub-package "dota2".

*/
package geyser
