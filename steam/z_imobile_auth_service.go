// Code generated by geyser. DO NOT EDIT.
// API interface: IMobileAuthService.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIMobileAuthService stores the Interfaces for interface IMobileAuthService.
var SchemaIMobileAuthService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetWGToken",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IMobileAuthService",
		Undocumented: true,
	},
)

// IMobileAuthService represents interface IMobileAuthService.
//
// This is an undocumented interface.
type IMobileAuthService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIMobileAuthService creates a new IMobileAuthService interface.
func NewIMobileAuthService(c *Client) (*IMobileAuthService, error) {
	si, err := SchemaIMobileAuthService.Get(schema.InterfaceKey{Name: "IMobileAuthService"})

	if err != nil {
		return nil, err
	}

	s := &IMobileAuthService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IMobileAuthService creates a new IMobileAuthService interface.
func (c *Client) IMobileAuthService() (*IMobileAuthService, error) {
	return NewIMobileAuthService(c)
}

/*
GetWGToken creates a Request for interface method GetWGToken.

This is an undocumented method.
*/
func (i *IMobileAuthService) GetWGToken() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetWGToken",
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