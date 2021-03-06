package schema

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var (
	smNameRegexp       = regexp.MustCompile(`^[[:alpha:]]\w*$`)
	smValidHTTPMethods = map[string]bool{
		http.MethodGet:  true,
		http.MethodPut:  true,
		http.MethodPost: true,
	}
)

// MethodKey is the key that uniquely identifies a `Method`.
type MethodKey struct {
	Name    string
	Version int
}

// Method holds the specification of an API interface method.
//
// The struct should be read-only.
type Method struct {
	Name         string       `json:"name"`
	Version      int          `json:"version"`
	HTTPMethod   string       `json:"httpmethod"`
	Params       MethodParams `json:"parameters"`
	Undocumented bool         `json:"undocumented"`

	key     MethodKey
	version string
}

// Validate checks if the method is valid.
//
// Returns an error of type `*InvalidMethodNameError` if the method has an invalid name.
//
// Returns an error of type `*InvalidMethodVersionError` if the method has an invalid version.
//
// Returns an error of type `*InvalidMethodHTTPMethodError` if the method has an invalid http
// method.
func (m *Method) Validate() error {
	return m.parse()
}

// Key returns the key identifying the method.
//
// Returns errors described in `Validate`.
func (m *Method) Key() (MethodKey, error) {
	err := m.parse()
	return m.key, err
}

// ValidateParams validates the given parameters against the params collection.
//
// Returns errors described in `MethodParams.ValidateParams`.
func (m *Method) ValidateParams(params url.Values) error {
	return m.Params.ValidateParams(params)
}

// FormatVersion returns the formatted version string (in the format "v<version>") to be used as a
// request path parameter.
func (m *Method) FormatVersion() string {
	if m.version == "" {
		m.version = fmt.Sprintf("v%d", m.Version)
	}

	return m.version
}

func (m *Method) parse() error {
	if m.key.Name != "" {
		return nil
	}

	if !smNameRegexp.MatchString(m.Name) {
		return &InvalidMethodNameError{Method: m}
	}

	if m.Version < 1 {
		return &InvalidMethodVersionError{Method: m}
	}

	if _, ok := smValidHTTPMethods[m.HTTPMethod]; !ok {
		return &InvalidMethodHTTPMethodError{Method: m}
	}

	m.key.Name = m.Name
	m.key.Version = m.Version

	return nil
}
