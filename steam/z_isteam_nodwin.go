// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamNodwin.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamNodwin stores the Interfaces for interface ISteamNodwin.
var SchemaISteamNodwin = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "NodwinPaymentNotification",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamNodwin",
		Undocumented: true,
	},
)

// ISteamNodwin represents interface ISteamNodwin.
//
// This is an undocumented interface.
type ISteamNodwin struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamNodwin creates a new ISteamNodwin interface.
func NewISteamNodwin(c *Client) (*ISteamNodwin, error) {
	si, err := SchemaISteamNodwin.Get(schema.InterfaceKey{Name: "ISteamNodwin"})

	if err != nil {
		return nil, err
	}

	s := &ISteamNodwin{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamNodwin creates a new ISteamNodwin interface.
func (c *Client) ISteamNodwin() (*ISteamNodwin, error) {
	return NewISteamNodwin(c)
}

/*
NodwinPaymentNotification creates a Request for interface method NodwinPaymentNotification.

This is an undocumented method.
*/
func (i *ISteamNodwin) NodwinPaymentNotification() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "NodwinPaymentNotification",
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