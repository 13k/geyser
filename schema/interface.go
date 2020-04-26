package schema

import (
	"regexp"
	"strconv"
)

var (
	siNameRegexp      = regexp.MustCompile(`^[A-Z]\w*$`)
	siNameAppIDRegexp = regexp.MustCompile(`^(\w+)_(\d+)$`)
)

// InterfaceKey is the key that uniquely identifies a `Interface`.
type InterfaceKey struct {
	Name  string
	AppID uint32
}

// Interface holds the specification of an API interface.
//
// The struct should be read-only.
type Interface struct {
	Name         string  `json:"name"`
	Methods      Methods `json:"methods"`
	Undocumented bool    `json:"undocumented"`

	key InterfaceKey
}

func (i *Interface) parse() error {
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
// Returns errors described in `Methods.Validate`.
func (i *Interface) Validate() error {
	if err := i.parse(); err != nil {
		return err
	}

	return i.Methods.Validate()
}

// Key returns the key identifying the interface.
//
// Interface names are formed by a base name and an optional AppID part, in the format
// `<BaseName>_<AppID>`. Interfaces can be non app-specific and omit the "_<AppID>" part, in these
// cases Key returns a key with AppID = 0.
//
// For example in "IDOTAMatch_570", the base name is "IDOTAMatch" and AppID is 570. In "ISteamApps",
// the base name is "ISteamApps" and AppID is 0.
//
// Key parses the interface name and extracts the base name and the AppID from the name.
//
// Returns errors described in `Validate`.
func (i *Interface) Key() (InterfaceKey, error) {
	err := i.parse()
	return i.key, err
}
