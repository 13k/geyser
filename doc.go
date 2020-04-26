/*
Package geyser provides HTTP clients for Steam API and Steam API-like endpoints.

NOTE: To avoid confusion between Go terminology and abstract API entities (terms like "interface"
and "method"), API entity names are expressed using capitalized words. Concrete Go names and
statements are surrounded by backticks.

Steam APIs are structured as trees, with the base endpoint being the root node. The root node has a
set of Interface nodes as children. Interfaces can be parameterized by AppID, so one Interface with
2 different versions results in 2 different nodes. Each Interface has a collection of Method leaf nodes.
Methods can be versioned, so a Method with 2 different versions results in 2 different nodes, and can
contain a set of Parameters (HTTP request parameters).

Schemas and client code are generated automatically by sub-package "apigen".

NOTE: A non-partner api key is used to fetch the API schema, so only non-partner Interfaces are
generated.

Schema

Types in sub-package "schema" represent the tree structure of an API schema:

		Schema
		|-- Interface<appID>
		    |-- Method<version>

A schema starts with a root `Schema` node that contains API information and an `Interfaces`
collection.

`Interfaces` is a collection of `Interface` nodes. It also provides helper methods to group
`Interface`s by name, extract AppIDs and so on.

`Interface` holds the specificiation of the Interface and a `Methods` collection.

`Methods` is a collection of `Method` nodes and provides helpers methods to group `Method`s by name,
extract versions and so on.

`Method` represents the specification of an Interface Method. It also contains a collection of
parameters `MethodParams`.

`MethodParams` is a collection of `MethodParam`. It also provides helpers for parameter validation.

`MethodParam` holds the specification of an Interface Method Parameter.

The specification for each Interface is exposed through `Schema<InterfaceName>` public variables.
These can also be used direcly by users for any other purpose, including instantiating individual
interface structs directly.

Node types provide JSON tags to encode/decode to/from JSON format. All of the collection types
provide JSON encoding/decoding methods. A root `Schema` node is then well-suited to encode/decode
API schemas to/from JSON format (i.e., when consuming "ISteamWebAPIUtil/GetSupportedAPIList").

Client Structure

Interfaces are grouped by their base name, with the AppID (if present) removed. For example: all
"IGCVersion_<ID>" Interfaces are grouped into a single struct named "IGCVersion" that is
instantiated by passing the AppID as parameter (`NewIGCVersion(client, appID)` or
`client.IGCVersion(appID)`).

Interfaces that don't have an AppID do not require an AppID parameter (e.g.: "ISteamApps" ->
`NewISteamApps(client)` or `client.NewISteamApps()`). Interfaces that have only one AppID still
require an AppID parameter.

The same grouping is done for Interface Methods. All Methods with the same name but different
versions are grouped by name and the version is parameterized (`iface.MethodName(version)`). Methods
that have a single version do not require a version parameter (`iface.MethodName()`).

The workflow to perform an HTTP request to an Interface Method is:

  Client
  |-> Interface([appID])
      |-> Method([version])
          |-> Request
              |-> Configure Request (optional)
              |-> Request execution
                  |-> HTTP response
                      |-> Check HTTP response status
                      |-> Parse response body or use the configured result value

Client is the root element. Interfaces are instantiated with methods in the `Client` struct
(possibly requiring an AppID parameter). Interface Methods are struct methods of the Interface
struct (possibly requiring a version parameter).

Parameters to the request must be set in the `Request` configuration step (see
`Request.SetOptions`). Parameters are validated automatically conforming the Interface Method
specification and the request execution will return an error if the validation failed.

The meaning and functionality of each Parameter for all Interface Methods is not well known, neither
are they well documented (here, in official documentation or anywhere else). The user will have to
do some experimenting and research to figure it out. That said, whenever Parameters are known, there
may exist manually implemented helper functions to generate values (either query parameters or form
data, as `net/url.Values`). These helper functions are named by concatenating the Interface struct
name, the Method name and either "Params" or "FormData", depending on the HTTP method (e.g.:
"ISteamRemoteStorage/GetCollectionDetails" -> "ISteamRemoteStorageGetCollectionDetailsFormData").

Results of requests can be obtained in two ways: parsing the response body manually, or
configuring the result object in the `Request` before executing.

Requests are always sent using `format=json`, so response bodies are (or seem to be) always in JSON
format. For manual parsing, checking "Content-Type" response header is recommended. When setting the
result object in the `Request` (see `Request.SetResult`), a JSON-unmarshable object is expected.

Given that each response has a specific format and result structs cannot be automatically generated,
decoding of a specific response is left to the user. When a response format is known, there will be
a manually implemented struct with the appropriate fields. These result structs are named by
concatenating the Interface struct name and the Method name (e.g.: "ISteamApps/GetAppList" ->
"ISteamAppsGetAppList"). Look for these result structs for each Interface Method to see if there's
any available at the time. Contributions to implement proper result structs are welcome!

Since not all similarly named Interfaces with different AppIDs are necessarily identical, the
Interface/Method grouping can result in generated Interface struct methods that are not present in a
given Interface. For example, given a (pseudo) schema:

  {
    IUsers_100 : [GetName/v1, GetName/v2],
    IUsers_200 : [GetName/v2, GetAge/v1],
  }

the resulting (pseudo) structure would be:

  {
    IUsers(appID) : [GetName(version), GetAge()],
  }

GetName and GetAge are generated struct methods of IUsers, regardless of AppID and version.

Accessing `IUsers(100).GetName(1)`, `IUsers(100).GetName(2)`, `IUsers(200).GetName(2)` and
`IUsers(200).GetAge()` is valid. But accessing `IUsers(100).GetAge()` and `IUsers(200).GetName(1)`
is invalid (returns error).

Available APIs

APIs that are available on different hosts and have different schemas are generated under
self-contained sub-packages.

These sub-packages will contain their own `Client` struct that are thin wrappers around the base
`Client`.

Currently, these are available:

    +===============+=============+
    | API           | sub-package |
    +===============+=============+
    | Steam Web API | steam       |
    +---------------+-------------+
    | Dota2 Web API | dota2       |
    +---------------+-------------+

Official Documentation

Steam Web API documentation can be found at:

https://partner.steamgames.com/doc/webapi_overview

https://developer.valvesoftware.com/wiki/Steam_Web_API

https://wiki.teamfortress.com/wiki/WebAPI

Undocumented APIs

There are several Interfaces and Methods not present in the official API schema (returned by
"ISteamWebAPIUtil/GetSupportedAPIList"), including game-specific schemas.

These undocumented Interfaces are fetched from third-party sources and are also generated along
with the official one.

Specifications of undocumented Methods don't define any Parameters, so no validation is performed
or any documentation is generated. More importantly, they also don't define the HTTP method to be
used. For now, these Methods default to GET HTTP method.

When an Interface or Method originates from an undocumented source, they'll have a comment
indicating it.

A comprehensive list of Interfaces and Methods, documented and undocumented, can be found at
https://steamapi.xpaw.me
*/
package geyser
