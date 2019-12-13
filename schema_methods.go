package geyser

import (
	"encoding/json"
)

type schemaMethodsIndex map[SchemaMethodKey]*SchemaMethod

func (i schemaMethodsIndex) Find(name string, version int) *SchemaMethod {
	return i[SchemaMethodKey{Name: name, Version: version}]
}

// SchemaMethods is a collection of `SchemaMethod`.
//
// The struct should be read-only.
type SchemaMethods struct {
	Methods []*SchemaMethod

	index       schemaMethodsIndex
	indexByName map[string][]*SchemaMethod
	versions    []int
}

// MustNewSchemaMethods is like `NewSchemaMethods` but panics if it returned an
// error.
func MustNewSchemaMethods(methods ...*SchemaMethod) *SchemaMethods {
	s, err := NewSchemaMethods(methods...)

	if err != nil {
		panic(err)
	}

	return s
}

// NewSchemaMethods creates a new collection.
//
// Returns errors described in `SchemaMethod.Key`.
func NewSchemaMethods(methods ...*SchemaMethod) (*SchemaMethods, error) {
	s := &SchemaMethods{Methods: methods}

	if err := s.buildIndex(); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SchemaMethods) buildIndex() error {
	s.index = make(schemaMethodsIndex)
	s.indexByName = make(map[string][]*SchemaMethod)

	for _, sm := range s.Methods {
		key, err := sm.Key()

		if err != nil {
			return err
		}

		s.index[*key] = sm
		s.indexByName[key.Name] = append(s.indexByName[key.Name], sm)
	}

	return nil
}

// UnmarshalJSON deserializes JSON data into `Methods` field.
func (s *SchemaMethods) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Methods); err != nil {
		return err
	}

	return s.buildIndex()
}

// MarshalJSON serializes `Methods` field to JSON data.
func (s SchemaMethods) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Methods)
}

// Find returns the method with the given name and version. It returns nil if
// none was found.
func (s *SchemaMethods) Find(name string, version int) *SchemaMethod {
	return s.index.Find(name, version)
}

// Get is like Find but returns an error of type `*InterfaceMethodNotFoundError`
// if none was found.
func (s *SchemaMethods) Get(name string, version int) (sm *SchemaMethod, err error) {
	sm = s.index.Find(name, version)

	if sm == nil {
		err = &InterfaceMethodNotFoundError{Name: name, Version: version}
	}

	return
}

// Add creates a new collection with the methods of both collections
// concatenated.
func (s *SchemaMethods) Add(other *SchemaMethods) (*SchemaMethods, error) {
	methods := append(s.Methods, other.Methods...)
	return NewSchemaMethods(methods...)
}

// GroupByName groups the methods by name, returning a map of `{ name => sub-collection }`.
//
// Each value in the map is a grouped `SchemaMethods` and can use group helpers.
//
// Returns errors described in `NewSchemaMethods`.
func (s *SchemaMethods) GroupByName() (map[string]*SchemaMethods, error) {
	var err error
	result := make(map[string]*SchemaMethods)

	for name, methods := range s.indexByName {
		if result[name], err = NewSchemaMethods(methods...); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// IsGroup checks if the collection contains only methods with the same name.
func (s *SchemaMethods) IsGroup() bool {
	return len(s.indexByName) == 1
}

// Versions collects the version of all methods in the collection.
//
// If the collection is not a group, it returns error `ErrMixedMethods`.
func (s *SchemaMethods) Versions() ([]int, error) {
	if !s.IsGroup() {
		return nil, ErrMixedMethods
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
