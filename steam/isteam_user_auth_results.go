package steam

// ISteamUserAuthAuthenticateUser can be used as result of method ISteamUserAuth/AuthenticateUser.
type ISteamUserAuthAuthenticateUser struct {
	AuthenticateUser ISteamUserAuthAuthenticateUserAuthenticateUser `json:"authenticateuser"`
}

// ISteamUserAuthAuthenticateUserAuthenticateUser is part of result struct
// ISteamUserAuthAuthenticateUser.
type ISteamUserAuthAuthenticateUserAuthenticateUser struct {
	Token       string `json:"token"`
	TokenSecure string `json:"token_secure"`
}
