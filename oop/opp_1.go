package main

import "fmt"

type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64{
	return r.height * r.width
}

func main(){
	rect := Rect{width:10, height:20}
	fmt.Println(rect.Area())
}