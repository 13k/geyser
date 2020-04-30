// Code generated by geyser. DO NOT EDIT.
// API interface: IClientStats.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIClientStats stores the Interfaces for interface IClientStats.
var SchemaIClientStats = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodPost,
				Name:         "ReportEvent",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IClientStats_1046930",
		Undocumented: false,
	},
)

// IClientStats represents interface IClientStats.
//
// Supported AppIDs: 1046930.
type IClientStats struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIClientStats creates a new IClientStats interface.
//
// Supported AppIDs: 1046930.
func NewIClientStats(c *Client, appID uint32) (*IClientStats, error) {
	si, err := SchemaIClientStats.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IClientStats",
	})

	if err != nil {
		return nil, err
	}

	s := &IClientStats{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IClientStats creates a new IClientStats interface.
//
// Supported AppIDs: 1046930.
func (c *Client) IClientStats(appID uint32) (*IClientStats, error) {
	return NewIClientStats(c, appID)
}

// ReportEvent creates a Request for interface method ReportEvent.
func (i *IClientStats) ReportEvent() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "ReportEvent",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
