package panic_recover

import "log"

func main(){
	panic(404)
}

func demo_recover()  {
	if r := recover(); r != nil{
		log.Print("Running error caught: %v", r)
	}
}