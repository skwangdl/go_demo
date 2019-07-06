package main

import "fmt"

func main(){

	//条件语句
	value := 1
	if value < 2{
		fmt.Println("a")
	}else {
		fmt.Println("b")
	}

	//选择语句
	switch value {
	case 0:
		fmt.Println(value)
	case 1:
		fmt.Println(value)
	}

	sum := 0
	for{			//go不支持while与do-while语句，使用for代替进行无限循环	
		sum ++
		if sum > 100{
			fmt.Println(sum)
			break
		}
	}

	OuterLoop:		//go break跳出固定循环，continue与Java类似
    for i := 0; i < 2; i++ {
        for j := 0; j < 5; j++ {
            switch j {
            case 2:
                fmt.Println(i, j)
                break OuterLoop
            case 3:
                fmt.Println(i, j)
                break OuterLoop
            }
        }
	}
	
	i := 0
	HERE:			//goto 标记位置
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE	//go goto语法，跳转到标记位置HERE
	}
}