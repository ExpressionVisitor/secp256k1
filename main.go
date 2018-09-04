package main

import (
	"fmt"
	"log"
)

func main() {

	/*================Sign begin================*/
	// data 待签名的数据
	// 转账地址：phenix3EHhkDgPmwMZrt7QtMH5QZ9hoVabEPZMm1
	// 接收地址：phenix38pLjBx7kTKmxNMN3tHGdLeexXX2Fio3pq
	// 转账金额：1.6
	data := []byte("phenix3EHhkDgPmwMZrt7QtMH5QZ9hoVabEPZMm1|phenix38pLjBx7kTKmxNMN3tHGdLeexXX2Fio3pq|1.6")
	//私钥
	priKey := []byte("Y8wfnrOAXtHKP3k3bsxCQvcmept6Vdvw")
	//签名
	sign, err := Sign(priKey, data)
	if err != nil {
		log.Fatalln(err)
	}
	/*================Sign end================*/

	/*================Verify begin================*/
	//公钥
	pubKey := "04d6ac902ec2a51cb196df8a8d1446c77f4779f32d01724ce1b0b0b010fd7942aef0bc9182dd42bf4e97ece78b8256135558353d6d7b1be780fc415794ab8ed714"
	//验证签名
	flag, err := Verify(sign, data, pubKey)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(flag)
	/*================Verify end================*/

	//转账测试
	//TranPost(data, sign)
}
