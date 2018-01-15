package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("Square roots: ", math.Sqrt(4))
	fmt.Println("Random: ", rand.Intn(100))
}

func add(x, y float32) float32 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}

// func main() {
// 	// fmt.Println("Welcome to Go!")
// 	// foo()
// 	var num1, num2 float32 = 5.6, 2.3
// 	// num1, num2 := 5.6, 2.3
// 	// var num2 float64 = 2.3
// 	fmt.Println(add(num1, num2))

// 	w1, w2 := "Hey", "world"
// 	fmt.Println(multiple(w1, w2))

// 	var a int = 62
// 	var b float64 = float64(a)
// 	fmt.Println(b)
// }

func main() {
	x := 15
	a := &x // memory address of x
	fmt.Println(a)
	fmt.Println(*a) //what is stored at that memory address
	*a = 5
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
}
