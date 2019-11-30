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

type SchemaInterfaces struct {
	Interfaces []*SchemaInterface

	index       schemaInterfacesIndex
	indexByName map[string][]*SchemaInterface
	appIDs      []uint32
}

func MustNewSchemaInterfaces(interfaces ...*SchemaInterface) *SchemaInterfaces {
	s, err := NewSchemaInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return s
}

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

func (s *SchemaInterfaces) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Interfaces); err != nil {
		return err
	}

	return s.buildIndex()
}

func (s SchemaInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Interfaces)
}

func (s *SchemaInterfaces) Find(name string, appID uint32) *SchemaInterface {
	return s.index.Find(name, appID)
}

func (s *SchemaInterfaces) Get(name string, appID uint32) (si *SchemaInterface, err error) {
	si = s.index.Find(name, appID)

	if si == nil {
		err = &InterfaceNotFoundError{Name: name, AppID: appID}
	}

	return
}

func (s *SchemaInterfaces) GroupByName() map[string]*SchemaInterfaces {
	result := make(map[string]*SchemaInterfaces)

	for name, interfaces := range s.indexByName {
		result[name] = MustNewSchemaInterfaces(interfaces...)
	}

	return result
}

func (s *SchemaInterfaces) IsGroup() bool {
	return len(s.indexByName) == 1
}

func (s *SchemaInterfaces) GroupName() (string, error) {
	if !s.IsGroup() {
		return "", &MixedInterfacesError{}
	}

	if len(s.Interfaces) == 0 {
		return "", &EmptyInterfacesError{}
	}

	return s.Interfaces[0].BaseName()
}

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

func (s *SchemaInterfaces) GroupMethods() (map[string]*SchemaMethods, error) {
	if !s.IsGroup() {
		return nil, &MixedInterfacesError{}
	}

	if len(s.Interfaces) == 0 {
		return nil, nil
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

func (i *SchemaInterface) BaseName() (s string, err error) {
	err = i.parseName()
	s = i.baseName
	return
}

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

type SchemaMethods struct {
	Methods []*SchemaMethod

	index       schemaMethodsIndex
	indexByName map[string][]*SchemaMethod
	versions    []int
}

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

func (s *SchemaMethods) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Methods); err != nil {
		return err
	}

	s.buildIndex()
	return nil
}

func (s SchemaMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Methods)
}

func (s *SchemaMethods) Find(name string, version int) *SchemaMethod {
	return s.index.Find(name, version)
}

func (s *SchemaMethods) Get(name string, version int) (sm *SchemaMethod, err error) {
	sm = s.index.Find(name, version)

	if sm == nil {
		err = &InterfaceMethodNotFoundError{Name: name, Version: version}
	}

	return
}

func (s *SchemaMethods) Add(other *SchemaMethods) *SchemaMethods {
	methods := append(s.Methods, other.Methods...)
	return NewSchemaMethods(methods...)
}

func (s *SchemaMethods) GroupByName() map[string]*SchemaMethods {
	result := make(map[string]*SchemaMethods)

	for name, methods := range s.indexByName {
		result[name] = NewSchemaMethods(methods...)
	}

	return result
}

func (s *SchemaMethods) IsGroup() bool {
	return len(s.indexByName) == 1
}

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

func (m *SchemaMethod) ValidateParams(params url.Values) error {
	return m.Params.Validate(params)
}

type SchemaMethodParams struct {
	Params []*SchemaMethodParam
}

func NewSchemaMethodParams(params ...*SchemaMethodParam) *SchemaMethodParams {
	return &SchemaMethodParams{Params: params}
}

func (s *SchemaMethodParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Params)
}

func (s SchemaMethodParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Params)
}

func (s *SchemaMethodParams) Validate(params url.Values) error {
	for _, p := range s.Params {
		if err := p.Validate(params.Get(p.Name)); err != nil {
			return err
		}
	}

	return nil
}

type SchemaMethodParam struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Optional    bool   `json:"optional"`
	Description string `json:"description,omitempty"`
}

func (p *SchemaMethodParam) Validate(value string) error {
	if !p.Optional && value == "" {
		return &RequiredParameterError{Param: p}
	}

	return nil
}
