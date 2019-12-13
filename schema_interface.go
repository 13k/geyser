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
	Name         string         `json:"name"`
	Methods      *SchemaMethods `json:"methods"`
	Undocumented bool           `json:"undocumented"`

	appID    uint32
	baseName string
	key      *SchemaInterfaceKey
}

func (i *SchemaInterface) parse() error {
	if i.key != nil {
		return nil
	}

	if !siNameRegexp.MatchString(i.Name) {
		return &InvalidInterfaceNameError{Interface: i}
	}

	if matches := siNameAppIDRegexp.FindStringSubmatch(i.Name); matches != nil {
		appID64, err := strconv.ParseUint(matches[2], 10, 32)

		if err != nil {
			return &InvalidInterfaceNameError{Interface: i, Err: err}
		}

		i.baseName = matches[1]
		i.appID = uint32(appID64)
	} else {
		i.baseName = i.Name
	}

	i.key = &SchemaInterfaceKey{Name: i.baseName, AppID: i.appID}

	return nil
}

// Key returns the key identifying the interface.
//
// If the interface name is invalid, it returns an error of type
// `*InvalidInterfaceNameError`.
func (i *SchemaInterface) Key() (key *SchemaInterfaceKey, err error) {
	err = i.parse()
	key = i.key
	return
}

// GetMethods returns the underlying slice of `SchemaMethod`.
func (i *SchemaInterface) GetMethods() []*SchemaMethod {
	if i.Methods == nil {
		return nil
	}

	return i.Methods.Methods
}

// BaseName extracts the base name of the interface (without the AppID part).
//
// Returns errors described in `Key`.
func (i *SchemaInterface) BaseName() (s string, err error) {
	err = i.parse()
	s = i.baseName
	return
}

// AppID extracts the AppID of the interface (without the BaseName part).
//
// Returns errors described in `Key`.
func (i *SchemaInterface) AppID() (id uint32, err error) {
	err = i.parse()
	id = i.appID
	return
}
