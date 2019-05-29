package main

import "fmt"

func main(){
	//Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理
	defer func(){//defer保证语句一定能够执行，执行顺序为后定义的defer语句先执行
		fmt.Println("C")
		if err:=recover();err!=nil{
			fmt.Println(err) // 这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
}

func f(){
	fmt.Println("a")
    panic(55)
    fmt.Println("b")
    fmt.Println("f")
}