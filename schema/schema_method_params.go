package schema

import (
	"net/url"
)

// SchemaMethodParams is a collection of `SchemaMethodParam`.
type SchemaMethodParams []*SchemaMethodParam

// NewSchemaMethodParams creates a new collection.
func NewSchemaMethodParams(params ...*SchemaMethodParam) SchemaMethodParams {
	return SchemaMethodParams(params)
}

// ValidateParams validates the given parameters against each parameter in the collection.
//
// Returns errors described in `SchemaMethodParam.ValidateParam`.
func (c SchemaMethodParams) ValidateParams(params url.Values) error {
	for _, p := range c {
		if err := p.ValidateValue(params.Get(p.Name)); err != nil {
			return err
		}
	}

	return nil
}
