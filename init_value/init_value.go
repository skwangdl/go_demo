package main

import "fmt"

func main()  {
	v1 := 10		//int整型
	const PI float64 = 3.1415926
	const ZERO = 0.0

	fmt.Println(v1)
	fmt.Println(PI)
	fmt.Println(ZERO)

	_,_,nickName := GetName()
	fmt.Println(nickName)

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
	fmt.Println(v2, v3, v4)
	const(
		a = 1 << iota	//a == 1 iota在每个const开头被重设为0
		b = 1 << iota	//b == 2
		c = 1 << iota	//c == 4
	)
	fmt.Println(a, b, c)
	const(
		Sunday = iota	//Sunday == 0
		Monday			//Monday == 1
		Tuesday			//Tuesday == 2
		Wednesday		//Wednesday == 3
		Thursday		//Thursday == 4
		Friday			//Friday == 5
		Saturday		//Saturday == 6
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)

	v_bool := true		//boolean 类型
	fmt.Println(v_bool)

	v_remiander := 5 % 3	//求余数	
	fmt.Println(v_remiander)

	i,j := 1,2		//多个赋值
	fmt.Println(i,j)
	fmt.Println(i == j)	//等值判断
	str1, str2 := "a", "a"
	fmt.Println(str1 == str2)	//字符串等值判断，为true

	var t1 int32	//两个不同类型的整型数值不能直接比较，但各种类型的整型变量都可以直接与字面常量(Iiteral)进行比较
	var t2 int64
	t1, t2 = 1, 2
	if t1 == 1 || t2 == 2{
		fmt.Println("t1 and t2 are equal.")
	}

	var value1 complex64	//构建复数
	value1 = 3.2 + 12i
	value2 := complex(3.2, 12)
	z1 := real(value1)
	x1 := imag(value1)
	z2 := real(value2)
	x2 := imag(value2)
	fmt.Println(z1, x1, z2, x2)

	var str string			//字符串变量
	str = "hello kepler"
	ch := str[0]
	fmt.Printf("The length of \"%s\" is %d \n", str, len(str))
	fmt.Printf("The first character of \"%s\" is %c.\n", str, ch)

	array := [10]int{1,2,3,4,5,6,7,8,9,10}	//创建并遍历数组，在函数体中无法修改传入的数组的内容
	for i, v := range array{
		fmt.Println("Array element[", i, "]=", v)
	}
	var mySlice []int = array[:5]			//基于数组创建数组切片
	fmt.Println("\nElements of mySlice: ")
	for _,v := range mySlice{
		fmt.Print(v, " ")
	}
	fmt.Println()

	mySlice1 := make([]int, 5)			//初始元素个数为5的切片数组，元素初始值为0
	mySlice2 := make([]int, 5, 10)		//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	mySlice3 := []int{1, 2, 3, 4, 5}	//直接创建并初始化包含5个元素的数组切片
	for i, v := range mySlice1 {
		fmt.Println("mySlice1[", i, "] =", v)
	}
	for i, v := range mySlice2 {
		fmt.Println("mySlice2[", i, "] =", v)
	}
	for i, v := range mySlice3 {
		fmt.Println("mySlice3[", i, "] =", v)
	}
	fmt.Println("len(mySlice):", len(mySlice2))	//数组切片中当前所存储的元素个数
	fmt.Println("cap(mySlice):", cap(mySlice2))	//数组切片分配的空间大小

	tempMap := make(map[string] string)		//创建Map
	tempMap["a"] = "key1";
	tempMap["b"] = "key2"
	tempStr1, ok1 := tempMap["a"]
	if ok1 {
		fmt.Println("Found str ", tempStr1)
	} else {
		fmt.Println("Did not find str in tempMap.")
	}
	delete(tempMap, "b")			//从Map内删除元素
	tempStr2, ok2 := tempMap["b"]
	if ok2 {
		fmt.Println("Found str ", tempStr2)
	} else {
		fmt.Println("Did not find str in tempMap, it is deleted.")
	}
}

func GetName() (firstName, lastName, nickName string){
	return "May", "Chan", "chibi Maruko"
}