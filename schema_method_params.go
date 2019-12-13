package geyser

import (
	"encoding/json"
	"net/url"
)

// SchemaMethodParams is a collection of `SchemaMethodParam`.
//
// The struct should be read-only.
type SchemaMethodParams struct {
	Params []*SchemaMethodParam
}

// NewSchemaMethodParams creates a new collection.
func NewSchemaMethodParams(params ...*SchemaMethodParam) *SchemaMethodParams {
	return &SchemaMethodParams{Params: params}
}

// UnmarshalJSON deserializes JSON data into the `Params` field.
func (s *SchemaMethodParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &s.Params)
}

// MarshalJSON serializes the `Params` field to JSON data.
func (s SchemaMethodParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Params)
}

// Validate validates the given parameters against each parameter in the
// collection.
//
// Returns errors described in `SchemaMethodParam.Validate`.
func (s *SchemaMethodParams) Validate(params url.Values) error {
	for _, p := range s.Params {
		if err := p.Validate(params.Get(p.Name)); err != nil {
			return err
		}
	}

	return nil
}
