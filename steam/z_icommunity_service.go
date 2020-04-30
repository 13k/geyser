// Code generated by geyser. DO NOT EDIT.
// API interface: ICommunityService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaICommunityService stores the Interfaces for interface ICommunityService.
var SchemaICommunityService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetApps",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetUserPartnerEventNews",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBestEventsForUser",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ICommunityService",
		Undocumented: true,
	},
)

// ICommunityService represents interface ICommunityService.
//
// This is an undocumented interface.
type ICommunityService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewICommunityService creates a new ICommunityService interface.
func NewICommunityService(c *Client) (*ICommunityService, error) {
	si, err := SchemaICommunityService.Get(schema.InterfaceKey{Name: "ICommunityService"})

	if err != nil {
		return nil, err
	}

	s := &ICommunityService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ICommunityService creates a new ICommunityService interface.
func (c *Client) ICommunityService() (*ICommunityService, error) {
	return NewICommunityService(c)
}

/*
GetApps creates a Request for interface method GetApps.

This is an undocumented method.
*/
func (i *ICommunityService) GetApps() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetApps",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetBestEventsForUser creates a Request for interface method GetBestEventsForUser.

This is an undocumented method.
*/
func (i *ICommunityService) GetBestEventsForUser() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBestEventsForUser",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetUserPartnerEventNews creates a Request for interface method GetUserPartnerEventNews.

This is an undocumented method.
*/
func (i *ICommunityService) GetUserPartnerEventNews() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetUserPartnerEventNews",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
