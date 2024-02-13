package encoding

import (
	"crypto/sha256"
	"encoding/hex"
)

func Encode(secret string) string {
	h := sha256.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum(nil))
}
