package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/chacha20poly1305"
)

// EncryptMessage шифрует сообщение перед отправкой в бд
func EncryptMessage(text string, key []byte) (string, error) {
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err = rand.Read(nonce); err != nil {
		return "", err
	}

	ciphertext := aead.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptMessage расшифровыввает сообщение при получении из бд
func DecryptMessage(cryptoText string, key []byte) (string, error) {
	ciphertext, _ := base64.StdEncoding.DecodeString(cryptoText)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", err
	}

	nonceSize := chacha20poly1305.NonceSizeX
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plainText, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}
