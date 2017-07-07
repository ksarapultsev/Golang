package main

import "fmt"

var V string

func main() {
	fmt.Println("Hello World")
	fmt.Println(V)
}

func ReturnInt() int {
	var i int = 1
	return i
}

func ReturnFloat() float32 {
	var s float32 = 1.1
	return s
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	var aarray = [3]int{1, 2, 3}
	return aarray[0:3]
}
