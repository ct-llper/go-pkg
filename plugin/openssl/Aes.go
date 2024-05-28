package openssl

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func AesEncrypt(src []byte, key []byte) (encrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

func AesDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func AesEncryptCBC(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecryptCBC(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

// AesCtrCrypt Ctr加密
func AesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {

	//1. 创建cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2. 创建分组模式，在crypto/cipher包中
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	//3. 加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

func AesEncryptCFB(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		//panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}
func AesDecryptCFB(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

func AesEncryptOFB(data []byte, key []byte) ([]byte, error) {
	data = PKCS7Padding(data, aes.BlockSize)
	block, _ := aes.NewCipher([]byte(key))
	out := make([]byte, aes.BlockSize+len(data))
	iv := out[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(out[aes.BlockSize:], data)
	return out, nil
}

func AesDecryptOFB(data []byte, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(key))
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("data is not a multiple of the block size")
	}

	out := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(out, data)

	out = PKCS7UnPadding(out)
	return out, nil
}

/*

密码学中的高级加密标准（Advanced Encryption Standard，AES），又称Rijndael加密法，是美国联邦政府采用的一种区块加密标准。这个标准用来替代原先的DES（Data Encryption Standard），已经被多方分析且广为全世界所使用。AES中常见的有三种解决方案，分别为AES-128、AES-192和AES-256。如果采用真正的128位加密技术甚至256位加密技术，蛮力攻击要取得成功需要耗费相当长的时间。

AES 有五种加密模式：

电码本模式（Electronic Codebook Book (ECB)）、
密码分组链接模式（Cipher Block Chaining (CBC)）、
计算器模式（Counter (CTR)）、
密码反馈模式（Cipher FeedBack (CFB)）
输出反馈模式（Output FeedBack (OFB)）
ECB模式
出于安全考虑，golang默认并不支持ECB模式。


ECB 电子密码本模式：相同的明文块产生相同的密文块，容易并行运算，但也可能对明文进行攻击。
CBC 加密分组链接模式：一块明文加密后和上一块密文进行链接，不利于并行，但安全性比ECB好，是SSL,IPSec的标准。
CFB 加密反馈模式：将上一次密文与密钥运算，再加密。隐藏明文模式，不利于并行，误差传递。
OFB 输出反馈模式：将上一次处理过的密钥与密钥运算，再加密。隐藏明文模式，不利于并行，有可能明文攻击，误差传递。

PKCS5Padding的填充方式是差多少字节就填数字多少；刚好每一不足16字节时，那么就会加一组填充为16.还有其他填充模式【Nopadding,ISO10126Padding】（不影响算法，加密解密时一致就行）

*/

//
//// AesEncrypt AES加密字节数组，返回字节数组
//func AesEncrypt(originalBytes, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	originalBytes = PKCS5Padding(originalBytes, blockSize)
//	// originalBytes = ZeroPadding(originalBytes, block.BlockSize())
//	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
//	cipherBytes := make([]byte, len(originalBytes))
//	// 根据CryptBlocks⽅法的说明，如下⽅式初始化crypted也可以
//	// crypted := originalBytes
//	blockMode.CryptBlocks(cipherBytes, originalBytes)
//	return cipherBytes, nil
//}
//
//// AesDecrypt AES解密字节数组，返回字节数组
//func AesDecrypt(cipherBytes, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	blockSize := block.BlockSize()
//	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
//	originalBytes := make([]byte, len(cipherBytes))
//	// origData := cipherBytes
//	blockMode.CryptBlocks(originalBytes, cipherBytes)
//	originalBytes = PKCS5UnPadding(originalBytes)
//	// origData = ZeroUnPadding(origData)
//	return originalBytes, nil
//}

// AesEncryptString AES加密⽂本，返回对加密后字节数组进⾏base64处理后字符串
//func AesEncryptString(originalText string, key []byte) (string, error) {
//	cipherBytes, err := AesEncrypt([]byte(originalText), key)
//	if err != nil {
//		return "", err
//	}
//	base64str := base64.StdEncoding.EncodeToString(cipherBytes)
//	return base64str, nil
//}
//
//// AesDecryptString AES解密⽂本，对Base64处理后的加密⽂本进⾏AES解密，返回解密后明⽂
//func AesDecryptString(cipherText string, key []byte) (string, error) {
//	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherText)
//	cipherBytes, err := AesDecrypt(cipherBytes, key)
//	if err != nil {
//		return "", err
//	}
//
//	return string(cipherBytes), nil
//}
