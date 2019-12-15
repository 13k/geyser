package geyser

// SteamWebAPIUtilGetServerInfo holds the result of the method ISteamWebAPIUtil/GetServerInfo.
type SteamWebAPIUtilGetServerInfo struct {
	ServerTime       uint64 `json:"servertime"`
	ServerTimeString string `json:"servertimestring"`
}

// SteamWebAPIUtilGetSupportedAPIList holds the result of the method ISteamWebAPIUtil/GetSupportedAPIList.
type SteamWebAPIUtilGetSupportedAPIList struct {
	API *Schema `json:"apilist"`
}
