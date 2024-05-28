package openssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// EncryptOAEP 加密
func EncryptOAEP(publicKey []byte, text string) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	pub := pubInterface.(*rsa.PublicKey)

	secretMessage := []byte(text)
	rng := rand.Reader
	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, pub, secretMessage, nil)
	if err != nil {
		return "", err
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	return ciphertext, nil
}

// DecryptOAEP 解密
func DecryptOAEP(privateKey []byte, ciphertext string) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	cipherData, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader
	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, priv, cipherData, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
