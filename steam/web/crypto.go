package web

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1" //nolint: gosec
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
)

func randomBytes(size int) ([]byte, error) {
	buf := make([]byte, size)

	if _, err := rand.Read(buf); err != nil {
		return nil, fmt.Errorf("error reading crypto random source: %w", err)
	}

	return buf, nil
}

func hexString(data []byte) string {
	return hex.EncodeToString(data)
}

func base64String(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func sha1Hash(data []byte) []byte {
	arr := sha1.Sum(data)
	return arr[:]
}

func rsaPubKeyFromHexStrings(smod, sexp string) (*rsa.PublicKey, error) {
	mod, ok := big.NewInt(0).SetString(smod, 16)

	if !ok {
		return nil, fmt.Errorf("invalid RSA public key modulus %q", smod)
	}

	exp, err := strconv.ParseInt(sexp, 16, 0)

	if err != nil {
		return nil, fmt.Errorf("invalid RSA public key exponent %q: %w", sexp, err)
	}

	key := &rsa.PublicKey{
		N: mod,
		E: int(exp),
	}

	return key, nil
}

func encryptPKCS1v15(pub *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}
