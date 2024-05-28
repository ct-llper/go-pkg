package openssl

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Test_Aes(t *testing.T) {
	source := "hello world"
	fmt.Println("原字符：", source)
	//16byte密钥
	key := "1443flfsaWfdas"
	encryptCode := AesEncrypt([]byte(source), []byte(key))
	fmt.Println("密文：", string(encryptCode))

	decryptCode := AesDecrypt(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}

func Test_AesCBC(t *testing.T) {
	orig := "hello world"
	key := "0123456789012345"
	fmt.Println("原文：", orig)
	encryptCode := AesEncryptCBC(orig, key)
	fmt.Println("密文：", encryptCode)
	decryptCode := AesDecryptCBC(encryptCode, key)
	fmt.Println("解密结果：", decryptCode)
}

func Test_AesCtr(t *testing.T) {
	source := "hello world"
	fmt.Println("原字符：", source)

	key := "1443flfsaWfdasds"
	encryptCode, _ := AesCtrCrypt([]byte(source), []byte(key))
	fmt.Println("密文：", string(encryptCode))

	decryptCode, _ := AesCtrCrypt(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}

func Test_AesCFB(t *testing.T) {
	source := "hello world"
	fmt.Println("原字符：", source)
	key := "ABCDEFGHIJKLMNO1" //16位
	encryptCode := AesEncryptCFB([]byte(source), []byte(key))
	fmt.Println("密文：", hex.EncodeToString(encryptCode))
	decryptCode := AesDecryptCFB(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}

func Test_AesOFB(t *testing.T) {
	source := "hello world"
	fmt.Println("原字符：", source)
	key := "1111111111111111" //16位  32位均可
	encryptCode, _ := AesEncryptOFB([]byte(source), []byte(key))
	fmt.Println("密文：", hex.EncodeToString(encryptCode))
	decryptCode, _ := AesDecryptOFB(encryptCode, []byte(key))

	fmt.Println("解密：", string(decryptCode))
}
