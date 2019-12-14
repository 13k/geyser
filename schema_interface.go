package geyser

import (
	"regexp"
	"strconv"
)

var (
	siNameRegexp      = regexp.MustCompile(`^[A-Z]\w*$`)
	siNameAppIDRegexp = regexp.MustCompile(`^(\w+)_(\d+)$`)
)

// SchemaInterfaceKey is the key that uniquely identifies a `SchemaInterface`.
type SchemaInterfaceKey struct {
	Name  string
	AppID uint32
}

// SchemaInterface holds the specification of an API interface.
//
// The struct should be read-only.
type SchemaInterface struct {
	Name         string        `json:"name"`
	Methods      SchemaMethods `json:"methods"`
	Undocumented bool          `json:"undocumented"`

	key SchemaInterfaceKey
}

func (i *SchemaInterface) parse() error {
	if i.key.Name != "" {
		return nil
	}

	if !siNameRegexp.MatchString(i.Name) {
		return &InvalidInterfaceNameError{Interface: i}
	}

	var appID uint32
	baseName := i.Name

	if matches := siNameAppIDRegexp.FindStringSubmatch(i.Name); matches != nil {
		appID64, err := strconv.ParseUint(matches[2], 10, 32)

		if err != nil {
			return &InvalidInterfaceNameError{Interface: i, Err: err}
		}

		baseName = matches[1]
		appID = uint32(appID64)
	}

	i.key.Name = baseName
	i.key.AppID = appID

	return nil
}

// Validate checks if the interface and all its methods are valid.
//
// Returns an error of type `*InvalidInterfaceNameError` if the interface name is invalid.
//
// Returns errors described in `SchemaMethods.Validate`.
func (i *SchemaInterface) Validate() error {
	if err := i.parse(); err != nil {
		return err
	}

	return i.Methods.Validate()
}

// Key returns the key identifying the interface.
//
// Returns errors described in `Validate`.
func (i *SchemaInterface) Key() (SchemaInterfaceKey, error) {
	err := i.parse()
	return i.key, err
}
