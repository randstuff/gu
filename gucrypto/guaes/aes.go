package guaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"io/ioutil"

	"guencode/guzip"
	"gulogger"
)

// The key length can be 32, 24, 16  bytes (OR in bits: 128, 192 or 256)
func EncryptDir(tDirSourcePath string, tDestZipFilePath string, tDestEncryptedArchiveFilePath string, tKey string) {

	gulogger.InitLogger()

	guzip.Zip(tDirSourcePath, tDestZipFilePath)
	encrypt_file_data(tDestZipFilePath, tDestEncryptedArchiveFilePath, tKey)

}

// The key length can be 32, 24, 16  bytes (OR in bits: 128, 192 or 256)
func encrypt_file_data(src_file_path, des_file_path string, tKey string) {

	var ciphertext, plaintext []byte
	//	var err error

	// The key length can be 32, 24, 16  bytes (OR in bits: 128, 192 or 256)
	//	key := []byte("longer means more possible keys ")

	key := []byte(tKey)

	plaintext, err := ioutil.ReadFile(src_file_path)

	if err == nil {
		ciphertext, err = encrypt(key, plaintext)

		if err == nil {
			gulogger.Info.Println("Encrypting : ", src_file_path)
			err = ioutil.WriteFile(des_file_path, ciphertext, 0755)
			if err == nil {
				gulogger.Info.Println("DONE - Encrypted files : ", des_file_path)
			}
		} else {
			gulogger.Error.Println(err, " - ", src_file_path)
		}

	} else {
		gulogger.Error.Println(err)
	}

	return

}

func decrypt_file_data(src_file_path, des_file_path string) {

	var ciphertext, plaintext []byte
	var err error

	// The key length can be 32, 24, 16  bytes (OR in bits: 128, 192 or 256)
	key := []byte("longer means more possible keys ")

	ciphertext, _ = ioutil.ReadFile(src_file_path)
	if plaintext, err = decrypt(key, ciphertext); err != nil {
		//		log.Fatal(err)
		gulogger.Error.Println(err, " - ", src_file_path)
	}
	ioutil.WriteFile(des_file_path, plaintext, 0755)

	return
}

func encrypt(key, text []byte) (ciphertext []byte, err error) {

	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}

	ciphertext = make([]byte, aes.BlockSize+len(string(text)))

	// iv =  initialization vector
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return
}

func decrypt(key, ciphertext []byte) (plaintext []byte, err error) {

	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	plaintext = ciphertext

	return
}
