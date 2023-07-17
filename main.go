package main

import "fmt"

func add(x,y int) int{
	return x+y;
}

func main() {
	var x, y int = 1,5
	fmt.Println("Hello, Go in CodeSandbox!")
	fmt.Printf("Addition of %d and %d is: %d",x, y,add(x,y))
}
