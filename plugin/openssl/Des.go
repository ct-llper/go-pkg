package openssl

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// DesEncrypt DES加密字节数组，返回字节数组
func DesEncrypt(originalBytes, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	originalBytes = PKCS5Padding(originalBytes, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherArr := make([]byte, len(originalBytes))
	blockMode.CryptBlocks(cipherArr, originalBytes)
	return cipherArr, nil
}

// DesDecrypt DES解密字节数组，返回字节数组
func DesDecrypt(cipherBytes, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	originalText := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(originalText, cipherBytes)
	originalText = PKCS5UnPadding(originalText)
	return originalText, nil
}

// DesEncryptString DES加密⽂本，返回加密后⽂本
func DesEncryptString(originalText string, key []byte) (string, error) {
	cipherArr, err := DesEncrypt([]byte(originalText), key)
	if err != nil {
		return "", err
	}
	base64str := base64.StdEncoding.EncodeToString(cipherArr)
	return base64str, nil
}

// DesDecryptString 对加密⽂本进⾏DES解密，返回解密后明⽂
func DesDecryptString(cipherText string, key []byte) (string, error) {
	cipherArr, _ := base64.StdEncoding.DecodeString(cipherText)
	cipherArr, err := DesDecrypt(cipherArr, key)
	if err != nil {
		return "", err
	}
	return string(cipherArr), nil
}

// ---------------DES ECB加密--------------------

// DesECBEncrypt DES ECB加密
// data: 明文数据
// key: 密钥字符串
// 返回密文数据
func DesECBEncrypt(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	// pkcs5填充
	data = PKCS5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out
}

// DesECBDecrypt DES ECB解密
//---------------DES ECB解密--------------------
// data: 密文数据
// key: 密钥字符串
// 返回明文数据
func DesECBDecrypt(data, key []byte) []byte {
	//NewCipher创建一个新的加密块
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil
	}

	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		//Encrypt加密第一个块，将其结果保存到dst
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	// pkcs5填充
	out = PKCS5UnPadding(out)

	return out
}

// DesCBCEncrypt
// ---------------DES CBC加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCBCEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = PKCS5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText
}

// DesCBCDecrypt
// ---------------DES CBC解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCBCDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.CryptBlocks(cryptText, data)
	// pkcs5填充
	cryptText = PKCS5UnPadding(cryptText)

	return cryptText
}

// DesCTREncrypt
// ---------------DES CTR加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCTREncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = PKCS5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCTR(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesCTRDecrypt
// ---------------DES CTR解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCTRDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCTR(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = PKCS5UnPadding(cryptText)

	return cryptText
}

// DesOFBEncrypt
// ---------------DES OFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesOFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = PKCS5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewOFB(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesOFBDecrypt
// ---------------DES OFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesOFBDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewOFB(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = PKCS5UnPadding(cryptText)

	return cryptText
}

// DesCFBEncrypt
// ---------------DES CFB加密--------------------
// data: 明文数据
// key: 密钥字符串
// iv:iv向量
// 返回密文数据
func DesCFBEncrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	// pkcs5填充
	data = PKCS5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCFBDecrypter(block, iv)
	blockMode.XORKeyStream(cryptText, data)
	return cryptText
}

// DesCFBDecrypt
// ---------------DES CFB解密--------------------
// data: 密文数据
// key: 密钥字符串
// iv:iv向量
// 返回明文数据
func DesCFBDecrypt(data, key, iv []byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil
	}

	blockMode := cipher.NewCFBEncrypter(block, iv)
	cryptText := make([]byte, len(data))
	blockMode.XORKeyStream(cryptText, data)

	// pkcs5填充
	cryptText = PKCS5UnPadding(cryptText)

	return cryptText
}

/*
DES加密算法，为对称加密算法中的一种。是以64比特的明文为一个单位来进行加密，超过64比特的数据，要求按固定的64比特的大小分组，分组有很多模式。DES使用的密钥长度为64比特，但由于每隔7个比特设置一个奇偶校验位，因此其密钥长度实际为56比特。奇偶校验为最简单的错误检测码，即根据一组二进制代码中1的个数是奇数或偶数来检测错误。

加密模式
ECB模式 全称Electronic Codebook模式，译为电子密码本模式
CBC模式 全称Cipher Block Chaining模式，译为密文分组链接模式
CFB模式 全称Cipher FeedBack模式，译为密文反馈模式
OFB模式 全称Output Feedback模式，译为输出反馈模式。
CTR模式 全称Counter模式，译为计数器模式。
填充方式
当明文长度不为分组长度的整数倍时，需要在最后一个分组中填充一些数据使其凑满一个分组长度。

NoPadding
API或算法本身不对数据进行处理，加密数据由加密双方约定填补算法。例如若对字符串数据进行加解密，可以补充\0或者空格，然后trim
PKCS5Padding
加密前：数据字节长度对8取余，余数为m，若m>0,则补足8-m个字节，字节数值为8-m，即差几个字节就补几个字节，字节数值即为补充的字节数，若为0则补充8个字节的8
解密后：取最后一个字节，值为m，则从数据尾部删除m个字节，剩余数据即为加密前的原文。
加密字符串为为AAA，则补位为AAA55555;加密字符串为BBBBBB，则补位为BBBBBB22；加密字符串为CCCCCCCC，则补位为CCCCCCCC88888888。
PKCS7Padding
PKCS7Padding 的填充方式和PKCS5Padding 填充方式一样。只是加密块的字节数不同。PKCS5Padding明确定义了加密块是8字节，PKCS7Padding加密快可以是1-255之间。

*/
