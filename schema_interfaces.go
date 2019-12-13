package geyser

import (
	"encoding/json"
)

type schemaInterfacesIndex map[SchemaInterfaceKey]*SchemaInterface

func (i schemaInterfacesIndex) Find(name string, appID uint32) *SchemaInterface {
	return i[SchemaInterfaceKey{Name: name, AppID: appID}]
}

// SchemaInterfaces is a collection of `SchemaInterface`.
//
// The struct should be read-only.
type SchemaInterfaces struct {
	Interfaces []*SchemaInterface

	index       schemaInterfacesIndex
	indexByName map[string][]*SchemaInterface
	appIDs      []uint32
}

// MustNewSchemaInterfaces is like `NewSchemaInterfaces` but panics if it
// returned an error.
func MustNewSchemaInterfaces(interfaces ...*SchemaInterface) *SchemaInterfaces {
	s, err := NewSchemaInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return s
}

// NewSchemaInterfaces creates a new collection.
//
// Returns errors described in `SchemaInterface.Key`.
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
		key, err := si.Key()

		if err != nil {
			return err
		}

		s.index[*key] = si
		s.indexByName[key.Name] = append(s.indexByName[key.Name], si)
	}

	return nil
}

// UnmarshalJSON deserializes JSON data into the `Interfaces` field.
func (s *SchemaInterfaces) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &s.Interfaces); err != nil {
		return err
	}

	return s.buildIndex()
}

// MarshalJSON serializes the `Interfaces` field to JSON data.
func (s SchemaInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Interfaces)
}

// Find returns the interface with given name and appID. It returns nil if none
// was found.
func (s *SchemaInterfaces) Find(name string, appID uint32) *SchemaInterface {
	return s.index.Find(name, appID)
}

// Get is like `Find` but returns an error of type `*InterfaceNotFoundError` if
// none was found.
func (s *SchemaInterfaces) Get(name string, appID uint32) (si *SchemaInterface, err error) {
	si = s.index.Find(name, appID)

	if si == nil {
		err = &InterfaceNotFoundError{Name: name, AppID: appID}
	}

	return
}

// GroupByName groups the interfaces by name, returning a map of `{ name => sub-collection }`.
//
// Each value in the map is a grouped `SchemaInterfaces` and can use group helpers.
//
// Returns errors described in `NewSchemaInterfaces`.
func (s *SchemaInterfaces) GroupByName() (map[string]*SchemaInterfaces, error) {
	var err error
	result := make(map[string]*SchemaInterfaces)

	for name, interfaces := range s.indexByName {
		if result[name], err = NewSchemaInterfaces(interfaces...); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// IsGroup checks if the collection contains only interfaces with the same name.
func (s *SchemaInterfaces) IsGroup() bool {
	return len(s.indexByName) == 1
}

// GroupName returns the common name between all contained interfaces.
//
// If the collection is empty, it returns error `ErrEmptyInterfaces`.
//
// If the collection is not a group, it returns error `ErrMixedInterfaces`.
func (s *SchemaInterfaces) GroupName() (string, error) {
	if len(s.Interfaces) == 0 {
		return "", ErrEmptyInterfaces
	}

	if !s.IsGroup() {
		return "", ErrMixedInterfaces
	}

	return s.Interfaces[0].BaseName()
}

// AppIDs collects the AppID of all interfaces in the collection.
//
// If the collection is not a group, it returns error `ErrMixedInterfaces`.
func (s *SchemaInterfaces) AppIDs() ([]uint32, error) {
	if !s.IsGroup() {
		return nil, ErrMixedInterfaces
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

// GroupMethods groups interfaces methods by interface name.
//
// If the collection is not a group, it returns error `ErrMixedInterfaces`.
//
// Returns errors described in `SchemaMethods.GroupByName` and
// `SchemaMethods.Add`.
func (s *SchemaInterfaces) GroupMethods() (map[string]*SchemaMethods, error) {
	if len(s.Interfaces) == 0 {
		return nil, nil
	}

	if !s.IsGroup() {
		return nil, ErrMixedInterfaces
	}

	var err error
	result := map[string]*SchemaMethods{}

	for _, si := range s.Interfaces {
		var grouped map[string]*SchemaMethods
		grouped, err = si.Methods.GroupByName()

		if err != nil {
			return nil, err
		}

		for methodName, methods := range grouped {
			if result[methodName] == nil {
				result[methodName] = methods
			} else {
				if result[methodName], err = result[methodName].Add(methods); err != nil {
					return nil, err
				}
			}
		}
	}

	return result, nil
}
