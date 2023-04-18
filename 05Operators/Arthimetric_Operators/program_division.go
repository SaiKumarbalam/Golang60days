package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	for i := 10; i >= 0; i-- {
		fmt.Println(i)
	}
}
