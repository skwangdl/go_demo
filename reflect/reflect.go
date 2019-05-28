package main

import (
	"fmt"
	"reflect"
)

type Bird struct{
	Name string
	LifeExpectance int
}

func (b *Bird) Fly(){
	fmt.Println("I am flying...")
}

func main(){
	sparrow := &Bird{"Sparrow", 3}		//创建实例
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++{
		f := s.Field(i)
		fmt.Println("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}