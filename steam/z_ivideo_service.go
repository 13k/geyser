// Code generated by geyser. DO NOT EDIT.
// API interface: IVideoService.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIVideoService stores the Interfaces for interface IVideoService.
var SchemaIVideoService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "SetVideoBookmark",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetVideoBookmarks",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IVideoService",
		Undocumented: true,
	},
)

// IVideoService represents interface IVideoService.
//
// This is an undocumented interface.
type IVideoService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIVideoService creates a new IVideoService interface.
func NewIVideoService(c *Client) (*IVideoService, error) {
	si, err := SchemaIVideoService.Get(schema.InterfaceKey{Name: "IVideoService"})

	if err != nil {
		return nil, err
	}

	s := &IVideoService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IVideoService creates a new IVideoService interface.
func (c *Client) IVideoService() (*IVideoService, error) {
	return NewIVideoService(c)
}

/*
GetVideoBookmarks creates a Request for interface method GetVideoBookmarks.

This is an undocumented method.
*/
func (i *IVideoService) GetVideoBookmarks() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetVideoBookmarks",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
SetVideoBookmark creates a Request for interface method SetVideoBookmark.

This is an undocumented method.
*/
func (i *IVideoService) SetVideoBookmark() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "SetVideoBookmark",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
