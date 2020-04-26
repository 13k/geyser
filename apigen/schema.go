package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"sort"

	"github.com/go-resty/resty/v2"

	"github.com/13k/geyser/schema"
)

const (
	schemaURLSteam      = "https://api.steampowered.com/ISteamWebAPIUtil/GetSupportedAPIList/v1"
	schemaURLUndocSteam = "https://raw.githubusercontent.com/SteamDatabase/UndocumentedAPI/master/api.json"
	schemaURLUndocDota  = "https://raw.githubusercontent.com/SteamDatabase/UndocumentedAPI/master/dota2.json"

	errfSchemaDuplicateInterfaces  = "Schema contains duplicate interfaces (%q)"
	errfRemoteSchemaResponse       = "Error fetching remote schema: %s"
	warnfSchemaMergeMethodConflict = "[WARN] Schema.Merge(): method %s/%s/%d already exists on destination schema"
)

var (
	errMissingSteamAPIKey = errors.New("API key is required when fetching remote Steam API schema")

	remoteSchemaReqUndocSteam = &RemoteSchemaRequest{
		URL:    schemaURLUndocSteam,
		Params: nil,
		Undoc:  true,
	}

	remoteSchemaReqUndocDota = &RemoteSchemaRequest{
		URL:    schemaURLUndocDota,
		Params: nil,
		Undoc:  true,
	}
)

type RemoteSchemaRequest struct {
	URL    string
	Params url.Values
	Undoc  bool
}

func remoteSchemaReqSteam(apiKey string) *RemoteSchemaRequest {
	params := url.Values{}
	params.Set("key", apiKey)

	return &RemoteSchemaRequest{
		URL:    schemaURLSteam,
		Params: params,
		Undoc:  false,
	}
}

type Schema struct {
	API *schema.Schema `json:"apilist"`

	relPath     string
	pkgPath     string
	pkgName     string
	filenames   map[string]string
	keyedByName map[string]*schema.Interface
}

func NewSchema(cacheFile string, remotes ...*RemoteSchemaRequest) (*Schema, error) {
	var schema *Schema
	var err error

	if cacheFile != "" {
		schema, err = NewCachedSchema(cacheFile)

		if err != nil && !os.IsNotExist(err) {
			return nil, err
		}
	}

	if schema == nil {
		var mergeSchema *Schema

		for _, req := range remotes {
			mergeSchema, err = NewRemoteSchema(req)

			if err != nil {
				return nil, err
			}

			mergeSchema.Normalize(req.Undoc)

			if schema == nil {
				schema = mergeSchema
				continue
			}

			if err = schema.Merge(mergeSchema); err != nil {
				return nil, err
			}
		}

		if cacheFile != "" {
			if err = schema.Dump(cacheFile); err != nil {
				return nil, err
			}
		}
	}

	return schema, nil
}

func NewSteamSchema(cacheFile string, apiKey string) (*Schema, error) {
	if cacheFile == "" && apiKey == "" {
		return nil, errMissingSteamAPIKey
	}

	schema, err := NewSchema(cacheFile, remoteSchemaReqSteam(apiKey), remoteSchemaReqUndocSteam)

	if err != nil {
		return nil, err
	}

	if err := schema.Validate(); err != nil {
		return nil, err
	}

	spec := schemaSpecs["steam"]

	schema.pkgPath = spec.PkgPath
	schema.pkgName = spec.PkgName
	schema.filenames = spec.Filenames
	schema.relPath = spec.RelPath

	return schema, nil
}

func NewDotaSchema(cacheFile string) (*Schema, error) {
	schema, err := NewSchema(cacheFile, remoteSchemaReqUndocDota)

	if err != nil {
		return nil, err
	}

	if err := schema.Validate(); err != nil {
		return nil, err
	}

	spec := schemaSpecs["dota2"]

	schema.pkgPath = spec.PkgPath
	schema.pkgName = spec.PkgName
	schema.filenames = spec.Filenames
	schema.relPath = spec.RelPath

	return schema, nil
}

func NewRemoteSchema(req *RemoteSchemaRequest) (*Schema, error) {
	log.Printf("Fetching remote schema from %s", req.URL)

	client := resty.New()

	resp, err := client.R().
		SetDoNotParseResponse(true).
		SetQueryParamsFromValues(req.Params).
		Get(req.URL)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf(errfRemoteSchemaResponse, resp.Status())
	}

	body := resp.RawBody()

	defer body.Close()

	dec := json.NewDecoder(body)
	schema := &Schema{}

	if err := dec.Decode(schema); err != nil {
		return nil, err
	}

	return schema, nil
}

func NewCachedSchema(cacheFile string) (*Schema, error) {
	data, err := ioutil.ReadFile(cacheFile)

	if err != nil {
		return nil, err
	}

	schema := &Schema{}

	if err = json.Unmarshal(data, schema); err != nil {
		return nil, err
	}

	return schema, nil
}

func (s *Schema) buildIndex() error {
	s.keyedByName = make(map[string]*schema.Interface)

	for _, sif := range s.API.Interfaces {
		if _, ok := s.keyedByName[sif.Name]; ok {
			return fmt.Errorf(errfSchemaDuplicateInterfaces, sif.Name)
		}

		s.keyedByName[sif.Name] = sif
	}

	return nil
}

func (s *Schema) Validate() error {
	return s.API.Validate()
}

func (s *Schema) Filename(group schema.InterfacesGroup) string {
	return s.filenames[group.Name()]
}

func (s *Schema) Normalize(undoc bool) {
	for _, i := range s.API.Interfaces {
		i.Undocumented = undoc

		for _, m := range i.Methods {
			m.Undocumented = undoc

			if m.HTTPMethod == "" {
				m.HTTPMethod = "GET"
			}
		}
	}
}

func (s *Schema) Merge(other *Schema) error {
	if err := s.buildIndex(); err != nil {
		return err
	}

	if err := other.buildIndex(); err != nil {
		return err
	}

	var result []*schema.Interface

	// append all interfaces in `s` that are not in `other`
	for _, si := range s.API.Interfaces {
		if _, ok := other.keyedByName[si.Name]; !ok {
			result = append(result, si)
		}
	}

	for _, oi := range other.API.Interfaces {
		si := s.keyedByName[oi.Name]

		// append interface in `other` that is not in `s`
		if si == nil {
			result = append(result, oi)
			continue
		}

		// merge interface that belongs to both

		var oMethods []*schema.Method

		// append all methods in `other` that are not in `s`
		for _, om := range oi.Methods {
			omKey, err := om.Key()

			if err != nil {
				return err
			}

			if _, err := si.Methods.Get(omKey); err == nil {
				// ignore conflicting method (keep method in `s`)
				log.Printf(warnfSchemaMergeMethodConflict, oi.Name, om.Name, om.Version)
				continue
			}

			oMethods = append(oMethods, om)
		}

		// prepend all methods in `s`
		methods, err := schema.NewMethods(append(si.Methods, oMethods...)...)

		if err != nil {
			return err
		}

		// create a new interface with merged methods
		mergedInterface := &schema.Interface{
			Name:    si.Name,
			Methods: methods,
		}

		result = append(result, mergedInterface)
	}

	interfaces, err := schema.NewInterfaces(result...)

	if err != nil {
		return err
	}

	s.API.Interfaces = interfaces

	return nil
}

func (s *Schema) Dump(cacheFile string) error {
	f, err := os.Create(cacheFile)

	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)

	return enc.Encode(s)
}

type interfaceGroupIterator func(string, schema.InterfacesGroup) error

func (s *Schema) eachSortedInterfaceGroup(fn interfaceGroupIterator) error {
	groups, err := s.API.Interfaces.GroupByBaseName()

	if err != nil {
		return err
	}

	groupNames := make([]string, 0, len(groups))

	for groupName := range groups {
		groupNames = append(groupNames, groupName)
	}

	sort.Strings(groupNames)

	for _, groupName := range groupNames {
		group := groups[groupName]

		if err := fn(groupName, group); err != nil {
			return err
		}
	}

	return nil
}
