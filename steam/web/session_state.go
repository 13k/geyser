package web

//go:generate stringer -type=SessionState

// SessionState identifies a Session's current state.
type SessionState uint8

// Session states.
const (
	SessionInvalid SessionState = iota
	SessionAuthenticated
	SessionCaptchaRequiredInvalidLogin
	SessionCaptchaRequired
	SessionEmailCodeRequired
	SessionTwoFactorCodeRequired
	SessionInvalidLogin
)
