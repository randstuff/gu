package guhash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"os"
)

func MD5File(tFilePath string) (string, error) {

	var buf [4096]byte
	file, err := os.Open(tFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := md5.New()

	for read, err := file.Read(buf[:]); err != io.EOF && read != 0; read, err = file.Read(buf[:]) {
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return "", err
		}
		hash.Write(buf[:read])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func SHA1File(tFilePath string) (string, error) {

	var buf [4096]byte
	file, err := os.Open(tFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha1.New()

	for read, err := file.Read(buf[:]); err != io.EOF && read != 0; read, err = file.Read(buf[:]) {
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return "", err
		}
		hash.Write(buf[:read])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func SHA256File(tFilePath string) (string, error) {

	var buf [4096]byte
	file, err := os.Open(tFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()

	for read, err := file.Read(buf[:]); err != io.EOF && read != 0; read, err = file.Read(buf[:]) {
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return "", err
		}
		hash.Write(buf[:read])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func SHA512File(tFilePath string) (string, error) {

	var buf [4096]byte
	file, err := os.Open(tFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha512.New()

	for read, err := file.Read(buf[:]); err != io.EOF && read != 0; read, err = file.Read(buf[:]) {
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			return "", err
		}
		hash.Write(buf[:read])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
