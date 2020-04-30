// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2StreamSystem.

package dota2

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIDOTA2StreamSystem stores the Interfaces for interface IDOTA2StreamSystem.
var SchemaIDOTA2StreamSystem = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBroadcasterInfo",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IDOTA2StreamSystem",
		Undocumented: true,
	},
)

// IDOTA2StreamSystem represents interface IDOTA2StreamSystem.
//
// This is an undocumented interface.
type IDOTA2StreamSystem struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2StreamSystem creates a new IDOTA2StreamSystem interface.
func NewIDOTA2StreamSystem(c *Client) (*IDOTA2StreamSystem, error) {
	si, err := SchemaIDOTA2StreamSystem.Get(schema.InterfaceKey{Name: "IDOTA2StreamSystem"})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2StreamSystem{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2StreamSystem creates a new IDOTA2StreamSystem interface.
func (c *Client) IDOTA2StreamSystem() (*IDOTA2StreamSystem, error) {
	return NewIDOTA2StreamSystem(c)
}

/*
GetBroadcasterInfo creates a Request for interface method GetBroadcasterInfo.

This is an undocumented method.
*/
func (i *IDOTA2StreamSystem) GetBroadcasterInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBroadcasterInfo",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
