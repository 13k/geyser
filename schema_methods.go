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
func (c SchemaMethods) Get(key SchemaMethodKey) (*SchemaMethod, error) {
	for _, sm := range c {
		if k, err := sm.Key(); err != nil {
			return nil, err
		} else {
			if k == key {
				return sm, nil
			}
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

// SchemaMethodsGroup is a group of `SchemaMethod`s with the same name.
type SchemaMethodsGroup map[SchemaMethodKey]*SchemaMethod

// Versions collects the unique versions of all methods in the group.
func (g SchemaMethodsGroup) Versions() []int {
	var versions []int

	visited := make(map[int]bool)

	for key := range g {
		if !visited[key.Version] {
			versions = append(versions, key.Version)
		}
	}

	return versions
}
