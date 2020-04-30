// Code generated by geyser. DO NOT EDIT.
// API interface: ITFSystem.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaITFSystem stores the Interfaces for interface ITFSystem.
var SchemaITFSystem = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetWorldStatus",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ITFSystem_440",
		Undocumented: false,
	},
)

// ITFSystem represents interface ITFSystem.
//
// Supported AppIDs: 440.
type ITFSystem struct {
	Client    *Client
	Interface *schema.Interface
}

// NewITFSystem creates a new ITFSystem interface.
//
// Supported AppIDs: 440.
func NewITFSystem(c *Client, appID uint32) (*ITFSystem, error) {
	si, err := SchemaITFSystem.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "ITFSystem",
	})

	if err != nil {
		return nil, err
	}

	s := &ITFSystem{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ITFSystem creates a new ITFSystem interface.
//
// Supported AppIDs: 440.
func (c *Client) ITFSystem(appID uint32) (*ITFSystem, error) {
	return NewITFSystem(c, appID)
}

// GetWorldStatus creates a Request for interface method GetWorldStatus.
func (i *ITFSystem) GetWorldStatus() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetWorldStatus",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
