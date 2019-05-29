package main

import "fmt"

func main(){
	var j int = 5
	a := func()(func()) {	//定义变量a的引用为闭包，闭包内的变量可以被闭包外更改，函数作为方法的返回值；
		// 只有内部的匿名函数才能访问变量i
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}