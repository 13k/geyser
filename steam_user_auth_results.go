package geyser

// SteamUserAuthAuthenticateUser holds the result of the method ISteamUserAuth/AuthenticateUser.
type SteamUserAuthAuthenticateUser struct {
	AuthenticateUser SteamUserAuthAuthenticateUserAuthenticateUser `json:"authenticateuser"`
}

type SteamUserAuthAuthenticateUserAuthenticateUser struct {
	Token       string `json:"token"`
	TokenSecure string `json:"token_secure"`
}

// SteamUserAuthAuthenticateUserTicket holds the result of the method ISteamUserAuth/AuthenticateUserTicket.
type SteamUserAuthAuthenticateUserTicket struct{}
