// Code generated by geyser. DO NOT EDIT.
// API interface: IEconItems.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIEconItems stores the Interfaces for interface IEconItems.
var SchemaIEconItems = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchema",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to return the names in. Defaults to returning string keys.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchemaItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to return the names in. Defaults to returning string keys.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The first item id to return. Defaults to 0. Response will indicate next value to query if applicable.",
						Name:        "start",
						Optional:    true,
						Type:        "int32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchemaOverview",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to return the names in. Defaults to returning string keys.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSchemaURL",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetStoreMetaData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to results in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetStoreStatus",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_440",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetEquippedPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Return items equipped for this class id",
						Name:        "class_id",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSchemaURL",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetStoreMetaData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to results in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_570",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchema",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to return the names in. Defaults to returning string keys.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_620",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchema",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to return the names in. Defaults to returning string keys.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      2,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSchemaURL",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      2,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetStoreMetaData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to results in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_730",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetEquippedPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Return items equipped for this class id",
						Name:        "class_id",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSchemaURL",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetStoreMetaData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to results in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_205790",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_221540",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_238460",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetEquippedPlayerItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Steam ID to fetch items for",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Return items equipped for this class id",
						Name:        "class_id",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconItems_583950",
		Undocumented: false,
	},
)

// IEconItems represents interface IEconItems.
//
// Supported AppIDs: 440, 570, 620, 730, 205790, 221540, 238460, 583950.
type IEconItems struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIEconItems creates a new IEconItems interface.
//
// Supported AppIDs: 440, 570, 620, 730, 205790, 221540, 238460, 583950.
func NewIEconItems(c *Client, appID uint32) (*IEconItems, error) {
	si, err := SchemaIEconItems.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IEconItems",
	})

	if err != nil {
		return nil, err
	}

	s := &IEconItems{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IEconItems creates a new IEconItems interface.
//
// Supported AppIDs: 440, 570, 620, 730, 205790, 221540, 238460, 583950.
func (c *Client) IEconItems(appID uint32) (*IEconItems, error) {
	return NewIEconItems(c, appID)
}

/*
GetEquippedPlayerItems creates a Request for interface method GetEquippedPlayerItems.

Parameters

  * steamid [uint64] (required): The Steam ID to fetch items for
  * class_id [uint32] (required): Return items equipped for this class id
*/
func (i *IEconItems) GetEquippedPlayerItems() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetEquippedPlayerItems",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetPlayerItems creates a Request for interface method GetPlayerItems.

Parameters

  * steamid [uint64] (required): The Steam ID to fetch items for
*/
func (i *IEconItems) GetPlayerItems() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetPlayerItems",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetSchema creates a Request for interface method GetSchema.

Supported versions: 1, 2.

Parameters (v1)

  * language [string]: The language to return the names in. Defaults to returning string keys.

Parameters (v2)

  * language [string]: The language to return the names in. Defaults to returning string keys.
*/
func (i *IEconItems) GetSchema(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSchema",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetSchemaItems creates a Request for interface method GetSchemaItems.

Parameters

  * language [string]: The language to return the names in. Defaults to returning string keys.
  * start [int32]: The first item id to return. Defaults to 0. Response will indicate next value to query if applicable.
*/
func (i *IEconItems) GetSchemaItems() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSchemaItems",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetSchemaOverview creates a Request for interface method GetSchemaOverview.

Parameters

  * language [string]: The language to return the names in. Defaults to returning string keys.
*/
func (i *IEconItems) GetSchemaOverview() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSchemaOverview",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetSchemaURL creates a Request for interface method GetSchemaURL.

Supported versions: 1, 2.
*/
func (i *IEconItems) GetSchemaURL(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSchemaURL",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

/*
GetStoreMetaData creates a Request for interface method GetStoreMetaData.

Parameters

  * language [string]: The language to results in.
*/
func (i *IEconItems) GetStoreMetaData() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetStoreMetaData",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}

// GetStoreStatus creates a Request for interface method GetStoreStatus.
func (i *IEconItems) GetStoreStatus() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetStoreStatus",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}