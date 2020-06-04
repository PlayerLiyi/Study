package main

import (
	"fmt"
)

func arrDemo() {
	//数组声明
	var arr1 [10]int
	fmt.Println(arr1)
	arr1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1, arr1[len(arr1)-1])

	//简短声明
	arr2 := [...]int{9, 10}
	arr2[0] = 11
	//arr2[10] =12 //数组越界
	fmt.Println(arr2)
	//4数组遍历
	for i := 0; i < len(arr1); i++ {
		fmt.Print(arr1[i], " ")
	}
	for _, v := range arr1 {
		fmt.Print(v, " ")
	}
	// 数组作为函数的参数, 任然是实参
	arrasParam(arr1)
	fmt.Println(arr1[0])
}

//数组作为函数的参数
func arrasParam(arr [10]int) {
	arr[0] = 10
}

//切片 修改数组
//切片的基本使用(本质是对数组的引用)
func sliceDemo() {
	//基于数组创建切片
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:5] //左闭右开
	fmt.Println(slice1)
	//基于切片创建切片
	slice2 := slice1[:2]
	fmt.Println(slice2)
	slice2[0] = 20
	fmt.Println(arr)
	fmt.Println(slice1)

	fmt.Println(slice2)
	//使用make 函数创建切片
	slice3 := make([]int, 5, 10)
	fmt.Println(len(slice3), cap(slice3), slice3)
	slice4 := make([]int, 5, 10)
	fmt.Println(len(slice4), cap(slice4), slice3)
	//fmt.Println(slice4[5]) //切片不能越界

}
func main() {
	arrDemo()
	sliceDemo()
}
