// Code generated by geyser. DO NOT EDIT.
// API interface: ITFItems.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaITFItems stores the Interfaces for interface ITFItems.
var SchemaITFItems = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetGoldenWrenches",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetGoldenWrenches",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      2,
			},
		),
		Name:         "ITFItems_440",
		Undocumented: false,
	},
)

// ITFItems represents interface ITFItems.
//
// Supported AppIDs: 440.
type ITFItems struct {
	Client    *Client
	Interface *schema.Interface
}

// NewITFItems creates a new ITFItems interface.
//
// Supported AppIDs: 440.
func NewITFItems(c *Client, appID uint32) (*ITFItems, error) {
	si, err := SchemaITFItems.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "ITFItems",
	})

	if err != nil {
		return nil, err
	}

	s := &ITFItems{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ITFItems creates a new ITFItems interface.
//
// Supported AppIDs: 440.
func (c *Client) ITFItems(appID uint32) (*ITFItems, error) {
	return NewITFItems(c, appID)
}

/*
GetGoldenWrenches creates a Request for interface method GetGoldenWrenches.

Supported versions: 1, 2.
*/
func (i *ITFItems) GetGoldenWrenches(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGoldenWrenches",
		Version: version,
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