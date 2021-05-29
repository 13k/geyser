package web

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/13k/go-steam/steamid"
)

// RSAKeyResult is the result of an RSA Key request.
type RSAKeyResult struct {
	PublicKeyMod string `json:"publickey_mod"`
	PublicKeyExp string `json:"publickey_exp"`
	Success      bool   `json:"success"`
	Timestamp    string `json:"timestamp"`
	TokenGID     string `json:"token_gid"`

	pub *rsa.PublicKey
	ts  *time.Time
}

// PubKey returns the RSA public key formed by the PublicKey* fields.
func (r *RSAKeyResult) PubKey() (*rsa.PublicKey, error) {
	if r.pub == nil {
		var err error

		r.pub, err = rsaPubKeyFromHexStrings(r.PublicKeyMod, r.PublicKeyExp)

		if err != nil {
			return nil, fmt.Errorf("geyser/steam/web: error parsing public key: %w", err)
		}
	}

	return r.pub, nil
}

// Time returns the parsed Timestamp.
func (r *RSAKeyResult) Time() (time.Time, error) {
	if r.ts == nil {
		n, err := strconv.ParseInt(r.Timestamp, 10, 64)

		if err != nil {
			return time.Time{}, fmt.Errorf("geyser/steam/web: error parsing unix timestamp: %w", err)
		}

		ts := time.Unix(n, 0)
		r.ts = &ts
	}

	return *r.ts, nil
}

// LoginResult is the result of a login request.
type LoginResult struct {
	Success            bool                           `json:"success"`
	Complete           bool                           `json:"login_complete"`
	CaptchaNeeded      bool                           `json:"captcha_needed"`
	CaptchaGID         int64                          `json:"captcha_gid"`
	EmailCodeNeeded    bool                           `json:"emailauth_needed"`
	EmailSteamID       steamid.SteamID                `json:"emailsteamid"`
	TwoFactorNeeded    bool                           `json:"requires_twofactor"`
	ClearPassword      bool                           `json:"clear_password_field"`
	Message            string                         `json:"message"`
	OAuthJSON          string                         `json:"oauth"`
	RedirectURI        string                         `json:"redirect_uri"`
	TransferURLs       []string                       `json:"transfer_urls"`
	TransferParameters *LoginResultTransferParameters `json:"transfer_parameters"`

	oauth *LoginResultOAuth
}

// OAuth parses the OAuthJSON field.
func (r *LoginResult) OAuth() (*LoginResultOAuth, error) {
	if r.oauth != nil {
		return r.oauth, nil
	}

	if r.OAuthJSON == "" {
		return nil, nil
	}

	result := &LoginResultOAuth{}

	if err := json.Unmarshal([]byte(r.OAuthJSON), result); err != nil {
		return nil, err
	}

	r.oauth = result

	return r.oauth, nil
}

// LoginResultTransferParameters are transfer parameters of a successful login request.
type LoginResultTransferParameters struct {
	SteamIDString string `json:"steamid"`
	TokenSecure   string `json:"token_secure"`
	Auth          string `json:"auth"`
	RememberLogin bool   `json:"remember_login"`

	steamid steamid.SteamID
}

// SteamID parses the SteamIDString field.
func (p *LoginResultTransferParameters) SteamID() (steamid.SteamID, error) {
	if p.steamid == 0 {
		n, err := strconv.ParseUint(p.SteamIDString, 10, 64)

		if err != nil {
			return 0, err
		}

		p.steamid = steamid.SteamID(n)
	}

	return p.steamid, nil
}

// LoginResultOAuth are OAuth values of a successful login request.
type LoginResultOAuth struct {
	SteamIDString string `json:"steamid"`
	AccountName   string `json:"account_name"`
	Token         string `json:"oauth_token"`
	WGToken       string `json:"wgtoken"`
	WGTokenSecure string `json:"wgtoken_secure"`

	steamid steamid.SteamID
}

// SteamID parses the SteamIDString field.
func (p *LoginResultOAuth) SteamID() (steamid.SteamID, error) {
	if p.steamid == 0 {
		n, err := strconv.ParseUint(p.SteamIDString, 10, 64)

		if err != nil {
			return 0, err
		}

		p.steamid = steamid.SteamID(n)
	}

	return p.steamid, nil
}
