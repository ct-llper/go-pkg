package openssl

import "bytes"

// 末尾填充字节

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize        // 要填充的值和个数
	padText := bytes.Repeat([]byte{byte(padding)}, padding) // 要填充的二进制数组
	return append(ciphertext, padText...)                   // 填充到数据末端
}

func PKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])   // 获取二进制数组最后一个数值
	return data[:(length - unPadding)] // 截取开始至总长度减去填充值之间的有效数据
}

// ZeroPadding 末尾填充0
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize // 要填充的个数
	padText := bytes.Repeat([]byte{0}, padding)      // 要填充的0二进制数组
	return append(ciphertext, padText...)            // 填充到数据末端
}

// ZeroUnPadding 去除填充的0
func ZeroUnPadding(data []byte) []byte {
	// 去除满足条件的子切片
	return bytes.TrimRightFunc(data, func(r rune) bool {
		return r == rune(0)
	})
}

//pkcs5补码算法
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

//pkcs5减码算法
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

//zero补码算法
func zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)
}

//zero减码算法
func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

// PKCS7Padding 补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
