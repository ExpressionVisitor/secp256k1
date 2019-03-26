package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"strings"
)

// Sign 签名
// 参数：
//      priKey ：私钥
//      data   ：待签名的数据
// 返回值：
//      签名后的数据
func Sign(priKey, data []byte) ([]byte, error) {
	Md5Inst := md5.New()
	Md5Inst.Write(data)
	hash := Md5Inst.Sum([]byte(""))

	x, y := SECP256K1().ScalarBaseMult(priKey)
	sk := ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: SECP256K1(),
			X:     x,
			Y:     y,
		},
		D: new(big.Int).SetBytes(priKey),
	}
	r, s, err := ecdsa.Sign(rand.Reader, &sk, hash)
	if err != nil {
		return nil, err
	}

	sign := append(r.Bytes(), s.Bytes()...)
	if len(sign) != 64 {
		//fmt.Println(sign)
		return Sign(priKey, data)
	}
	return sign, nil
}

// Verify 验证签名
// 参数：
//      sign   ：签名数据
//      data   ：待签名的数据
//      pubKey ：公钥
// 返回值：
//      验证结果
func Verify(sign, data []byte, pubKey string) (bool, error) {
	Md5Inst := md5.New()
	Md5Inst.Write(data)
	hash := Md5Inst.Sum([]byte(""))

	var pk = strings.TrimPrefix(pubKey, "04")
	x, err := hex.DecodeString(pk[:64])
	if err != nil {
		return false, err
	}
	y, err := hex.DecodeString(pk[64:])
	if err != nil {
		return false, err
	}
	publicKey := ecdsa.PublicKey{
		Curve: SECP256K1(),
		X:     new(big.Int).SetBytes(x),
		Y:     new(big.Int).SetBytes(y),
	}

	if !ecdsa.Verify(&publicKey, hash, new(big.Int).SetBytes(sign[:32]), new(big.Int).SetBytes(sign[32:])) {
		return false, nil
	}

	return true, nil
}

// CreatePubKey 由私钥生成公钥
// 参数：
//      sk   ：16进制私钥
// 返回值：
//      公钥
func CreatePubKey(sk string) string {
	k, err := hex.DecodeString(sk)
	if err != nil {
		log.Fatalln(err)
	}
	x, y := SECP256K1().ScalarBaseMult(k)

	publicKey := ecdsa.PublicKey{
		Curve: SECP256K1(),
		X:     x,
		Y:     y,
	}

	pubKey := append([]byte{0x04}, publicKey.X.Bytes()...)
	pubKey = append(pubKey, publicKey.Y.Bytes()...)
	return hex.EncodeToString(pubKey)
}

// TranPost 转账
// 参数：
//      data   ：待签名的数据
//      sign   ：签名数据
func TranPost(data, sign []byte) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	// Add the other fields
	fw, err := w.CreateFormField("data")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = fw.Write(data); err != nil {
		log.Fatalln(err)
	}

	// Add the other fields
	fw, err = w.CreateFormField("sign")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err = fw.Write(sign); err != nil {
		log.Fatalln(err)
	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", "http://121.201.80.40:8888/kcoin/transign", &b)
	if err != nil {
		log.Fatalln(err)
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		log.Fatalln("StatusCode", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
