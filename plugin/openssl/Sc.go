package openssl

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// SCEncrypt 对称加密（Symmetric Cryptography）
func SCEncrypt(originalBytes, key []byte, scType string) ([]byte, error) {
	// 1、实例化密码器block（参数为密钥）
	var err error
	var block cipher.Block
	switch scType {
	case "des":
		block, err = des.NewCipher(key)
	case "3des":
		block, err = des.NewTripleDESCipher(key)
	case "aes":
		block, err = aes.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 2、对明文进行填充（参数为原始字节切片和密码对象的区块个数）
	paddingBytes := PKCS5Padding(originalBytes, blockSize)

	// 3、实例化加密模式（参数为密码对象和密钥）
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])

	// 4、对填充字节后的明文进行加密（参数为加密字节切片和填充字节切片）
	cipherBytes := make([]byte, len(paddingBytes))
	blockMode.CryptBlocks(cipherBytes, paddingBytes)
	return cipherBytes, nil
}

// SCDecrypt 对称解密
func SCDecrypt(cipherBytes, key []byte, scType string) ([]byte, error) {
	// 1、实例化密码器block（参数为密钥）
	var err error
	var block cipher.Block
	switch scType {
	case "des":
		block, err = des.NewCipher(key)
	case "3des":
		block, err = des.NewTripleDESCipher(key)
	case "aes":
		block, err = aes.NewCipher(key)
	}
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()

	// 2、实例化解密模式（参数为密码对象和密钥）
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])

	// 3、对密文进行解密（参数为填充字节切片和加密字节切片）
	paddingBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(paddingBytes, cipherBytes)

	// 4、去除填充的字节（参数为填充切片）
	originalBytes := PKCS5UnPadding(paddingBytes)
	return originalBytes, nil
}

// SCEncryptString 封装字符串对称加密
func SCEncryptString(originalText, key, scType string) (string, error) {
	cipherBytes, err := SCEncrypt([]byte(originalText), []byte(key), scType)
	if err != nil {
		return "", err
	}
	// base64 编码（encoded）
	base64str := base64.StdEncoding.EncodeToString(cipherBytes)
	return base64str, nil
}

// SCDecryptString 封装字符串对称解密
func SCDecryptString(cipherText, key, scType string) (string, error) {
	// base64 解码（decode）
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherText)
	cipherBytes, err := SCDecrypt(cipherBytes, []byte(key), scType)
	if err != nil {
		return "", err
	}
	return string(cipherBytes), nil
}

// ScPKCS5Padding 末尾填充字节
func ScPKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize // 要填充的值和个数
	slice1 := []byte{byte(padding)}            // 要填充的单个二进制值
	slice2 := bytes.Repeat(slice1, padding)    // 要填充的二进制数组
	return append(data, slice2...)             // 填充到数据末端
}

// ScPKCS5UnPadding 去除填充的字节
func ScPKCS5UnPadding(data []byte) []byte {
	unPadding := data[len(data)-1]                // 获取二进制数组最后一个数值
	result := data[:(len(data) - int(unPadding))] // 截取开始至总长度减去填充值之间的有效数据
	return result
}

// ScZerosPadding 末尾填充0
func ScZerosPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize // 要填充的个数
	slice1 := []byte{0}                        // 要填充的单个0数据
	slice2 := bytes.Repeat(slice1, padding)    // 要填充的0二进制数组
	return append(data, slice2...)             // 填充到数据末端
}

// ScZerosUnPadding 去除填充的0
func ScZerosUnPadding(data []byte) []byte {
	return bytes.TrimRightFunc(data, func(r rune) bool { // 去除满足条件的子切片
		return r == 0
	})
}
