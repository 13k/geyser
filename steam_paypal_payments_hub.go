// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ISteamPayPalPaymentsHub.

package geyser

import "net/http"

// SchemaSteamPayPalPaymentsHub stores the SchemaInterfaces for interface ISteamPayPalPaymentsHub.
var SchemaSteamPayPalPaymentsHub = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: MustNewSchemaMethods(
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "PayPalPaymentsHubPaymentNotification",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamPayPalPaymentsHub",
		Undocumented: true,
	},
)

// SteamPayPalPaymentsHub represents interface ISteamPayPalPaymentsHub.
//
// This is an undocumented interface.
type SteamPayPalPaymentsHub struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewSteamPayPalPaymentsHub creates a new SteamPayPalPaymentsHub interface.
func NewSteamPayPalPaymentsHub(c *Client) (*SteamPayPalPaymentsHub, error) {
	si, err := SchemaSteamPayPalPaymentsHub.Get(SchemaInterfaceKey{Name: "ISteamPayPalPaymentsHub"})

	if err != nil {
		return nil, err
	}

	s := &SteamPayPalPaymentsHub{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// SteamPayPalPaymentsHub creates a new SteamPayPalPaymentsHub interface.
func (c *Client) SteamPayPalPaymentsHub() (*SteamPayPalPaymentsHub, error) {
	return NewSteamPayPalPaymentsHub(c)
}

/*
PayPalPaymentsHubPaymentNotification creates a Request for interface method PayPalPaymentsHubPaymentNotification.

This is an undocumented method.
*/
func (i *SteamPayPalPaymentsHub) PayPalPaymentsHubPaymentNotification() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "PayPalPaymentsHubPaymentNotification",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamPayPalPaymentsHubPayPalPaymentsHubPaymentNotification{},
	}

	return req, nil
}