package main

import (
	"unsafe"
	"fmt"
	"unicode/utf8"
)
var g int = 1000

//复合赋值
func varValue() {
	var i int 
	i+=10
	i+=10
	var f float32
	var s string
	s = "Hello"
	fmt.Println(g,i,f,s)
	x := 10
	y :=20
	fmt.Println(x,y)
	//多重赋值 -------> 不引入第三变量直接交换
	x,y = y,x
	fmt.Println(x,y)
}
// _ 标识符

//_,err = io.Copy(des,src)
//常量声明

func constDemo()  {
	const pi float32= 3.14
	const x float32 = 3.14 *2
	fmt.Println(pi,x)
	//批量声明
	const (
		a = 100
		b
		c = 0
		d
	)
	fmt.Println(a,b,c,d)
	const (
		e = iota 
		f
		g = 100+iota
		h
	)
	fmt.Println(e,f,g,h)
}

//数据类型
func DataType () {
	// bool 类型
	var flag bool
	fmt.Println(flag)
	flag = true
	fmt.Println(flag)
}
//
func IntDemo()  {
	// int 类型的长度
	var i int = 100
	fmt.Println(unsafe.Sizeof(i))
	a,b := 10,20

	fmt.Println(a+b,a-b,a*b,a/b,a%b)
	//在 GO 语言中 ++  -- 是表达式，不是语句
	a++
	i=a
	fmt.Println(i,a)
}
// 各种不同类型的变量与零值的比较<高质量的C++>编程  林锐，附录B。
func FloatDemo()  {
	//float32 与 float64 之间需要强制转换
	var f1 float32 = 12.5
	var f2 float64 =0.05
	fmt.Println(f1,f2)
	f1 = float32(f2)
	f2 = float64(f1)
	fmt.Println(f1,f2)
	f3 := 1.99E8
	f4 := 1.99e-3
	fmt.Println(f3,f4)
	//复数
	c1 := 10 + 0i
	fmt.Println(c1,imag(c1),real(c1))
}

func StringDemo()  {
	s1 := "Hello"
	s2 := "World"
	s3 := s1 + s2
	fmt.Println(s3,len(s3))
	//取字符和取子串
	fmt.Println(s3[0:5])
	//字符串中的字符不能被改变，但是可以给变量重新赋值
	//字符串遍历

	for i:=0 ;i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	fmt.Println(utf8.RuneCountInString(s3))
	//按照字符遍历
	for _,v := range s3 {
		fmt.Println(string(v))
	}
	// unicode 与 utf-8、utf-16 之间的关系
	//unicode 就是一种字符编码。

	//原生字符串字面值

	str := `
	窗前名月光
	疑似地上霜
	举头望明月
	低头思故乡
	`
	fmt.Println(str)
}


//指针演示
func PointerDemo()  {
	var i int = 100
	var pi *int 
	fmt.Println(pi)
	pi=&i
	fmt.Println(pi,i,*pi)

	var ppi **int
	ppi = &pi
	fmt.Println(ppi,*ppi, &pi,pi,i,*pi,**ppi)
}
type Score int
type Age byte
func TypeDemo()  {
	var s Score = 98
	fmt.Println(s)
	var a Age =20
	fmt.Println(a)
	//s = i // 不能直接赋值
	var i int = 100
	s = Score(i)
	fmt.Println(s)

}

func Stdio() {
	//标准输入
	var a,b int
	//fmt.Scan(&a,&b)
	fmt.Scanln(&a,&b)
	//标准输出
	fmt.Println("a = ",a,"b = ",b)
}
func main()  {
	var a int = 100
	b := 3.14
	var s string = "Hello"
	var p = "QQS"

	i,j,k := 10,20,30

	fmt.Println(a,b,s,p,i,j,k)
	varValue()
	constDemo()
	DataType()
	IntDemo()
	FloatDemo()
	StringDemo()
	PointerDemo()
	TypeDemo()
	Stdio()
}
