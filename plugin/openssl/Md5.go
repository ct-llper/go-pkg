package openssl

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/*

MD5信息摘要算法是一种被广泛使用的密码散列函数，可以产生出一个128位（16进制，32个字符）的散列值（hash value），用于确保信息传输完整一致。

*/
