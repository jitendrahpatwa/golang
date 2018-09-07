package main

import "fmt"

func main() {
	fmt.Println("Loop Conditions!")
	var a int = 15
	var b int
 	numbers := [6]int{1,2,3,5}

	for a := 0; a < 10; a++ {
		fmt.Printf("value of a: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("value of aaz: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("value of x = %d at %d\n", x,i)
	}
}	