package main

import (
	"fmt"
	"log"
)

func main() {

	/*================Sign begin================*/
	// data 待签名的数据
	// 转账地址：phenix2CkSyo9K5rrdw7aV4gkEAYNfAFjAi7pEgH
	// 接收地址：phenix2G9nhGgH8J5w8E2cF18DGiqxUq9Lf7p468
	// 转账金额：1.6
	data := []byte("phenix2CkSyo9K5rrdw7aV4gkEAYNfAFjAi7pEgH|phenix2G9nhGgH8J5w8E2cF18DGiqxUq9Lf7p468|1.6")
	//私钥
	priKey := []byte("lVKXDndWw1yJBuJXYNUxm0IA31dmOVQX")
	//签名
	sign, err := Sign(priKey, data)
	if err != nil {
		log.Fatalln(err)
	}
	/*================Sign end================*/

	/*================Verify begin================*/
	//公钥
	pubKey := "04884fa0ce7d1310ab87fbd2680a21959db648ff6771248f5e2fecc45179fdbd26039b5684f6cdf5fb4f2f288e12cb982a1b3fc84b112f3cbba1b4e47ac1e04a73"
	//验证签名
	flag, err := Verify(sign, data, pubKey)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(flag)
	/*================Verify end================*/

	//转账测试
	TranPost(data, sign)
}
