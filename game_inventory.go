// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IGameInventory.

package geyser

import "net/http"

// SchemaGameInventory stores the SchemaInterfaces for interface IGameInventory.
var SchemaGameInventory = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: MustNewSchemaMethods(
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetItemDefArchive",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IGameInventory",
		Undocumented: true,
	},
)

// GameInventory represents interface IGameInventory.
//
// This is an undocumented interface.
type GameInventory struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewGameInventory creates a new GameInventory interface.
func NewGameInventory(c *Client) (*GameInventory, error) {
	si, err := SchemaGameInventory.Get(SchemaInterfaceKey{Name: "IGameInventory"})

	if err != nil {
		return nil, err
	}

	s := &GameInventory{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// GameInventory creates a new GameInventory interface.
func (c *Client) GameInventory() (*GameInventory, error) {
	return NewGameInventory(c)
}

/*
GetItemDefArchive creates a Request for interface method GetItemDefArchive.

This is an undocumented method.
*/
func (i *GameInventory) GetItemDefArchive() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "GetItemDefArchive",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &GameInventoryGetItemDefArchive{},
	}

	return req, nil
}