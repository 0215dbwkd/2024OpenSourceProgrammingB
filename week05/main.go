package main

import (
	"fmt" // c언어의 #include <stdio.h>
	"math"
	"reflect"
	"strings"
)

func main() {
	// var i int
	// i = 55

	//var i int = 55

	var f float64 = 1.92

	//f := 1.92
	i := "55" //:=단축 연산자 var 안써도 괜찮은 타입도 안적음(뒤에 오는걸로 추적)

	fmt.Println(strings.Title("kim inha"))
	fmt.Printf("%f %f %f\n", f, math.Ceil(f), math.Floor(f))
	fmt.Println(reflect.TypeOf(f), reflect.TypeOf(i))
	fmt.Println("f is", f)
	fmt.Println("i is", i)
	fmt.Println("i is", i)
	fmt.Print("i is ", i, "\n")
	//fmt.Printf("i is %d\n", i)
	fmt.Printf("i is %s\n", i)

}
