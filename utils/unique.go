package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func GetUniq(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	ss := hex.EncodeToString(h.Sum(nil))
	ss = ss[len(ss)-32:]
	return ss
}
