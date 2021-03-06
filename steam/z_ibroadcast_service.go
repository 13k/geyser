// Code generated by geyser. DO NOT EDIT.
// API interface: IBroadcastService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIBroadcastService stores the Interfaces for interface IBroadcastService.
var SchemaIBroadcastService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "PostGameDataFrameRTMP",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "AppID of the game being broadcasted",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Broadcasters SteamID",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Valid RTMP token for the Broadcaster",
						Name:        "rtmp_token",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "game data frame expressing current state of game (string, zipped, whatever)",
						Name:        "frame_data",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "PostChatMessage",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "MuteBroadcastChatUser",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "RemoveUserChatText",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetRTMPInfo",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "SetRTMPInfo",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "UpdateChatMessageFlair",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBroadcastUploadStats",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBroadcastViewerStats",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "StartBuildClip",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetBuildClipStatus",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetClipDetails",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IBroadcastService",
		Undocumented: false,
	},
)

// IBroadcastService represents interface IBroadcastService.
type IBroadcastService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIBroadcastService creates a new IBroadcastService interface.
func NewIBroadcastService(c *Client) (*IBroadcastService, error) {
	si, err := SchemaIBroadcastService.Get(schema.InterfaceKey{Name: "IBroadcastService"})

	if err != nil {
		return nil, err
	}

	s := &IBroadcastService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IBroadcastService creates a new IBroadcastService interface.
func (c *Client) IBroadcastService() (*IBroadcastService, error) {
	return NewIBroadcastService(c)
}

/*
GetBroadcastUploadStats creates a Request for interface method GetBroadcastUploadStats.

This is an undocumented method.
*/
func (i *IBroadcastService) GetBroadcastUploadStats() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBroadcastUploadStats",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetBroadcastViewerStats creates a Request for interface method GetBroadcastViewerStats.

This is an undocumented method.
*/
func (i *IBroadcastService) GetBroadcastViewerStats() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBroadcastViewerStats",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetBuildClipStatus creates a Request for interface method GetBuildClipStatus.

This is an undocumented method.
*/
func (i *IBroadcastService) GetBuildClipStatus() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBuildClipStatus",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetClipDetails creates a Request for interface method GetClipDetails.

This is an undocumented method.
*/
func (i *IBroadcastService) GetClipDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetClipDetails",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetRTMPInfo creates a Request for interface method GetRTMPInfo.

This is an undocumented method.
*/
func (i *IBroadcastService) GetRTMPInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetRTMPInfo",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
MuteBroadcastChatUser creates a Request for interface method MuteBroadcastChatUser.

This is an undocumented method.
*/
func (i *IBroadcastService) MuteBroadcastChatUser() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "MuteBroadcastChatUser",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
PostChatMessage creates a Request for interface method PostChatMessage.

This is an undocumented method.
*/
func (i *IBroadcastService) PostChatMessage() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "PostChatMessage",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
PostGameDataFrameRTMP creates a Request for interface method PostGameDataFrameRTMP.

Parameters

  * appid [uint32] (required): AppID of the game being broadcasted
  * steamid [uint64] (required): Broadcasters SteamID
  * rtmp_token [string] (required): Valid RTMP token for the Broadcaster
  * frame_data [string] (required): game data frame expressing current state of game (string, zipped, whatever)
*/
func (i *IBroadcastService) PostGameDataFrameRTMP() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "PostGameDataFrameRTMP",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
RemoveUserChatText creates a Request for interface method RemoveUserChatText.

This is an undocumented method.
*/
func (i *IBroadcastService) RemoveUserChatText() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "RemoveUserChatText",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
SetRTMPInfo creates a Request for interface method SetRTMPInfo.

This is an undocumented method.
*/
func (i *IBroadcastService) SetRTMPInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "SetRTMPInfo",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
StartBuildClip creates a Request for interface method StartBuildClip.

This is an undocumented method.
*/
func (i *IBroadcastService) StartBuildClip() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "StartBuildClip",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
UpdateChatMessageFlair creates a Request for interface method UpdateChatMessageFlair.

This is an undocumented method.
*/
func (i *IBroadcastService) UpdateChatMessageFlair() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "UpdateChatMessageFlair",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
