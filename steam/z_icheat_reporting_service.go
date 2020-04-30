// Code generated by geyser. DO NOT EDIT.
// API interface: ICheatReportingService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaICheatReportingService stores the Interfaces for interface ICheatReportingService.
var SchemaICheatReportingService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "ReportCheatData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "steamid of the user running and reporting the cheat.",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "The appid.",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "path and file name of the cheat executable.",
						Name:        "pathandfilename",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "web url where the cheat was found and downloaded.",
						Name:        "webcheaturl",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "local system time now.",
						Name:        "time_now",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "local system time when cheat process started. ( 0 if not yet run )",
						Name:        "time_started",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "local system time when cheat process stopped. ( 0 if still running )",
						Name:        "time_stopped",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "descriptive name for the cheat.",
						Name:        "cheatname",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "process ID of the running game.",
						Name:        "game_process_id",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "process ID of the cheat process that ran",
						Name:        "cheat_process_id",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "cheat param 1",
						Name:        "cheat_param_1",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "cheat param 2",
						Name:        "cheat_param_2",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "data collection in json format",
						Name:        "cheat_data_dump",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ICheatReportingService",
		Undocumented: false,
	},
)

// ICheatReportingService represents interface ICheatReportingService.
type ICheatReportingService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewICheatReportingService creates a new ICheatReportingService interface.
func NewICheatReportingService(c *Client) (*ICheatReportingService, error) {
	si, err := SchemaICheatReportingService.Get(schema.InterfaceKey{Name: "ICheatReportingService"})

	if err != nil {
		return nil, err
	}

	s := &ICheatReportingService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ICheatReportingService creates a new ICheatReportingService interface.
func (c *Client) ICheatReportingService() (*ICheatReportingService, error) {
	return NewICheatReportingService(c)
}

/*
ReportCheatData creates a Request for interface method ReportCheatData.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): steamid of the user running and reporting the cheat.
  * appid [uint32] (required): The appid.
  * pathandfilename [string] (required): path and file name of the cheat executable.
  * webcheaturl [string] (required): web url where the cheat was found and downloaded.
  * time_now [uint64] (required): local system time now.
  * time_started [uint64] (required): local system time when cheat process started. ( 0 if not yet run )
  * time_stopped [uint64] (required): local system time when cheat process stopped. ( 0 if still running )
  * cheatname [string] (required): descriptive name for the cheat.
  * game_process_id [uint32] (required): process ID of the running game.
  * cheat_process_id [uint32] (required): process ID of the cheat process that ran
  * cheat_param_1 [uint64] (required): cheat param 1
  * cheat_param_2 [uint64] (required): cheat param 2
  * cheat_data_dump [string] (required): data collection in json format
*/
func (i *ICheatReportingService) ReportCheatData() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "ReportCheatData",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
