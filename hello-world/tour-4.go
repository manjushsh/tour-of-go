package main

import (
	"fmt"
)

func main() {
	basicPointer()
	structEx()
	arraysGo()
}

func basicPointer() {
	val := 54
	point := &val
	*point = 32
	fmt.Println(point, *point)
}

type Person struct {
	Name string
	Age  byte
}

func structEx() {
	person1 := Person{"John Doe", 54}
	pointer := &person1
	pointer.Age = 64
	fmt.Println(*pointer)
}

func arraysGo() {
	// Fix sized
	a1 := [2]int{}
	a1[0] = 1
	a1[1] = 3
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(a1, a1[0], a1[1], primes)

	// Dynamic / Slices
	
}
