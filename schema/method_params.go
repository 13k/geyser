package schema

import (
	"net/url"
)

// MethodParams is a collection of `MethodParam`s.
type MethodParams []*MethodParam

// NewMethodParams creates a new collection.
func NewMethodParams(params ...*MethodParam) MethodParams {
	return MethodParams(params)
}

// ValidateParams validates the given parameters against each parameter in the collection.
//
// Returns errors described in `MethodParam.ValidateParam`.
func (c MethodParams) ValidateParams(params url.Values) error {
	for _, p := range c {
		if err := p.ValidateValue(params.Get(p.Name)); err != nil {
			return err
		}
	}

	return nil
}
