package geyser

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
)

var (
	ifaceNameAppIDRegexp = regexp.MustCompile(`^(\w+)_(\d+)$`)
)

type siKey struct {
	Name  string
	AppID uint32
}

type schemaInterfacesIndex map[siKey]*SchemaInterface

func (i schemaInterfacesIndex) Find(name string, appID uint32) *SchemaInterface {
	return i[siKey{Name: name, AppID: appID}]
}

// SchemaInterfaces is a collection of SchemaInterface.
//
// The struct should be read-only.
type SchemaInterfaces struct {
	Interfaces []*SchemaInterface

	index       schemaInterfacesIndex
	indexByName map[string][]*SchemaInterface
	appIDs      []uint32
}

// MustNewSchemaInterfaces is like NewSchemaInterfaces but panics if it returned an error.
func MustNewSchemaInterfaces(interfaces ...*SchemaInterface) *SchemaInterfaces {
	s, err := NewSchemaInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return s
}

// NewSchemaInterfaces creates a new SchemaInterfaces collection.
//
// If any of the interface names cannot be parsed, it returns an error.
func NewSchemaInterfaces(interfaces ...*SchemaInterface) (*SchemaInterfaces, error) {
	s := &SchemaInterfaces{Interfaces: interfaces}

	if err := s.buildIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SchemaInterfaces) buildIndex() error {
	s.index = make(schemaInterfacesIndex)
	s.indexByName = make(map[string][]*SchemaInterface)

	for _, si := range s.Interfaces {
		key, err := si.getKey()

		if err != nil {
			return err
		}

		s.index[*key] = si
		s.indexByName[key.Name] = append(s.indexByName[key.Name], si)
	}

	return nil
}

// UnmarshalJSON deserializes JSON data into the slice of SchemaInterface.
func (s *SchemaInterfaces) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Interfaces); err != nil {
		return err
	}

	return s.buildIndex()
}

// MarshalJSON serializes the slice of SchemaInterface to JSON data.
func (s SchemaInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Interfaces)
}

// Find returns the interface with given name and appID. It returns nil if none was found.
func (s *SchemaInterfaces) Find(name string, appID uint32) *SchemaInterface {
	return s.index.Find(name, appID)
}

// Get is like Find but returns an error of type *InterfaceNotFoundError if none was found.
func (s *SchemaInterfaces) Get(name string, appID uint32) (si *SchemaInterface, err error) {
	si = s.index.Find(name, appID)

	if si == nil {
		err = &InterfaceNotFoundError{Name: name, AppID: appID}
	}

	return
}

// GroupByName groups the interfaces by name, returning a map of { name => sub-collection }.
//
// Each value in the map is a grouped SchemaInterfaces and can use group helpers.
func (s *SchemaInterfaces) GroupByName() map[string]*SchemaInterfaces {
	result := make(map[string]*SchemaInterfaces)

	for name, interfaces := range s.indexByName {
		result[name] = MustNewSchemaInterfaces(interfaces...)
	}

	return result
}

// IsGroup checks if the collection contains only interfaces with the same name.
func (s *SchemaInterfaces) IsGroup() bool {
	return len(s.indexByName) == 1
}

// GroupName returns the common name between all contained interfaces.
//
// If the collection is empty, It returns an error of type *EmptyInterfacesError.
//
// If the collection is not a group, it returns an error of type *MixedInterfacesError.
func (s *SchemaInterfaces) GroupName() (string, error) {
	if len(s.Interfaces) == 0 {
		return "", &EmptyInterfacesError{}
	}

	if !s.IsGroup() {
		return "", &MixedInterfacesError{}
	}

	return s.Interfaces[0].BaseName()
}

// AppIDs collects the AppID of all interfaces in the collection.
//
// If the collection is not a group, it returns an error of type *MixedInterfacesError.
func (s *SchemaInterfaces) AppIDs() ([]uint32, error) {
	if !s.IsGroup() {
		return nil, &MixedInterfacesError{}
	}

	if s.appIDs == nil {
		m := map[uint32]bool{}

		for key := range s.index {
			if key.AppID != 0 && !m[key.AppID] {
				s.appIDs = append(s.appIDs, key.AppID)
				m[key.AppID] = true
			}
		}
	}

	return s.appIDs, nil
}

// GroupMethods groups all intefaces methods by name.
//
// If the collection is not a group, it returns an error of type *MixedInterfacesError.
func (s *SchemaInterfaces) GroupMethods() (map[string]*SchemaMethods, error) {
	if len(s.Interfaces) == 0 {
		return nil, nil
	}

	if !s.IsGroup() {
		return nil, &MixedInterfacesError{}
	}

	result := map[string]*SchemaMethods{}

	for _, si := range s.Interfaces {
		for methodName, methods := range si.Methods.GroupByName() {
			if result[methodName] == nil {
				result[methodName] = methods
			} else {
				result[methodName] = result[methodName].Add(methods)
			}
		}
	}

	return result, nil
}

// SchemaInterface holds the specification of an API interface.
//
// The struct should be read-only.
type SchemaInterface struct {
	Name    string         `json:"name"`
	Methods *SchemaMethods `json:"methods"`

	appID    uint32
	baseName string
	key      *siKey
}

func (i *SchemaInterface) parseName() error {
	if i.key != nil {
		return nil
	}

	if matches := ifaceNameAppIDRegexp.FindStringSubmatch(i.Name); matches != nil {
		appID64, err := strconv.ParseUint(matches[2], 10, 32)

		if err != nil {
			return err
		}

		i.baseName = matches[1]
		i.appID = uint32(appID64)
	} else {
		i.baseName = i.Name
	}

	i.key = &siKey{Name: i.baseName, AppID: i.appID}

	return nil
}

func (i *SchemaInterface) getKey() (key *siKey, err error) {
	err = i.parseName()
	key = i.key
	return
}

// BaseName extracts the base name of the interface (without the AppID part).
func (i *SchemaInterface) BaseName() (s string, err error) {
	err = i.parseName()
	s = i.baseName
	return
}

// AppID extracts the AppID of the interface (without the BaseName part).
func (i *SchemaInterface) AppID() (id uint32, err error) {
	err = i.parseName()
	id = i.appID
	return
}

type smKey struct {
	Name    string
	Version int
}

type schemaMethodsIndex map[smKey]*SchemaMethod

func (i schemaMethodsIndex) Find(name string, version int) *SchemaMethod {
	return i[smKey{Name: name, Version: version}]
}

// SchemaMethods is a collection of SchemaMethod.
//
// The struct should be read-only.
type SchemaMethods struct {
	Methods []*SchemaMethod

	index       schemaMethodsIndex
	indexByName map[string][]*SchemaMethod
	versions    []int
}

// NewSchemaMethods creates a SchemaMethods collection.
func NewSchemaMethods(methods ...*SchemaMethod) *SchemaMethods {
	s := &SchemaMethods{Methods: methods}
	s.buildIndex()
	return s
}

func (s *SchemaMethods) buildIndex() {
	s.index = make(schemaMethodsIndex)
	s.indexByName = make(map[string][]*SchemaMethod)

	for _, sm := range s.Methods {
		key := sm.getKey()
		s.index[*key] = sm
		s.indexByName[key.Name] = append(s.indexByName[key.Name], sm)
	}
}

// UnmarshalJSON deserializes JSON data into the slice of SchemaMethod.
func (s *SchemaMethods) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Methods); err != nil {
		return err
	}

	s.buildIndex()
	return nil
}

// MarshalJSON serializes the slice of SchemaMethod to JSON data.
func (s SchemaMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Methods)
}

// Find returns the method with the given name and version. It returns nil if none was found.
func (s *SchemaMethods) Find(name string, version int) *SchemaMethod {
	return s.index.Find(name, version)
}

// Get is like Find but returns an error of type *InterfaceMethodNotFoundError if none was found.
func (s *SchemaMethods) Get(name string, version int) (sm *SchemaMethod, err error) {
	sm = s.index.Find(name, version)

	if sm == nil {
		err = &InterfaceMethodNotFoundError{Name: name, Version: version}
	}

	return
}

// Add creates a new SchemaMethods collection with the methods of both collections concatenated.
func (s *SchemaMethods) Add(other *SchemaMethods) *SchemaMethods {
	methods := append(s.Methods, other.Methods...)
	return NewSchemaMethods(methods...)
}

// GroupByName groups the methods by name, returning a map of { name => sub-collection }.
//
// Each value in the map is a grouped SchemaMethods and can use group helpers.
func (s *SchemaMethods) GroupByName() map[string]*SchemaMethods {
	result := make(map[string]*SchemaMethods)

	for name, methods := range s.indexByName {
		result[name] = NewSchemaMethods(methods...)
	}

	return result
}

// IsGroup checks if the collection contains only methods with the same name.
func (s *SchemaMethods) IsGroup() bool {
	return len(s.indexByName) == 1
}

// Versions collects the version of all methods in the collection.
//
// If the collection is not a group, it returns an error of type *MixedMethodsError.
func (s *SchemaMethods) Versions() ([]int, error) {
	if !s.IsGroup() {
		return nil, &MixedMethodsError{}
	}

	if s.versions == nil {
		m := map[int]bool{}

		for key := range s.index {
			if key.Version != 0 && !m[key.Version] {
				s.versions = append(s.versions, key.Version)
				m[key.Version] = true
			}
		}
	}

	return s.versions, nil
}

// SchemaMethod holds the specification of an API interface method.
//
// The struct should be read-only.
type SchemaMethod struct {
	Name       string              `json:"name"`
	Version    int                 `json:"version"`
	HTTPMethod string              `json:"httpmethod"`
	Params     *SchemaMethodParams `json:"parameters,omitempty"`

	key              *smKey
	versionPathParam string
}

func (m *SchemaMethod) getKey() *smKey {
	if m.key == nil {
		m.key = &smKey{Name: m.Name, Version: m.Version}
	}

	return m.key
}

func (m *SchemaMethod) getVersionPathParam() string {
	if m.versionPathParam == "" {
		m.versionPathParam = fmt.Sprintf("v%d", m.Version)
	}

	return m.versionPathParam
}

// ValidateParams validates the given parameters against the Params collection.
func (m *SchemaMethod) ValidateParams(params url.Values) error {
	return m.Params.Validate(params)
}

// SchemaMethodParams is a collection of SchemaMethodParam.
//
// The struct should be read-only.
type SchemaMethodParams struct {
	Params []*SchemaMethodParam
}

// NewSchemaMethodParams creates a new SchemaMethodParams collection.
func NewSchemaMethodParams(params ...*SchemaMethodParam) *SchemaMethodParams {
	return &SchemaMethodParams{Params: params}
}

// UnmarshalJSON deserializes JSON data into the slice of SchemaMethodParam.
func (s *SchemaMethodParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Params)
}

// MarshalJSON serializes the slice of SchemaMethodParam to JSON data.
func (s SchemaMethodParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Params)
}

// Validate validates the given parameters against each parameter in the collection.
func (s *SchemaMethodParams) Validate(params url.Values) error {
	for _, p := range s.Params {
		if err := p.Validate(params.Get(p.Name)); err != nil {
			return err
		}
	}

	return nil
}

// SchemaMethodParam holds the specification of an API interface method parameter.
//
// The struct should be read-only.
type SchemaMethodParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description,omitempty"`
}

// Validate validates the given value against the parameter specification.
//
// If the parameter is required and the value is empty, it returns an error of type
// *RequiredParameterError.
func (p *SchemaMethodParam) Validate(value string) error {
	if !p.Optional && value == "" {
		return &RequiredParameterError{Param: p}
	}

	return nil
}
