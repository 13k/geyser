package schema

// Interfaces is a collection of `Interface`s.
type Interfaces []*Interface

// MustNewInterfaces is like `NewInterfaces` but panics if it returned an error.
func MustNewInterfaces(interfaces ...*Interface) Interfaces {
	c, err := NewInterfaces(interfaces...)

	if err != nil {
		panic(err)
	}

	return c
}

// NewInterfaces creates a new collection.
//
// Returns errors described in `Interface.Validate`.
func NewInterfaces(interfaces ...*Interface) (Interfaces, error) {
	c := Interfaces(interfaces)

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return c, nil
}

// Validate checks if all contained interfaces are valid.
//
// Returns errors described in `Interface.Validate`.
func (c Interfaces) Validate() error {
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
// Returns errors described in `Interface.Key`.
func (c Interfaces) Get(key InterfaceKey) (*Interface, error) {
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

// GroupByBaseName groups the interfaces by base name.
//
// The definition of base name is described in `Interface.Key`.
//
// Returns errors described in `Interface.Key`.
func (c Interfaces) GroupByBaseName() (map[string]InterfacesGroup, error) {
	result := make(map[string]InterfacesGroup)

	for _, si := range c {
		key, err := si.Key()

		if err != nil {
			return nil, err
		}

		if result[key.Name] == nil {
			result[key.Name] = make(InterfacesGroup)
		}

		result[key.Name][key] = si
	}

	return result, nil
}

/*
InterfacesGroup is a group of `Interface`s with the same base name.

The definition of base name is described in `Interface.Key`.

It's a regular map and therefore provides no guarantees on consistency:

* Keys are not guaranteed to be correctly associated to their respective interfaces

* Interfaces are not guaranteed to be unique for each key

* Interfaces are not guaranteed to have the same base name

The group creator is responsible for ensuring consistency. Group consumers can assume it's
consistent.

Behavior of inconsistent groups is undefined.
*/
type InterfacesGroup map[InterfaceKey]*Interface

// Name returns the common base name of all interfaces in the group.
//
// The definition of base name is described in `Interface.Key`.
func (g InterfacesGroup) Name() (name string) {
	for key := range g {
		name = key.Name
		break
	}

	return
}

// AppIDs collects the AppID of all interfaces in the group.
//
// Interfaces with no AppID (0) are omitted.
func (g InterfacesGroup) AppIDs() []uint32 {
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
// Returns errors described in `Methods.GroupByName`.
func (g InterfacesGroup) GroupMethods() (map[string]MethodsGroup, error) {
	var result map[string]MethodsGroup

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
