package main

import "fmt"

type Integer int
func (a Integer) Less (b Integer) bool{
	return a < b
}

func (a *Integer) Add (b Integer){
	*a += b
}

func main(){
	var a Integer = 1
	fmt.Println(a.Less(2))
	a.Add(3)
	fmt.Println(a)
}