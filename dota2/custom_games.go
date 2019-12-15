// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ICustomGames.

package dota2

import (
	"github.com/13k/geyser"
	"net/http"
)

// SchemaCustomGames stores the SchemaInterfaces for interface ICustomGames.
var SchemaCustomGames = geyser.MustNewSchemaInterfaces(
	&geyser.SchemaInterface{
		Methods: geyser.MustNewSchemaMethods(
			&geyser.SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetGamePlayerCounts",
				Params:       geyser.NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&geyser.SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetPopularGames",
				Params:       geyser.NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&geyser.SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSuggestedGames",
				Params:       geyser.NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&geyser.SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "GetWorkshopVoteQueue",
				Params:       geyser.NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ICustomGames",
		Undocumented: true,
	},
)

// CustomGames represents interface ICustomGames.
//
// This is an undocumented interface.
type CustomGames struct {
	Client    *Client
	Interface *geyser.SchemaInterface
}

// NewCustomGames creates a new CustomGames interface.
func NewCustomGames(c *Client) (*CustomGames, error) {
	si, err := SchemaCustomGames.Get(geyser.SchemaInterfaceKey{Name: "ICustomGames"})

	if err != nil {
		return nil, err
	}

	s := &CustomGames{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// CustomGames creates a new CustomGames interface.
func (c *Client) CustomGames() (*CustomGames, error) {
	return NewCustomGames(c)
}

/*
GetGamePlayerCounts creates a Request for interface method GetGamePlayerCounts.

This is an undocumented method.
*/
func (i *CustomGames) GetGamePlayerCounts() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(geyser.SchemaMethodKey{
		Name:    "GetGamePlayerCounts",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &CustomGamesGetGamePlayerCounts{},
	}

	return req, nil
}

/*
GetPopularGames creates a Request for interface method GetPopularGames.

This is an undocumented method.
*/
func (i *CustomGames) GetPopularGames() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(geyser.SchemaMethodKey{
		Name:    "GetPopularGames",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &CustomGamesGetPopularGames{},
	}

	return req, nil
}

/*
GetSuggestedGames creates a Request for interface method GetSuggestedGames.

This is an undocumented method.
*/
func (i *CustomGames) GetSuggestedGames() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(geyser.SchemaMethodKey{
		Name:    "GetSuggestedGames",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &CustomGamesGetSuggestedGames{},
	}

	return req, nil
}

/*
GetWorkshopVoteQueue creates a Request for interface method GetWorkshopVoteQueue.

This is an undocumented method.
*/
func (i *CustomGames) GetWorkshopVoteQueue() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(geyser.SchemaMethodKey{
		Name:    "GetWorkshopVoteQueue",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &CustomGamesGetWorkshopVoteQueue{},
	}

	return req, nil
}