package web

import (
	"net/url"

	"github.com/go-playground/form/v4"
)

var (
	formEncoder *form.Encoder
)

func init() {
	formEncoder = form.NewEncoder()
}

type requestForm interface {
	FormValues() (url.Values, error)
}

type rsaKeyForm struct {
	Username string `form:"username"` // plain-text
}

var _ requestForm = (*rsaKeyForm)(nil)

func (f *rsaKeyForm) FormValues() (url.Values, error) {
	return formEncoder.Encode(f)
}

type loginForm struct {
	Username      string `form:"username"`                  // plain-text
	Password      string `form:"password"`                  // base64-encoded PKCS1v15 encrypted
	KeyTimestamp  int64  `form:"rsatimestamp"`              // plain-text unix timestamp (seconds)
	CaptchaGID    int64  `form:"captchagid"`                // plain-text number
	Captcha       string `form:"captcha_text,omitempty"`    // plain-text captcha challenge answer
	EmailSteamID  uint64 `form:"emailsteamid,omitempty"`    // plain-text number
	EmailCode     string `form:"emailauth,omitempty"`       // plain-text
	TwoFactorCode string `form:"twofactorcode,omitempty"`   // plain-text
	OAuthClientID string `form:"oauth_client_id,omitempty"` // plain-text hardcoded value
	OAuthScope    string `form:"oauth_scope,omitempty"`     // plain-text space-separated list of oauth scope items
	LoginName     string `form:"loginfriendlyname"`         // plain-text hardcoded value
	RememberLogin bool   `form:"remember_login"`            // plain-text boolean ("true"/"false")
}

var _ requestForm = (*loginForm)(nil)

func (f *loginForm) FormValues() (url.Values, error) {
	return formEncoder.Encode(f)
}
