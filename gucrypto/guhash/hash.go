package guhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func MD5string(mystr string) string {

	h := md5.New()
	h.Write([]byte(mystr))
	md5_hash := hex.EncodeToString(h.Sum(nil))

	return md5_hash
}

func SHA1string(mystr string) string {

	h := sha1.New()
	h.Write([]byte(mystr))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func SHA256string(mystr string) string {

	h := sha256.New()
	h.Write([]byte(mystr))
	sha256_hash := hex.EncodeToString(h.Sum(nil))

	return sha256_hash
}

// // // //

func MD5bytes(myBytes []byte) string {

	h := md5.New()
	h.Write(myBytes)
	md5_hash := hex.EncodeToString(h.Sum(nil))

	return md5_hash
}

func SHA1bytes(myBytes []byte) string {

	h := sha1.New()
	h.Write(myBytes)
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func SHA256bytes(myBytes []byte) string {

	h := sha256.New()
	h.Write(myBytes)
	sha256_hash := hex.EncodeToString(h.Sum(nil))

	return sha256_hash
}
