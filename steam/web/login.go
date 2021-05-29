package web

import (
	"crypto/rsa"
	"strings"
	"time"

	"github.com/13k/go-steam/steamid"
)

// LoginOptions are options to pass to `Session.Login`.
type LoginOptions struct {
	Password      string
	Captcha       string
	EmailCode     string
	TwoFactorCode string
	Language      string
	Mobile        bool
}

// Copy returns a new copy of LoginOptions.
func (o *LoginOptions) Copy() *LoginOptions {
	ocopy := *o
	return &ocopy
}

// LoginCredentials are credentials given to `Client.Login`.
type LoginCredentials struct {
	Username      string
	Password      string
	Key           *rsa.PublicKey
	KeyTimestamp  time.Time
	CaptchaGID    int64
	Captcha       string
	EmailSteamID  steamid.SteamID
	EmailCode     string
	TwoFactorCode string
	LoginName     string
	OAuthClientID string
	OAuthScope    []string
}

func (c *LoginCredentials) form() (*loginForm, error) {
	xpasswd, err := encryptPKCS1v15(c.Key, []byte(c.Password))

	if err != nil {
		return nil, err
	}

	passwd := base64String(xpasswd)

	f := &loginForm{
		Username:      c.Username,
		Password:      passwd,
		EmailCode:     c.EmailCode,
		EmailSteamID:  c.EmailSteamID.Uint64(),
		TwoFactorCode: c.TwoFactorCode,
		CaptchaGID:    c.CaptchaGID,
		Captcha:       c.Captcha,
		LoginName:     c.LoginName,
		KeyTimestamp:  c.KeyTimestamp.Unix(),
		RememberLogin: true,
		OAuthClientID: c.OAuthClientID,
		OAuthScope:    strings.Join(c.OAuthScope, " "),
	}

	return f, nil
}
