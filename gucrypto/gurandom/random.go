package gurandom

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"
const numBytes = "01234567890"
const filechunk = 8192

func RandInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func RandStringBytes(n int) string {

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func RandBytes(n int) []byte {

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func RandMD5() string {

	s := RandStringBytes(128)

	h := md5.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func RandSHA1() string {

	s := RandStringBytes(128)

	h := sha1.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func RandSHA256() string {

	s := RandStringBytes(128)

	h := sha256.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}

func RandSHA512() string {

	s := RandStringBytes(128)

	h := sha512.New()
	h.Write([]byte(s))
	hash := hex.EncodeToString(h.Sum(nil))

	return hash
}
