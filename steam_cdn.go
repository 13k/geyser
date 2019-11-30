// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ISteamCDN

package geyser

import "net/http"

// SchemaSteamCDN stores the SchemaInterfaces for interface ISteamCDN
var SchemaSteamCDN = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "SetClientFilters",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Steam name of CDN property",
						Name:        "cdnname",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "comma-separated list of allowed IP address blocks in CIDR format - blank for not used",
						Name:        "allowedipblocks",
						Optional:    true,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "comma-separated list of allowed client network AS numbers - blank for not used",
						Name:        "allowedasns",
						Optional:    true,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "comma-separated list of allowed client IP country codes in ISO 3166-1 format - blank for not used",
						Name:        "allowedipcountries",
						Optional:    true,
						Type:        "string",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "SetPerformanceStats",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Steam name of CDN property",
						Name:        "cdnname",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Outgoing network traffic in Mbps",
						Name:        "mbps_sent",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Incoming network traffic in Mbps",
						Name:        "mbps_recv",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Percent CPU load",
						Name:        "cpu_percent",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Percent cache hits",
						Name:        "cache_hit_percent",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ISteamCDN",
	},
)

// SteamCDN represents interface ISteamCDN
type SteamCDN struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewSteamCDN creates a new SteamCDN interface
func NewSteamCDN(c *Client) (*SteamCDN, error) {
	si, err := SchemaSteamCDN.Get("ISteamCDN", 0)

	if err != nil {
		return nil, err
	}

	s := &SteamCDN{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// SteamCDN creates a new SteamCDN interface
func (c *Client) SteamCDN() (*SteamCDN, error) {
	return NewSteamCDN(c)
}

// SetClientFilters creates a Request for interface method SetClientFilters
func (i *SteamCDN) SetClientFilters() (*Request, error) {
	sm, err := i.Interface.Methods.Get("SetClientFilters", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamCDNSetClientFilters{},
	}

	return req, nil
}

// SetPerformanceStats creates a Request for interface method SetPerformanceStats
func (i *SteamCDN) SetPerformanceStats() (*Request, error) {
	sm, err := i.Interface.Methods.Get("SetPerformanceStats", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamCDNSetPerformanceStats{},
	}

	return req, nil
}
