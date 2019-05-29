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
	const(				//iota被重设为0
		v2 = iota		//v2 == 0
		v3 = iota		//v3 == 1
		v4 = iota		//v4 == 2
	)
	const(
		a = 1 << iota	//a == 1 iota在每个const开头被重设为0
		b = 1 << iota	//b == 2
		c = 1 << iota	//c == 4
	)
	const(
		Sunday = iota	//Sunday == 0
		Monday			//Monday == 1
		Tuesday			//Tuesday == 2
		Wednesday		//Wednesday == 3
		Thursday		//Thursday == 4
		Friday			//Friday == 5
		Saturday		//Saturday == 6
	)

	fmt.Println(v1)
	fmt.Println(PI)
	fmt.Println(ZERO)
	_,_,nickName := GetName()
	fmt.Println(nickName)
	fmt.Println(v2, v3, v4)
	fmt.Println(a, b, c)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

}

func GetName() (firstName, lastName, nickName string){
	return "May", "Chan", "chibi Maruko"
}