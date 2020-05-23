
package main

import "fmt"

const (
	a1 = iota
	a2
	a3
)


func main()  {

	var (
		name string
		age int
	)
	name = "jack"
	age = 12
	fmt.Println("Hello World!!!")
	fmt.Println("name %s",name)
	fmt.Println("age %d",age)

	fmt.Println("a1 %d",a1)
	fmt.Println("a2 %d",a2)
	fmt.Println("a3 %d",a3)
}