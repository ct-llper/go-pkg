package openssl

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/ripemd160"
	"io"
)

func Sha1(in string) string {
	// 产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。
	h := sha1.New()
	// 写入要处理的字节。
	h.Write([]byte(in))

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。
	return hex.EncodeToString(h.Sum(nil))
	// return hex.EncodeToString(h.Sum([]byte("")))

}

func Ripemd160(str string) string {
	h := ripemd160.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256(str string) string {
	h := sha256.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha256Sum(str string) string {
	sum := sha256.Sum256([]byte(str))
	hashCode2 := hex.EncodeToString(sum[:])

	fmt.Printf("%x\n", sum)
	return hashCode2
}

func Sha512(str string) string {
	h := sha512.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Sha521Sum(str string) string {
	sum := sha512.Sum512([]byte(str))
	fmt.Printf("%x\n", sum)
	return ""
}

/*

SHA-1可以生成一个被称为消息摘要的160位（20字节）散列值，散列值通常的呈现形式为40个十六进制数。


*/
