package geyser

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

// SchemaMethodKey is the key that uniquely identifies a `SchemaMethod`.
type SchemaMethodKey struct {
	Name    string
	Version int
}

// SchemaMethod holds the specification of an API interface method.
//
// The struct should be read-only.
type SchemaMethod struct {
	Name         string              `json:"name"`
	Version      int                 `json:"version"`
	HTTPMethod   string              `json:"httpmethod"`
	Params       *SchemaMethodParams `json:"parameters"`
	Undocumented bool                `json:"undocumented"`

	key     *SchemaMethodKey
	version string
}

func (m *SchemaMethod) versionPathParam() string {
	if m.version == "" {
		m.version = fmt.Sprintf("v%d", m.Version)
	}

	return m.version
}

func (m *SchemaMethod) parse() error {
	if m.key != nil {
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

	m.key = &SchemaMethodKey{Name: m.Name, Version: m.Version}

	return nil
}

// Key returns the key identifying the method.
//
// If the method has an invalid name, an error of type `*InvalidMethodNameError`
// is returned.
//
// If the method has an invalid version, an error of type
// `*InvalidMethodVersionError` is returned.
//
// If the method has an invalid http method, an error of type
// `*InvalidMethodHTTPMethodError` is returned.
func (m *SchemaMethod) Key() (key *SchemaMethodKey, err error) {
	err = m.parse()
	key = m.key
	return
}

// GetParams returns the underlying slice of `SchemaMethodParam`.
func (m *SchemaMethod) GetParams() []*SchemaMethodParam {
	if m.Params == nil {
		return nil
	}

	return m.Params.Params
}

// ValidateParams validates the given parameters against the params collection.
func (m *SchemaMethod) ValidateParams(params url.Values) error {
	return m.Params.Validate(params)
}
