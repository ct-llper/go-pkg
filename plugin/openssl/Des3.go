package openssl

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// TripleDesEncrypt 3DES加密字节数组，返回字节数组
func TripleDesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	originalBytes = PKCS5Padding(originalBytes, block.BlockSize())
	// originalBytes = ZeroPadding(originalBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherArr := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(cipherArr, originalBytes)
	return cipherArr, nil
}

// TripleDesDecrypt 3DES解密字节数组，返回字节数组
func TripleDesDecrypt(cipherBytes, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	originalArr := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(originalArr, cipherBytes)
	originalArr = PKCS5UnPadding(originalArr)
	// origData = ZeroUnPadding(origData)
	return originalArr, nil
}

// TripleDesEncrypt2Str 3DES加密字符串，返回base64处理后字符串
func TripleDesEncrypt2Str(originalText string, key []byte) (string, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	originalData := PKCS5Padding([]byte(originalText), block.BlockSize())
	// originalData = ZeroPadding(originalData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cipherArr := make([]byte, len(originalData))
	blockMode.CryptBlocks(cipherArr, originalData)
	cipherText := base64.StdEncoding.EncodeToString(cipherArr)
	return cipherText, nil
}

// TripleDesDecrypt2Str 3DES解密base64处理后的加密字符串，返回明⽂字符串
func TripleDesDecrypt2Str(cipherText string, key []byte) (string, error) {
	cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	originalArr := make([]byte, len(cipherArr))
	blockMode.CryptBlocks(originalArr, cipherArr)
	originalArr = PKCS5UnPadding(originalArr)
	// origData = ZeroUnPadding(origData)
	return string(originalArr), nil
}
