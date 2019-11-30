package geyser

import (
	"crypto/aes"
	"crypto/rand"
	"net/url"
	"strconv"

	"github.com/faceit/go-steam"
	"github.com/faceit/go-steam/cryptoutil"
	"github.com/faceit/go-steam/protocol/steamlang"
)

func SteamUserAuthAuthenticateUserFormData(steamID uint64, loginKey string) (url.Values, error) {
	sessionKey := make([]byte, 32)

	if _, err := rand.Read(sessionKey); err != nil {
		return nil, err
	}

	encryptedSessionKey := cryptoutil.RSAEncrypt(steam.GetPublicKey(steamlang.EUniverse_Public), sessionKey)

	ciph, err := aes.NewCipher(sessionKey)

	if err != nil {
		return nil, err
	}

	encryptedLoginKey := cryptoutil.SymmetricEncrypt(ciph, []byte(loginKey))

	values := url.Values{}

	values.Add("steamid", strconv.FormatUint(steamID, 10))
	values.Add("sessionkey", string(encryptedSessionKey))
	values.Add("encrypted_loginkey", string(encryptedLoginKey))

	return values, nil
}
