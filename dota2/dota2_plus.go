// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IDOTA2Plus.

package dota2

import (
	"github.com/13k/geyser"
	"net/http"
)

// SchemaDOTA2Plus stores the SchemaInterfaces for interface IDOTA2Plus.
var SchemaDOTA2Plus = geyser.MustNewSchemaInterfaces(
	&geyser.SchemaInterface{
		Methods: geyser.MustNewSchemaMethods(
			&geyser.SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetPlusStatsData",
				Params:       geyser.NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IDOTA2Plus",
		Undocumented: true,
	},
)

// DOTA2Plus represents interface IDOTA2Plus.
//
// This is an undocumented interface.
type DOTA2Plus struct {
	Client    *Client
	Interface *geyser.SchemaInterface
}

// NewDOTA2Plus creates a new DOTA2Plus interface.
func NewDOTA2Plus(c *Client) (*DOTA2Plus, error) {
	si, err := SchemaDOTA2Plus.Get(geyser.SchemaInterfaceKey{Name: "IDOTA2Plus"})

	if err != nil {
		return nil, err
	}

	s := &DOTA2Plus{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// DOTA2Plus creates a new DOTA2Plus interface.
func (c *Client) DOTA2Plus() (*DOTA2Plus, error) {
	return NewDOTA2Plus(c)
}

/*
GetPlusStatsData creates a Request for interface method GetPlusStatsData.

This is an undocumented method.
*/
func (i *DOTA2Plus) GetPlusStatsData() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(geyser.SchemaMethodKey{
		Name:    "GetPlusStatsData",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &DOTA2PlusGetPlusStatsData{},
	}

	return req, nil
}