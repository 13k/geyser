package steam

import (
	"crypto/aes"
	"crypto/rand"
	"net/url"
	"strconv"

	"github.com/13k/go-steam"
	"github.com/13k/go-steam-resources/steamlang"
	"github.com/13k/go-steam/cryptoutil"
)

// ISteamUserAuthAuthenticateUserFormData generates request form data for method
// ISteamUserAuth/AuthenticateUser.
func ISteamUserAuthAuthenticateUserFormData(steamID uint64, loginKey string) (url.Values, error) {
	sessionKey := make([]byte, 32)

	if _, err := rand.Read(sessionKey); err != nil {
		return nil, err
	}

	encryptedSessionKey, err := cryptoutil.RSAEncrypt(steam.GetPublicKey(steamlang.EUniverse_Public), sessionKey)

	if err != nil {
		return nil, err
	}

	ciph, err := aes.NewCipher(sessionKey)

	if err != nil {
		return nil, err
	}

	encryptedLoginKey, err := cryptoutil.SymmetricEncrypt(ciph, []byte(loginKey))

	if err != nil {
		return nil, err
	}

	values := url.Values{}

	values.Add("steamid", strconv.FormatUint(steamID, 10))
	values.Add("sessionkey", string(encryptedSessionKey))
	values.Add("encrypted_loginkey", string(encryptedLoginKey))

	return values, nil
}
