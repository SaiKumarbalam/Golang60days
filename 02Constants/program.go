package main

import "fmt"

func main() {
	//Constants are the fixed values that cannot be changed once declared.
	//In Go, we use the const keyword to create constant variables.
	const lightSpeed = 256366
	fmt.Println(lightSpeed)

	//lightSpeed = 5623663 //  Error Code this cannot be changed
	fmt.Println(lightSpeed)

}
