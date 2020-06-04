package main
// 包的声明语句， main包包含了main 函数， main的包名称和所在的目录名称不相同
// G0 语言程序组成部分
//导入需要使用的包
import (
	
	"fmt"
	"01helloWorld/day04/math"
	"01helloWorld/day04/math2"
//	"time"
//	"os"
)
// 如果是多个包，放到一起
//只能导入在代码使用的包，如果包导入了，没有使用，导致语法错误


//函数 
func main()  {
	fmt.Println("HelloWorld!!!")

	add := math.Add(100,100)
	sub := math.Sub(100,100)
	mul := math2.Mul(100,100)
	dev := math2.Dev(100,100)
	fmt.Println(add,sub,mul,dev)
}