package main

import "fmt" // c언어의 #include <stdio.h>

func main() {
	// var i int
	// i = 55

	//var i int = 55

	i := 55 //:=단축 연산자 var 안써도 괜찮은 타입도 안적음(뒤에 오는걸로 추적)
	var f float32 = 1.92

	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	fmt.Printf("i is %d\n", i)

}
