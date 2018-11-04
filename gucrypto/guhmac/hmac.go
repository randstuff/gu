package guhmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GernerateSHA2_HMAC(tSecret string, tData string) string {

	h := hmac.New(sha256.New, []byte(tSecret))
	h.Write([]byte(tData))

	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
