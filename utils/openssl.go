package utils

import (
	"bytes"
	"context"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"net/url"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// pkcs7Padding 补码
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// pkcs7UnPadding 去码
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])

	if length < unPadding {
		return []byte("")
	}
	return origData[:(length - unPadding)]
}

// EncryptCode	base64 加密
func EncryptCode(ctx context.Context, codeKey, data string) (err error, code string) {
	// 加密key
	k := []byte(codeKey)
	// 转成字节数组
	origData := []byte(data)
	// 分组秘钥
	block, _ := des.NewCipher(k)

	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补码
	origData = pkcs7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	crypt := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(crypt, origData)

	return nil, hex.EncodeToString(crypt)
}

// DecryptCode 解密
func DecryptCode(ctx context.Context, codeKey, data string) (err error, code string) {
	// 兼容
	cryptByte, err := DecodeHexOrBase64(data)
	if err != nil {
		return err, code
	}
	k := []byte(codeKey)
	// 分组秘钥
	block, err := des.NewCipher(k)
	if err != nil {
		return err, code
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 验证是否完整
	if len(cryptByte)%blockSize != 0 {
		return errors.New("crypto/cipher: input not full blocks"), ""
	}
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(cryptByte))
	// 解密
	blockMode.CryptBlocks(orig, cryptByte)
	// 去码
	orig = pkcs7UnPadding(orig)
	if len(orig) <= 0 {
		return errors.New(" decrypt code err "), ""
	}
	return nil, string(orig)
}

// DecodeHexOrBase64 is hex or base64
func DecodeHexOrBase64(data string) ([]byte, error) {
	content, err := url.QueryUnescape(data)
	if err != nil {
		return nil, err
	}
	dat := []byte(content)
	isHex := true
	for _, v := range dat {
		if v >= 48 && v <= 57 || v >= 65 && v <= 70 || v >= 97 && v <= 102 {
			// isHex = true
		} else {
			isHex = false
			break
		}
	}
	if isHex {
		d, err := hex.DecodeString(content)
		if len(d) == 0 || err != nil {
			return base64.RawStdEncoding.DecodeString(content)
		}
		return d, err
	} else {
		return base64.RawStdEncoding.DecodeString(content)
	}
}

// Base64Encode 编码
func Base64Encode(src []byte, baseCode string) []byte {
	var coder = base64.NewEncoding(baseCode)
	return []byte(coder.EncodeToString(src))
}

// Base64Decode 解码
func Base64Decode(src []byte, baseCode string) ([]byte, error) {
	var coder = base64.NewEncoding(baseCode)
	return coder.DecodeString(string(src))
}

// sha1Encode sha1 加密
func sha1Encode(str string) string {
	// 产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()
	// 写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(str))
	// 这个用来得到最终的散列值的字符切片。Sum 的参数可以用来都现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	// SHA1 值经常以 16 进制输出，例如在 git commit 中。使用%x 来将散列结果格式化为 16 进制字符串。
	fmt.Println(str)
	fmt.Printf("%x\n", bs)
	return fmt.Sprintf("%x", bs)
}

// RsaEncrypt 加密
func RsaEncrypt(origData, publicKey []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext, privateKey []byte) ([]byte, error) {
	// 解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	// 解析PKCS1格式的私钥
	prIv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, prIv, ciphertext)
}
