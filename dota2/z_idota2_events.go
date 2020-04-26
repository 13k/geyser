// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2Events.

package dota2

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIDOTA2Events stores the Interfaces for interface IDOTA2Events.
var SchemaIDOTA2Events = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetArcanaVotes",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetCurrentTriviaQuestions",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetMutations",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetTriviaQuestionAnswersSummary",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IDOTA2Events",
		Undocumented: true,
	},
)

// IDOTA2Events represents interface IDOTA2Events.
//
// This is an undocumented interface.
type IDOTA2Events struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2Events creates a new IDOTA2Events interface.
func NewIDOTA2Events(c *Client) (*IDOTA2Events, error) {
	si, err := SchemaIDOTA2Events.Get(schema.InterfaceKey{Name: "IDOTA2Events"})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2Events{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2Events creates a new IDOTA2Events interface.
func (c *Client) IDOTA2Events() (*IDOTA2Events, error) {
	return NewIDOTA2Events(c)
}

/*
GetArcanaVotes creates a Request for interface method GetArcanaVotes.

This is an undocumented method.
*/
func (i *IDOTA2Events) GetArcanaVotes() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetArcanaVotes",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetCurrentTriviaQuestions creates a Request for interface method GetCurrentTriviaQuestions.

This is an undocumented method.
*/
func (i *IDOTA2Events) GetCurrentTriviaQuestions() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCurrentTriviaQuestions",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetMutations creates a Request for interface method GetMutations.

This is an undocumented method.
*/
func (i *IDOTA2Events) GetMutations() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetMutations",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetTriviaQuestionAnswersSummary creates a Request for interface method GetTriviaQuestionAnswersSummary.

This is an undocumented method.
*/
func (i *IDOTA2Events) GetTriviaQuestionAnswersSummary() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetTriviaQuestionAnswersSummary",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}