package cryptogram

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Sha1Encrypt sha1加密
func Sha1Encrypt(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// Sha256Encrypt sha256加密
func Sha256Encrypt(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	mdStr := hex.EncodeToString(bs)
	return mdStr
}
