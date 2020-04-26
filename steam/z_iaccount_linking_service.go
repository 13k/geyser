// Code generated by geyser. DO NOT EDIT.
// API interface: IAccountLinkingService.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIAccountLinkingService stores the Interfaces for interface IAccountLinkingService.
var SchemaIAccountLinkingService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetLinkedAccountInfo",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IAccountLinkingService",
		Undocumented: true,
	},
)

// IAccountLinkingService represents interface IAccountLinkingService.
//
// This is an undocumented interface.
type IAccountLinkingService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIAccountLinkingService creates a new IAccountLinkingService interface.
func NewIAccountLinkingService(c *Client) (*IAccountLinkingService, error) {
	si, err := SchemaIAccountLinkingService.Get(schema.InterfaceKey{Name: "IAccountLinkingService"})

	if err != nil {
		return nil, err
	}

	s := &IAccountLinkingService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IAccountLinkingService creates a new IAccountLinkingService interface.
func (c *Client) IAccountLinkingService() (*IAccountLinkingService, error) {
	return NewIAccountLinkingService(c)
}

/*
GetLinkedAccountInfo creates a Request for interface method GetLinkedAccountInfo.

This is an undocumented method.
*/
func (i *IAccountLinkingService) GetLinkedAccountInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetLinkedAccountInfo",
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