package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Create a slice by make function
	var gpasSlice []float64
	gpasSlice = make([]float64, 3)
	gpasSlice[0] = 4.1
	gpasSlice[1] = 4.5
	gpasSlice[2] = 3.9
	fmt.Println(gpasSlice, reflect.TypeOf(gpasSlice))

	// Create a slice by slice literal
	// gpasSlice := []float64{4.1, 4.5, 3.9} // slice literal
	// fmt.Println(gpasSlice, reflect.TypeOf(gpasSlice))

	// Create a slice by slicing an existing array
	// gpas := [5]float64{3.5, 4.1, 4.5, 3.9, 4.23} //array := array literal 중요 외우기
	// fmt.Println(gpas, reflect.TypeOf(gpas))
	// gpasSlice := gpas[1:4] // slice := slicing array
	// fmt.Println(gpasSlice, reflect.TypeOf(gpasSlice))
}
