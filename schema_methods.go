package geyser

// SchemaMethods is a collection of `SchemaMethod`.
type SchemaMethods []*SchemaMethod

// MustNewSchemaMethods is like `NewSchemaMethods` but panics if it returned an error.
func MustNewSchemaMethods(methods ...*SchemaMethod) SchemaMethods {
	c, err := NewSchemaMethods(methods...)

	if err != nil {
		panic(err)
	}

	return c
}

// NewSchemaMethods creates a new collection.
//
// Returns errors described in `SchemaMethod.Validate`.
func NewSchemaMethods(methods ...*SchemaMethod) (SchemaMethods, error) {
	c := SchemaMethods(methods)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Validate checks if all contained methods are valid.
//
// Returns errors described in `SchemaMethod.Validate`.
func (c SchemaMethods) Validate() error {
	for _, sm := range c {
		if err := sm.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the method with the given key.
//
// Returns an error of type `*InterfaceMethodNotFoundError` if none was found.
//
// Returns errors described in `SchemaMethod.Key`.
func (c SchemaMethods) Get(key SchemaMethodKey) (*SchemaMethod, error) {
	for _, sm := range c {
		k, err := sm.Key()

		if err != nil {
			return nil, err
		}

		if k == key {
			return sm, nil
		}
	}

	return nil, &InterfaceMethodNotFoundError{Key: key}
}

// GroupByName groups the methods by name.
//
// Returns errors described in `SchemaMethod.Key`.
func (c SchemaMethods) GroupByName() (map[string]SchemaMethodsGroup, error) {
	result := make(map[string]SchemaMethodsGroup)

	for _, sm := range c {
		key, err := sm.Key()

		if err != nil {
			return nil, err
		}

		if result[key.Name] == nil {
			result[key.Name] = make(SchemaMethodsGroup)
		}

		result[key.Name][key] = sm
	}

	return result, nil
}

/*
SchemaMethodsGroup is a group of `SchemaMethod`s with the same name.

It's a regular map and therefore provides no guarantees on consistency:

* Keys are not guaranteed to be correctly associated to their respective methods
* Methods are not guaranteed to be unique for each key
* Methods are not guaranteed to have the same name

The group creator is responsible for ensuring consistency. Group consumers can assume it's
consistent.

Behavior of inconsistent groups is undefined.
*/
type SchemaMethodsGroup map[SchemaMethodKey]*SchemaMethod

// Versions collects the versions of all methods in the group.
func (g SchemaMethodsGroup) Versions() []int {
	var versions []int

	for key := range g {
		versions = append(versions, key.Version)
	}

	return versions
}
