package schema

// Methods is a collection of `Method`s.
type Methods []*Method

// MustNewMethods is like `NewMethods` but panics if it returned an error.
func MustNewMethods(methods ...*Method) Methods {
	c, err := NewMethods(methods...)

	if err != nil {
		panic(err)
	}

	return c
}

// NewMethods creates a new collection.
//
// Returns errors described in `Method.Validate`.
func NewMethods(methods ...*Method) (Methods, error) {
	c := Methods(methods)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Validate checks if all contained methods are valid.
//
// Returns errors described in `Method.Validate`.
func (c Methods) Validate() error {
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
// Returns errors described in `Method.Key`.
func (c Methods) Get(key MethodKey) (*Method, error) {
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
// Returns errors described in `Method.Key`.
func (c Methods) GroupByName() (map[string]MethodsGroup, error) {
	result := make(map[string]MethodsGroup)

	for _, sm := range c {
		key, err := sm.Key()

		if err != nil {
			return nil, err
		}

		if result[key.Name] == nil {
			result[key.Name] = make(MethodsGroup)
		}

		result[key.Name][key] = sm
	}

	return result, nil
}

/*
MethodsGroup is a group of `Method`s with the same name.

It's a regular map and therefore provides no guarantees on consistency:

* Keys are not guaranteed to be correctly associated to their respective methods

* Methods are not guaranteed to be unique for each key

* Methods are not guaranteed to have the same name

The group creator is responsible for ensuring consistency. Group consumers can assume it's
consistent.

Behavior of inconsistent groups is undefined.
*/
type MethodsGroup map[MethodKey]*Method

// Versions collects the versions of all methods in the group.
func (g MethodsGroup) Versions() []int {
	versions := make([]int, 0, len(g))

	for key := range g {
		versions = append(versions, key.Version)
	}

	return versions
}
