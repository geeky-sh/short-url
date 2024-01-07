package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetUniq(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	ss := hex.EncodeToString(h.Sum(nil))
	ss = ss[len(ss)-5:]
	return ss
}

func GetShortKey(size int) string {
	coll := "ABCDEFGHIJKLMNOPQRSTUVWQYZabcdefghijklmopqrstuvwxyz1234567890"

	key := make([]byte, size)
	for i := range key {
		ridx := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(coll))
		key[i] = coll[ridx]
	}
	return string(key)
}
