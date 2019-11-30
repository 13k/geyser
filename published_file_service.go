// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IPublishedFileService

package geyser

import "net/http"

// SchemaPublishedFileService stores the SchemaInterfaces for interface IPublishedFileService
var SchemaPublishedFileService = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "QueryFiles",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "enumeration EPublishedFileQueryType in clientenums.h",
						Name:        "query_type",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Current page",
						Name:        "page",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Cursor to paginate through the results (set to '*' for the first request).  Prefer this over using the page parameter, as it will allow you to do deep pagination.  When used, the page parameter will be ignored.",
						Name:        "cursor",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "(Optional) The number of results, per page to return.",
						Name:        "numperpage",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "App that created the files",
						Name:        "creator_appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "App that consumes the files",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Tags to match on. See match_all_tags parameter below",
						Name:        "requiredtags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "(Optional) Tags that must NOT be present on a published file to satisfy the query.",
						Name:        "excludedtags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "If true, then items must have all the tags specified, otherwise they must have at least one of the tags.",
						Name:        "match_all_tags",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Required flags that must be set on any returned items",
						Name:        "required_flags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Flags that must not be set on any returned items",
						Name:        "omitted_flags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Text to match in the item's title or description",
						Name:        "search_text",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "EPublishedFileInfoMatchingFileType",
						Name:        "filetype",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Find all items that reference the given item.",
						Name:        "child_publishedfileid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "If query_type is k_PublishedFileQueryType_RankedByTrend, then this is the number of days to get votes for [1,7].",
						Name:        "days",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "If query_type is k_PublishedFileQueryType_RankedByTrend, then limit result set just to items that have votes within the day range given",
						Name:        "include_recent_votes_only",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Allow stale data to be returned for the specified number of seconds.",
						Name:        "cache_max_age_seconds",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Language to search in and also what gets returned. Defaults to English.",
						Name:        "language",
						Optional:    true,
						Type:        "int32",
					},
					&SchemaMethodParam{
						Description: "Required key-value tags to match on.",
						Name:        "required_kv_tags",
						Optional:    false,
						Type:        "{message}",
					},
					&SchemaMethodParam{
						Description: "(Optional) If true, only return the total number of files that satisfy this query.",
						Name:        "totalonly",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "(Optional) If true, only return the published file ids of files that satisfy this query.",
						Name:        "ids_only",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return vote data",
						Name:        "return_vote_data",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return tags in the file details",
						Name:        "return_tags",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return key-value tags in the file details",
						Name:        "return_kv_tags",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return preview image and video details in the file details",
						Name:        "return_previews",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return child item ids in the file details",
						Name:        "return_children",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Populate the short_description field instead of file_description",
						Name:        "return_short_description",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return pricing information, if applicable",
						Name:        "return_for_sale_data",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Populate the metadata",
						Name:        "return_metadata",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return playtime stats for the specified number of days before today.",
						Name:        "return_playtime_stats",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "By default, if none of the other 'return_*' fields are set, only some voting details are returned. Set this to true to return the default set of details.",
						Name:        "return_details",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Strips BBCode from descriptions.",
						Name:        "strip_description_bbcode",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return the data for the specified revision.",
						Name:        "desired_revision",
						Optional:    true,
						Type:        "{enum}",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetDetails",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Set of published file Ids to retrieve details for.",
						Name:        "publishedfileids",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "If true, return tag information in the returned details.",
						Name:        "includetags",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return preview information in the returned details.",
						Name:        "includeadditionalpreviews",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return children in the returned details.",
						Name:        "includechildren",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return key value tags in the returned details.",
						Name:        "includekvtags",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return vote data in the returned details.",
						Name:        "includevotes",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return a short description instead of the full description.",
						Name:        "short_description",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, return pricing data, if applicable.",
						Name:        "includeforsaledata",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "If true, populate the metadata field.",
						Name:        "includemetadata",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Specifies the localized text to return. Defaults to English.",
						Name:        "language",
						Optional:    true,
						Type:        "int32",
					},
					&SchemaMethodParam{
						Description: "Return playtime stats for the specified number of days before today.",
						Name:        "return_playtime_stats",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Strips BBCode from descriptions.",
						Name:        "strip_description_bbcode",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return the data for the specified revision.",
						Name:        "desired_revision",
						Optional:    true,
						Type:        "{enum}",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetUserFiles",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Steam ID of the user whose files are being requested.",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "App Id of the app that the files were published to.",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "(Optional) Starting page for results.",
						Name:        "page",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "(Optional) The number of results, per page to return.",
						Name:        "numperpage",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "(Optional) Type of files to be returned.",
						Name:        "type",
						Optional:    true,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "(Optional) Sorting method to use on returned values.",
						Name:        "sortmethod",
						Optional:    true,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "(optional) Filter by privacy settings.",
						Name:        "privacy",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "(Optional) Tags that must be present on a published file to satisfy the query.",
						Name:        "requiredtags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "(Optional) Tags that must NOT be present on a published file to satisfy the query.",
						Name:        "excludedtags",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Required key-value tags to match on.",
						Name:        "required_kv_tags",
						Optional:    false,
						Type:        "{message}",
					},
					&SchemaMethodParam{
						Description: "(Optional) File type to match files to.",
						Name:        "filetype",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "App Id of the app that published the files, only matched if specified.",
						Name:        "creator_appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Match this cloud filename if specified.",
						Name:        "match_cloud_filename",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "Allow stale data to be returned for the specified number of seconds.",
						Name:        "cache_max_age_seconds",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Specifies the localized text to return. Defaults to English.",
						Name:        "language",
						Optional:    true,
						Type:        "int32",
					},
					&SchemaMethodParam{
						Description: "(Optional) If true, only return the total number of files that satisfy this query.",
						Name:        "totalonly",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "(Optional) If true, only return the published file ids of files that satisfy this query.",
						Name:        "ids_only",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return vote data",
						Name:        "return_vote_data",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return tags in the file details",
						Name:        "return_tags",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return key-value tags in the file details",
						Name:        "return_kv_tags",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return preview image and video details in the file details",
						Name:        "return_previews",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return child item ids in the file details",
						Name:        "return_children",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Populate the short_description field instead of file_description",
						Name:        "return_short_description",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return pricing information, if applicable",
						Name:        "return_for_sale_data",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Populate the metadata field",
						Name:        "return_metadata",
						Optional:    true,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return playtime stats for the specified number of days before today.",
						Name:        "return_playtime_stats",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Strips BBCode from descriptions.",
						Name:        "strip_description_bbcode",
						Optional:    false,
						Type:        "bool",
					},
					&SchemaMethodParam{
						Description: "Return the data for the specified revision.",
						Name:        "desired_revision",
						Optional:    true,
						Type:        "{enum}",
					},
				),
				Version: 1,
			},
		),
		Name: "IPublishedFileService",
	},
)

// PublishedFileService represents interface IPublishedFileService
type PublishedFileService struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewPublishedFileService creates a new PublishedFileService interface
func NewPublishedFileService(c *Client) (*PublishedFileService, error) {
	si, err := SchemaPublishedFileService.Get("IPublishedFileService", 0)

	if err != nil {
		return nil, err
	}

	s := &PublishedFileService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// PublishedFileService creates a new PublishedFileService interface
func (c *Client) PublishedFileService() (*PublishedFileService, error) {
	return NewPublishedFileService(c)
}

// GetDetails creates a Request for interface method GetDetails
func (i *PublishedFileService) GetDetails() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetDetails", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &PublishedFileServiceGetDetails{},
	}

	return req, nil
}

// GetUserFiles creates a Request for interface method GetUserFiles
func (i *PublishedFileService) GetUserFiles() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetUserFiles", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &PublishedFileServiceGetUserFiles{},
	}

	return req, nil
}

// QueryFiles creates a Request for interface method QueryFiles
func (i *PublishedFileService) QueryFiles() (*Request, error) {
	sm, err := i.Interface.Methods.Get("QueryFiles", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &PublishedFileServiceQueryFiles{},
	}

	return req, nil
}
