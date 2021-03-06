// Code generated by geyser. DO NOT EDIT.
// API interface: IQuestService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIQuestService stores the Interfaces for interface IQuestService.
var SchemaIQuestService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetCommunityItemDefinitions",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IQuestService",
		Undocumented: true,
	},
)

// IQuestService represents interface IQuestService.
//
// This is an undocumented interface.
type IQuestService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIQuestService creates a new IQuestService interface.
func NewIQuestService(c *Client) (*IQuestService, error) {
	si, err := SchemaIQuestService.Get(schema.InterfaceKey{Name: "IQuestService"})

	if err != nil {
		return nil, err
	}

	s := &IQuestService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IQuestService creates a new IQuestService interface.
func (c *Client) IQuestService() (*IQuestService, error) {
	return NewIQuestService(c)
}

/*
GetCommunityItemDefinitions creates a Request for interface method GetCommunityItemDefinitions.

This is an undocumented method.
*/
func (i *IQuestService) GetCommunityItemDefinitions() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCommunityItemDefinitions",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
