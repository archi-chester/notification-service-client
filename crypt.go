package nsclient

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

var keyPhrase = "6368616e676520746869732070617373776f726420746f206120736563726574"
var noncePhrase = "64a9433eae7ccceee2fc0eda"

// шифруем сообщение
func encrypt(plaintext []byte) ([]byte, error) {
	// Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(keyPhrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		err = fmt.Errorf("Ошибка в шифровании(aes.NewCipher): %s", err)
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.

	// nonce, _ := hex.DecodeString("64a9433eae7ccceee2fc0eda")
	// nonce := make([]byte, 12)
	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	panic(err.Error())
	// }

	nonce, _ := hex.DecodeString(noncePhrase)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		err = fmt.Errorf("Ошибка в шифровании(cipher.NewGCM): %s", err)
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)

	// возвращаем сообщение
	return ciphertext, nil
}

// расшифровываем сообщение
func decrypt(ciphertext []byte) ([]byte, error) { // Load your secret key from a safe place and reuse it across multiple
	// Seal/Open calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	// When decoded the key should be 16 bytes (AES-128) or 32 (AES-256).
	key, _ := hex.DecodeString(keyPhrase)
	nonce, _ := hex.DecodeString(noncePhrase)

	block, err := aes.NewCipher(key)
	if err != nil {
		err = fmt.Errorf("Ошибка в шифровании(aes.NewCipher): %s", err)
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		err = fmt.Errorf("Ошибка в шифровании(cipher.NewGCM): %s", err)
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		err = fmt.Errorf("Ошибка в шифровании (aesgcm.Open): %s", err)
		panic(err.Error())
	}

	// возвращаем сообщение
	return plaintext, nil
}
