package main

// You can group imports using ()
import (
	"fmt"
	"math"
)

// if both params of same datatype, you can define datatype at end instead of int x, int y
func add(x, y int) int {
	return x + y
}

/* for multiple returns, group return types like this. NOT RECOMENDED IN BIG FUNC */
func swap(name1, name2 string) (string, string) {
	return name2, name1
}

// As we defined x and y in return defination, x and y values exported by default
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func variab() {
	// No need to define type for intialized vars
	a, b, c := true, 1, "hello"
	var d int
	fmt.Println(a, b, c, d)
}

func main() {
	var x, y int = 1, 5
	fmt.Println("Hello, Go in CodeSandbox!")
	fmt.Println("Pi: ", math.Pi) //	p in Pi is capital as it is exported from math package
	fmt.Printf("Addition of %d and %d is: %d \n", x, y, add(x, y))
	fmt.Println(swap("john", "doe"))
	fmt.Println(split(22))
	variab()
}
