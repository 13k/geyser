package geyser

// SchemaInterfaces is a collection of `SchemaInterface`.
type SchemaInterfaces []*SchemaInterface

// MustNewSchemaInterfaces is like `NewSchemaInterfaces` but panics if it returned an error.
func MustNewSchemaInterfaces(interfaces ...*SchemaInterface) SchemaInterfaces {
	c, err := NewSchemaInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return c
}

// NewSchemaInterfaces creates a new collection.
//
// Returns errors described in `SchemaInterface.Validate`.
func NewSchemaInterfaces(interfaces ...*SchemaInterface) (SchemaInterfaces, error) {
	c := SchemaInterfaces(interfaces)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Validate checks if all contained interfaces are valid.
//
// Returns errors described in `SchemaInterface.Validate`.
func (c SchemaInterfaces) Validate() error {
	for _, si := range c {
		if err := si.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Get returns the interface with given key.
//
// Returns an error of type `*InterfaceNotFoundError` if none was found.
//
// Returns errors described in `SchemaInterface.Key`.
func (c SchemaInterfaces) Get(key SchemaInterfaceKey) (*SchemaInterface, error) {
	for _, si := range c {
		k, err := si.Key()

		if err != nil {
			return nil, err
		}

		if k == key {
			return si, nil
		}
	}

	return nil, &InterfaceNotFoundError{Key: key}
}

// GroupByName groups the interfaces by base name.
//
// The definition of base name is described in `SchemaInterface.Key`.
//
// Returns errors described in `SchemaInterface.Key`.
func (c SchemaInterfaces) GroupByBaseName() (map[string]SchemaInterfacesGroup, error) {
	result := make(map[string]SchemaInterfacesGroup)

	for _, si := range c {
		key, err := si.Key()

		if err != nil {
			return nil, err
		}

		if result[key.Name] == nil {
			result[key.Name] = make(SchemaInterfacesGroup)
		}

		result[key.Name][key] = si
	}

	return result, nil
}

/*
SchemaInterfacesGroup is a group of `SchemaInterface`s with the same base name.

The definition of base name is described in `SchemaInterface.Key`.

It's a regular map and therefore provides no guarantees on consistency:

* Keys are not guaranteed to be correctly associated to their respective interfaces
* Interfaces are not guaranteed to be unique for each key
* Interfaces are not guaranteed to have the same base name

The group creator is responsible for ensuring consistency. Group consumers can assume it's
consistent.

Behavior of inconsistent groups is undefined.
*/
type SchemaInterfacesGroup map[SchemaInterfaceKey]*SchemaInterface

// Name returns the common base name of all interfaces in the group.
//
// The definition of base name is described in `SchemaInterface.Key`.
func (g SchemaInterfacesGroup) Name() (name string) {
	for key := range g {
		name = key.Name
		break
	}

	return
}

// AppIDs collects the AppID of all interfaces in the group.
//
// Interfaces with no AppID (0) are omitted.
func (g SchemaInterfacesGroup) AppIDs() []uint32 {
	var appIDs []uint32

	for key := range g {
		if key.AppID != 0 {
			appIDs = append(appIDs, key.AppID)
		}
	}

	return appIDs
}

// GroupMethods groups interfaces methods by method name.
//
// Returns errors described in `SchemaMethods.GroupByName`.
func (g SchemaInterfacesGroup) GroupMethods() (map[string]SchemaMethodsGroup, error) {
	var result map[string]SchemaMethodsGroup

	for _, si := range g {
		grouped, err := si.Methods.GroupByName()

		if err != nil {
			return nil, err
		}

		if result == nil {
			result = grouped
			continue
		}

		for name, group := range grouped {
			if result[name] == nil {
				result[name] = group
				continue
			}

			for key, sm := range group {
				result[name][key] = sm
			}
		}
	}

	return result, nil
}
