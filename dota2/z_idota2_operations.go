// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2Operations.

package dota2

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIDOTA2Operations stores the Interfaces for interface IDOTA2Operations.
var SchemaIDOTA2Operations = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBannedWordList",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      2,
			},
		),
		Name:         "IDOTA2Operations",
		Undocumented: true,
	},
)

// IDOTA2Operations represents interface IDOTA2Operations.
//
// This is an undocumented interface.
type IDOTA2Operations struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2Operations creates a new IDOTA2Operations interface.
func NewIDOTA2Operations(c *Client) (*IDOTA2Operations, error) {
	si, err := SchemaIDOTA2Operations.Get(schema.InterfaceKey{Name: "IDOTA2Operations"})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2Operations{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2Operations creates a new IDOTA2Operations interface.
func (c *Client) IDOTA2Operations() (*IDOTA2Operations, error) {
	return NewIDOTA2Operations(c)
}

/*
GetBannedWordList creates a Request for interface method GetBannedWordList.

This is an undocumented method.
*/
func (i *IDOTA2Operations) GetBannedWordList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBannedWordList",
		Version: 2,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
