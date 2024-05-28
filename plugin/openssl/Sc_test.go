package openssl

import (
	"fmt"
	"testing"
)

func Test_Sc(t *testing.T) {
	// DES密钥
	key := "12345678" // 占8字节
	// 3DES密钥
	// key = "abcdefgh012345678" // 占24字节
	// AES密钥。key长度：16,24，32 bytes 对应 AES-128，AES-192，AES-256
	// key = "01234567abcdefgh" // 占16字节

	keyBytes := []byte(key)
	str := "a"
	cipherArr, err := SCEncrypt([]byte(str), keyBytes, "des")
	if err != nil {
		panic(err)
	}
	fmt.Printf("加密后的字节数组：%v\n", cipherArr)
	fmt.Printf("加密后的字节数组：%x\n", cipherArr)
	originalArr, _ := SCDecrypt(cipherArr, keyBytes, "des")
	fmt.Println(string(originalArr))

	fmt.Println("------------------------------------------------------")

	str = "字符串加密aa12"
	cipherText, _ := SCEncryptString(str, key, "des")
	fmt.Println(cipherText)

	originalText, _ := SCDecryptString(cipherText, key, "des")
	fmt.Println(originalText)
}
