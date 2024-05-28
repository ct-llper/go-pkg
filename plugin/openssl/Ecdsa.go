package openssl

//
//import (
//	"crypto/ecdsa"
//	"crypto/elliptic"
//	"crypto/rand"
//	"crypto/sha256"
//	"encoding/hex"
//	"fmt"
//	"log"
//	"math/big"
//	"testing"
//)
//
//// NewKeyPair ⽣成私钥和公钥，⽣成的私钥为结构体ecdsa.PrivateKey的指针
//func NewKeyPair() (ecdsa.PrivateKey, []byte) {
//	//⽣成secp256椭圆曲线
//	curve := elliptic.P256()
//	//产⽣的是⼀个结构体指针，结构体类型为ecdsa.PrivateKey
//	private, err := ecdsa.GenerateKey(curve, rand.Reader)
//	if err != nil {
//		log.Panic(err)
//	}
//	fmt.Printf("私钥：%x\n", private)
//	fmt.Printf("私钥X：%x\n", private.X.Bytes())
//	fmt.Printf("私钥Y：%x\n", private.Y.Bytes())
//	fmt.Printf("私钥D：%x\n", private.D.Bytes())
//	//x坐标与y坐标拼接在⼀起，⽣成公钥
//	pubKey := append(private.X.Bytes(), private.Y.Bytes()...)
//	//打印公钥，公钥⽤16进制打印出来⻓度为128，包含了x轴坐标与y轴坐标。
//	fmt.Printf("公钥：%x \n", pubKey)
//	return *private, pubKey
//}
//
//// MakeSignatureDerString ⽣成签名的DER格式
//func MakeSignatureDerString(r, s string) string {
//	// 获取R和S的⻓度
//	lenSigR := len(r) / 2
//	lenSigS := len(s) / 2
//	// 计算DER序列的总⻓度
//	lenSequence := lenSigR + lenSigS + 4
//	// 将10进制⻓度转16进制字符串
//	strLenSigR := DecimalToHex(int64(lenSigR))
//	strLenSigS := DecimalToHex(int64(lenSigS))
//	strLenSequence := DecimalToHex(int64(lenSequence))
//	// 拼凑DER编码
//	derString := "30" + strLenSequence
//	derString = derString + "02" + strLenSigR + r
//	derString = derString + "02" + strLenSigS + s
//	derString = derString + "01"
//	return derString
//}
//
//func DecimalToHex(n int64) string {
//	if n < 0 {
//		log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
//		return ""
//	}
//	if n == 0 {
//		return "0"
//	}
//
//	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
//	s := ""
//	for q := n; q > 0; q = q / 16 {
//		m := q % 16
//		if m > 9 && m < 16 {
//			m = hex[m]
//			s = fmt.Sprintf("%v%v", string(m), s)
//			continue
//		}
//		s = fmt.Sprintf("%v%v", m, s)
//	}
//	return s
//}
//
//// VerifySig 验证签名1
//func VerifySig(pubKey, message []byte, r, s *big.Int) bool {
//	curve := elliptic.P256()
//	//公钥的⻓度
//	keyLen := len(pubKey)
//	//前⼀半为x轴坐标，后⼀半为y轴坐标
//	x := big.Int{}
//	y := big.Int{}
//	x.SetBytes(pubKey[:(keyLen / 2)])
//	y.SetBytes(pubKey[(keyLen / 2):])
//	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
//	//根据交易哈希、公钥、数字签名验证成功。ecdsa.Verify func Verify(pub *PublicKey, hash[] byte, r * big.Int, s * big.Int) bool
//	res := ecdsa.Verify(&rawPubKey, message, r, s)
//	return res
//}
//
//// VerifySignature 验证签名2
//func VerifySignature(pubKey, message []byte, r, s string) bool {
//	curve := elliptic.P256()
//	//公钥的⻓度
//	keyLen := len(pubKey)
//	//前⼀半为x轴坐标，后⼀半为y轴坐标
//	x := big.Int{}
//	y := big.Int{}
//	x.SetBytes(pubKey[:(keyLen / 2)])
//	y.SetBytes(pubKey[(keyLen / 2):])
//	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
//	//根据交易哈希、公钥、数字签名验证成功。ecdsa.Verify func Verify(pub *PublicKey, hash[] byte, r * big.Int, s * big.Int) bool
//	rint := big.Int{}
//	sint := big.Int{}
//	rByte, _ := hex.DecodeString(r)
//	sByte, _ := hex.DecodeString(s)
//	rint.SetBytes(rByte)
//	sint.SetBytes(sByte)
//	//fmt.Println("------", rint.SetBytes(rByte))
//	//fmt.Println("------", sint.SetBytes(sByte))
//	res := ecdsa.Verify(&rawPubKey, message, &rint, &sint)
//	return res
//}
//
//// Test 验证过程
//func Test(t *testing.T) {
//	//1、⽣成签名
//	fmt.Println("1、⽣成签名-------------------------------")
//	//调⽤函数⽣成私钥与公钥
//	privKey, pubKey := NewKeyPair()
//	//信息的哈希
//	msg := sha256.Sum256([]byte("hello world"))
//	//根据私钥和信息的哈希进⾏数字签名，产⽣r和s
//	r, s, _ := ecdsa.Sign(rand.Reader, &privKey, msg[:])
//	//⽣成r、s字符串
//	fmt.Println("-------------------------------")
//	strSigR := fmt.Sprintf("%x", r)
//	strSigS := fmt.Sprintf("%x", s)
//	fmt.Println("r、s的10进制分别为：", r, s)
//	fmt.Println("r、s的16进制分别为：", strSigR, strSigS)
//	//r和s拼接在⼀起，形成数字签名的der格式
//	signatureDer := MakeSignatureDerString(strSigR, strSigS)
//	//打印数字签名的16进制显示
//	fmt.Println("数字签名DER格式为：", signatureDer)
//	fmt.Println()
//	//2、签名验证过程
//	fmt.Println("2、签名验证过程-------------------------------")
//	res := VerifySig(pubKey, msg[:], r, s)
//	fmt.Println("签名验证结果：", res)
//	res = VerifySignature(pubKey, msg[:], strSigR, strSigS)
//	fmt.Println("签名验证结果：", res)
//}
