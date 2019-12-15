// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IFriendMessagesService.

package geyser

import "net/http"

// SchemaFriendMessagesService stores the SchemaInterfaces for interface IFriendMessagesService.
var SchemaFriendMessagesService = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: MustNewSchemaMethods(
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetActiveMessageSessions",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetRecentMessages",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "MarkOfflineMessagesRead",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IFriendMessagesService",
		Undocumented: true,
	},
)

// FriendMessagesService represents interface IFriendMessagesService.
//
// This is an undocumented interface.
type FriendMessagesService struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewFriendMessagesService creates a new FriendMessagesService interface.
func NewFriendMessagesService(c *Client) (*FriendMessagesService, error) {
	si, err := SchemaFriendMessagesService.Get(SchemaInterfaceKey{Name: "IFriendMessagesService"})

	if err != nil {
		return nil, err
	}

	s := &FriendMessagesService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// FriendMessagesService creates a new FriendMessagesService interface.
func (c *Client) FriendMessagesService() (*FriendMessagesService, error) {
	return NewFriendMessagesService(c)
}

/*
GetActiveMessageSessions creates a Request for interface method GetActiveMessageSessions.

This is an undocumented method.
*/
func (i *FriendMessagesService) GetActiveMessageSessions() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "GetActiveMessageSessions",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &FriendMessagesServiceGetActiveMessageSessions{},
	}

	return req, nil
}

/*
GetRecentMessages creates a Request for interface method GetRecentMessages.

This is an undocumented method.
*/
func (i *FriendMessagesService) GetRecentMessages() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "GetRecentMessages",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &FriendMessagesServiceGetRecentMessages{},
	}

	return req, nil
}

/*
MarkOfflineMessagesRead creates a Request for interface method MarkOfflineMessagesRead.

This is an undocumented method.
*/
func (i *FriendMessagesService) MarkOfflineMessagesRead() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "MarkOfflineMessagesRead",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &FriendMessagesServiceMarkOfflineMessagesRead{},
	}

	return req, nil
}