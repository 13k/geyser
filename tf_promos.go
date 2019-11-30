// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ITFPromos.

package geyser

import "net/http"

// SchemaTFPromos stores the SchemaInterfaces for interface ITFPromos.
var SchemaTFPromos = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemID",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "GrantItem",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ITFPromos_205790",
	},
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemID",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "GrantItem",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ITFPromos_440",
	},
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemID",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "GrantItem",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "promoid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ITFPromos_570",
	},
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemID",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "PromoID",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "GrantItem",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "The promo ID to grant an item for",
						Name:        "PromoID",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ITFPromos_620",
	},
)

// TFPromos represents interface ITFPromos.
//
// Supported AppIDs: [205790 440 570 620].
type TFPromos struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewTFPromos creates a new TFPromos interface.
//
// Supported AppIDs: [205790 440 570 620].
func NewTFPromos(c *Client, appID uint32) (*TFPromos, error) {
	si, err := SchemaTFPromos.Get("ITFPromos", appID)

	if err != nil {
		return nil, err
	}

	s := &TFPromos{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// TFPromos creates a new TFPromos interface.
//
// Supported AppIDs: [205790 440 570 620].
func (c *Client) TFPromos(appID uint32) (*TFPromos, error) {
	return NewTFPromos(c, appID)
}

// GetItemID creates a Request for interface method GetItemID.
func (i *TFPromos) GetItemID() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetItemID", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &TFPromosGetItemID{},
	}

	return req, nil
}

// GrantItem creates a Request for interface method GrantItem.
func (i *TFPromos) GrantItem() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GrantItem", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &TFPromosGrantItem{},
	}

	return req, nil
}
