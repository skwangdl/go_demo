package main

import "errors"
import "fmt"

//大写开头的变量与函数为共有，可以被其他包调用；小写开头的变量与函数只能在本包内调用
func Add(a int, b int)(ret int, err error){		//方法定义多个返回值及其类型
	if a < 0 || b < 0{
		err = errors.New("should be no-negative numbers")
		return
	}
	return a + b, nil	//方法支持多重返回值
}

//不定参数方法
func myFunc(args ...int){
	for _,arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()
}

//任意参数方法
func myFuncAnyType(args ...interface{}){
	for _,arg := range args {
		fmt.Print(arg)
	}
	fmt.Println()
}

//判断任意参数类型
func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func main(){
	fmt.Println(Add(2, 3))
	myFunc(1,2)
	myFunc(1,2,3)
	myFuncAnyType("a", 1, 2)
	MyPrintf("a", 1)
}