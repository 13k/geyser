// Code generated by geyser. DO NOT EDIT.
// API interface: ICredentialsService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaICredentialsService stores the Interfaces for interface ICredentialsService.
var SchemaICredentialsService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "ValidateEmailAddress",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "SteamGuardPhishingReport",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ICredentialsService",
		Undocumented: true,
	},
)

// ICredentialsService represents interface ICredentialsService.
//
// This is an undocumented interface.
type ICredentialsService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewICredentialsService creates a new ICredentialsService interface.
func NewICredentialsService(c *Client) (*ICredentialsService, error) {
	si, err := SchemaICredentialsService.Get(schema.InterfaceKey{Name: "ICredentialsService"})

	if err != nil {
		return nil, err
	}

	s := &ICredentialsService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ICredentialsService creates a new ICredentialsService interface.
func (c *Client) ICredentialsService() (*ICredentialsService, error) {
	return NewICredentialsService(c)
}

/*
SteamGuardPhishingReport creates a Request for interface method SteamGuardPhishingReport.

This is an undocumented method.
*/
func (i *ICredentialsService) SteamGuardPhishingReport() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "SteamGuardPhishingReport",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
ValidateEmailAddress creates a Request for interface method ValidateEmailAddress.

This is an undocumented method.
*/
func (i *ICredentialsService) ValidateEmailAddress() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "ValidateEmailAddress",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
