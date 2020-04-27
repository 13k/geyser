// Code generated by geyser. DO NOT EDIT.
// API interface: IRemoteClientService.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIRemoteClientService stores the Interfaces for interface IRemoteClientService.
var SchemaIRemoteClientService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "NotifyRemotePacket",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "NotifyRegisterStatusUpdate",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "NotifyUnregisterStatusUpdate",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IRemoteClientService",
		Undocumented: true,
	},
)

// IRemoteClientService represents interface IRemoteClientService.
//
// This is an undocumented interface.
type IRemoteClientService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIRemoteClientService creates a new IRemoteClientService interface.
func NewIRemoteClientService(c *Client) (*IRemoteClientService, error) {
	si, err := SchemaIRemoteClientService.Get(schema.InterfaceKey{Name: "IRemoteClientService"})

	if err != nil {
		return nil, err
	}

	s := &IRemoteClientService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IRemoteClientService creates a new IRemoteClientService interface.
func (c *Client) IRemoteClientService() (*IRemoteClientService, error) {
	return NewIRemoteClientService(c)
}

/*
NotifyRegisterStatusUpdate creates a Request for interface method NotifyRegisterStatusUpdate.

This is an undocumented method.
*/
func (i *IRemoteClientService) NotifyRegisterStatusUpdate() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "NotifyRegisterStatusUpdate",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
NotifyRemotePacket creates a Request for interface method NotifyRemotePacket.

This is an undocumented method.
*/
func (i *IRemoteClientService) NotifyRemotePacket() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "NotifyRemotePacket",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
NotifyUnregisterStatusUpdate creates a Request for interface method NotifyUnregisterStatusUpdate.

This is an undocumented method.
*/
func (i *IRemoteClientService) NotifyUnregisterStatusUpdate() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "NotifyUnregisterStatusUpdate",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
