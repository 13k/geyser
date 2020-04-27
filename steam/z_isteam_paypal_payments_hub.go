// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamPayPalPaymentsHub.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamPayPalPaymentsHub stores the Interfaces for interface ISteamPayPalPaymentsHub.
var SchemaISteamPayPalPaymentsHub = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "PayPalPaymentsHubPaymentNotification",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamPayPalPaymentsHub",
		Undocumented: true,
	},
)

// ISteamPayPalPaymentsHub represents interface ISteamPayPalPaymentsHub.
//
// This is an undocumented interface.
type ISteamPayPalPaymentsHub struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamPayPalPaymentsHub creates a new ISteamPayPalPaymentsHub interface.
func NewISteamPayPalPaymentsHub(c *Client) (*ISteamPayPalPaymentsHub, error) {
	si, err := SchemaISteamPayPalPaymentsHub.Get(schema.InterfaceKey{Name: "ISteamPayPalPaymentsHub"})

	if err != nil {
		return nil, err
	}

	s := &ISteamPayPalPaymentsHub{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamPayPalPaymentsHub creates a new ISteamPayPalPaymentsHub interface.
func (c *Client) ISteamPayPalPaymentsHub() (*ISteamPayPalPaymentsHub, error) {
	return NewISteamPayPalPaymentsHub(c)
}

/*
PayPalPaymentsHubPaymentNotification creates a Request for interface method PayPalPaymentsHubPaymentNotification.

This is an undocumented method.
*/
func (i *ISteamPayPalPaymentsHub) PayPalPaymentsHubPaymentNotification() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "PayPalPaymentsHubPaymentNotification",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
