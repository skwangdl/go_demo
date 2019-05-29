package main

import "fmt"

func main()  {
	v1 := 10
	const PI float64 = 3.1415926
	const ZERO = 0.0
	// v2 := "kepler"
	// v3 := [10]int{1,2,3,4,5,6,7,8,9,10}		//数组
	// var v4 []int							//数组切片
	// var v5 struct{f int}					
	// var v6 *int								//指针
	// var v7 map[string]int					//map key为string, value为int类型
	// var v8 func(a int) int

	fmt.Println(v1)
	fmt.Println(PI)
	fmt.Println(ZERO)
	_,_,nickName := GetName()
	fmt.Println(nickName)
}

func GetName() (firstName, lastName, nickName string){
	return "May", "Chan", "chibi Maruko"
}