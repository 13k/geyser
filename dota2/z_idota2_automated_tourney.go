// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2AutomatedTourney.

package dota2

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIDOTA2AutomatedTourney stores the Interfaces for interface IDOTA2AutomatedTourney.
var SchemaIDOTA2AutomatedTourney = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetParticipationDetails",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetPlayerHistory",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetTournamentDetails",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IDOTA2AutomatedTourney",
		Undocumented: true,
	},
)

// IDOTA2AutomatedTourney represents interface IDOTA2AutomatedTourney.
//
// This is an undocumented interface.
type IDOTA2AutomatedTourney struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2AutomatedTourney creates a new IDOTA2AutomatedTourney interface.
func NewIDOTA2AutomatedTourney(c *Client) (*IDOTA2AutomatedTourney, error) {
	si, err := SchemaIDOTA2AutomatedTourney.Get(schema.InterfaceKey{Name: "IDOTA2AutomatedTourney"})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2AutomatedTourney{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2AutomatedTourney creates a new IDOTA2AutomatedTourney interface.
func (c *Client) IDOTA2AutomatedTourney() (*IDOTA2AutomatedTourney, error) {
	return NewIDOTA2AutomatedTourney(c)
}

/*
GetParticipationDetails creates a Request for interface method GetParticipationDetails.

This is an undocumented method.
*/
func (i *IDOTA2AutomatedTourney) GetParticipationDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetParticipationDetails",
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
GetPlayerHistory creates a Request for interface method GetPlayerHistory.

This is an undocumented method.
*/
func (i *IDOTA2AutomatedTourney) GetPlayerHistory() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetPlayerHistory",
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
GetTournamentDetails creates a Request for interface method GetTournamentDetails.

This is an undocumented method.
*/
func (i *IDOTA2AutomatedTourney) GetTournamentDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetTournamentDetails",
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
