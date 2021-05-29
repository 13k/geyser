package web

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/13k/go-steam/steamid"
)

const (
	sessionIDLength = 32

	cookieNameSessionID = "sessionid"
	cookieNameLanguage  = "Steam_Language"

	defaultLoginName = "geyser/web"
	oauthClientID    = "DE45CD61"
)

var (
	cookieBirthTime                 = &http.Cookie{Name: "birthtime", Value: "-3333"}
	cookieMobileClientPlatform      = &http.Cookie{Name: "mobileClient", Value: "android"}
	cookieUnsetMobileClientPlatform = &http.Cookie{Name: "mobileClient", MaxAge: -1}
	cookieMobileClientVersion       = &http.Cookie{Name: "mobileClientVersion", Value: "0 (2.1.3)"}
	cookieUnsetMobileClientVersion  = &http.Cookie{Name: "mobileClientVersion", MaxAge: -1}

	oauthScope = []string{
		"read_profile",
		"write_profile",
		"read_client",
		"write_client",
	}
)

// Session manages authentication and session information for Steam websites.
type Session struct {
	// State indicates current session state.
	State SessionState
	// SteamID is the user's SteamID. It's set after a successful login or when email code is required.
	SteamID steamid.SteamID
	// ID is the session ID, the SHA1 hash of 32 random bytes.
	ID string
	// Username of session's account.
	Username string
	// KeyResult is the result of a request to retrieve the user's RSA key.
	KeyResult *RSAKeyResult
	// LoginResult is the result of the previous `Login` attempt.
	LoginResult *LoginResult
	// OAuthResult is the OAuth result of a successful mobile login.
	OAuthResult *LoginResultOAuth
	// PubKey is the user's RSA public key.
	PubKey *rsa.PublicKey
	// PubKeyTimestamp is the user's RSA public key timestamp.
	PubKeyTimestamp time.Time

	client     *Client
	options    *LoginOptions
	captchaGID int64
}

// NewSession creates a new unauthenticated session for the given username.
func NewSession(username string) *Session {
	return &Session{
		State:      SessionInvalid,
		Username:   username,
		captchaGID: -1,
		client:     NewClient(),
	}
}

// CookieJar returns the cookie jar passed when the session was created or the default jar.
func (s *Session) CookieJar() http.CookieJar {
	return s.client.CookieJar()
}

// Cookies returns all stored cookies, keyed by domain.
func (s *Session) Cookies() map[*url.URL][]*http.Cookie {
	return s.client.cookies()
}

// AuthorizedDomains returns a list of domains for which this session is authorized.
//
// Returns nil if the session is not authenticated.
func (s *Session) AuthorizedDomains() []*url.URL {
	switch s.State {
	case SessionAuthenticated:
		return s.client.domains()
	default:
		return nil
	}
}

// CaptchaGID returns the ID of the captcha challenge.
//
// Returns -1 if the session is not awaiting a captcha answer.
func (s *Session) CaptchaGID() int64 {
	switch s.State {
	case SessionCaptchaRequired, SessionCaptchaRequiredInvalidLogin:
		return s.captchaGID
	default:
		return -1
	}
}

func (s *Session) genID() error {
	data, err := randomBytes(sessionIDLength)

	if err != nil {
		return err
	}

	s.ID = hexString(sha1Hash(data))

	return nil
}

func (s *Session) fetchKey() error {
	if s.PubKey != nil {
		return nil
	}

	result, err := s.client.RSAKey(s.Username)

	if err != nil {
		return err
	}

	s.KeyResult = result

	if !result.Success {
		return fmt.Errorf("geyser/steam/web: error fetching user key")
	}

	if s.PubKey, err = result.PubKey(); err != nil {
		return fmt.Errorf("geyser/steam/web: error fetching user key: %w", err)
	}

	if s.PubKeyTimestamp, err = result.Time(); err != nil {
		return fmt.Errorf("geyser/steam/web: error fetching user key: %w", err)
	}

	return nil
}

func (s *Session) credentials() *LoginCredentials {
	credentials := &LoginCredentials{
		Username:      s.Username,
		Password:      s.options.Password,
		Key:           s.PubKey,
		KeyTimestamp:  s.PubKeyTimestamp,
		CaptchaGID:    s.captchaGID,
		Captcha:       s.options.Captcha,
		EmailSteamID:  s.SteamID,
		EmailCode:     s.options.EmailCode,
		TwoFactorCode: s.options.TwoFactorCode,
		LoginName:     defaultLoginName,
	}

	if s.options.Mobile {
		credentials.OAuthClientID = oauthClientID
		credentials.OAuthScope = oauthScope
	}

	return credentials
}

// Login tries to authenticate the session with given password and options.
//
// It's a noop if already authenticated.
func (s *Session) Login(options *LoginOptions) (SessionState, error) {
	if s.State == SessionAuthenticated {
		return s.State, nil
	}

	if options.Password == "" {
		return s.State, errors.New("geyser/steam/web: password is required")
	}

	s.options = options.Copy()

	if s.options.Language == "" {
		s.options.Language = "english"
	}

	if err := s.fetchKey(); err != nil {
		return s.State, fmt.Errorf("geyser/steam/web: error fetching user key: %w", err)
	}

	if s.options.Mobile {
		s.client.setCookies([]*http.Cookie{
			cookieMobileClientPlatform,
			cookieMobileClientVersion,
		})
	}

	result, err := s.client.Login(s.credentials())

	if err != nil {
		return s.State, err
	}

	s.LoginResult = result

	switch {
	case result.Success && result.Complete:
		if err := s.loggedIn(result); err != nil {
			return s.State, err
		}
	case result.CaptchaNeeded:
		if result.ClearPassword {
			s.State = SessionCaptchaRequiredInvalidLogin
		} else {
			s.State = SessionCaptchaRequired
		}

		s.captchaGID = result.CaptchaGID
	case result.EmailCodeNeeded:
		s.SteamID = result.EmailSteamID
		s.State = SessionEmailCodeRequired
	case result.TwoFactorNeeded:
		s.State = SessionTwoFactorCodeRequired
	default:
		s.State = SessionInvalidLogin
	}

	return s.State, nil
}

func (s *Session) loggedIn(result *LoginResult) error {
	var err error

	if err = s.genID(); err != nil {
		return fmt.Errorf("geyser/steam/web: error generating session ID: %w", err)
	}

	if result.TransferParameters != nil {
		if s.SteamID, err = result.TransferParameters.SteamID(); err != nil {
			return fmt.Errorf("geyser/steam/web: error parsing SteamID: %w", err)
		}
	}

	if s.OAuthResult, err = result.OAuth(); err != nil {
		return fmt.Errorf("geyser/steam/web: error parsing OAuth result: %w", err)
	}

	if s.OAuthResult != nil {
		if s.SteamID, err = s.OAuthResult.SteamID(); err != nil {
			return fmt.Errorf("geyser/steam/web: error parsing SteamID: %w", err)
		}
	}

	mirrors := make([]*url.URL, len(result.TransferURLs))

	for i, transferURL := range result.TransferURLs {
		if mirrors[i], err = url.Parse(transferURL); err != nil {
			return fmt.Errorf("geyser/steam/web: error parsing transfer URL %q: %w", transferURL, err)
		}
	}

	s.State = SessionAuthenticated
	s.captchaGID = -1

	s.client.addMirrors(mirrors)

	s.client.transferCookies()

	s.client.setCookies([]*http.Cookie{
		cookieUnsetMobileClientPlatform,
		cookieUnsetMobileClientVersion,
		cookieBirthTime,
		{Name: cookieNameLanguage, Value: s.options.Language},
		{Name: cookieNameSessionID, Value: s.ID},
	})

	return nil
}
