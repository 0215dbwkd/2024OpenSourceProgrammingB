package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

func main() {
	//fmt.Sprintf("%d\n", rand.Intn(6)+1) //splint는 서식이 반영된 출력하는 대신
	r := fmt.Sprintf("%d\n", rand.Intn(6)+1)
	fmt.Println(reflect.TypeOf(r))
	fmt.Printf("%T\n", r)

	i := 1
	for i <= 100 {
		fmt.Printf("%3d점\n", i) //d앞에 2를 붙이면 10의 자리에 맞춰저서 나옴
		i = i + 1
	}
}
