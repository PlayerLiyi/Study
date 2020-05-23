package main

import "fmt"


func main() {

//bool类型 

	var b1 bool
	b1 = true
	fmt.Printf("%T %t\n",b1,b1)
    b2 := false
    fmt.Printf("%T %t\n",b2,b2)

	//多个byte 的结合
	//dingyi s1
	var s1 string ;
	s1 = "hello"
	fmt.Printf("%T %s\n",s1,s1)

	/*
	数据类型转换 :Type(value) uint8(a1)
	*/
}