package openssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// RsaEncrypt1 加密字节数组，返回字节数组
func RsaEncrypt1(publicKey, origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt1 解密字节数组，返回字节数组
func RsaDecrypt1(privateKey, ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

// RsaEncryptString 加密字符串，返回base64处理的字符串
func RsaEncryptString(publicKey []byte, origData string) (string, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	cipherArr, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(cipherArr), nil
	}
}

// RsaDecryptString 解密经过base64处理的加密字符串，返回加密前的明⽂
func RsaDecryptString(privateKey []byte, cipherText string) (string, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
	originalArr, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherArr)
	if err != nil {
		return "", err
	} else {
		return string(originalArr), nil
	}
}
