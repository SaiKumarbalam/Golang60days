package main

import "fmt"

func main() {
	/*
		Operator	  Name	    			    Description
		&				AND					    Sets each bit to 1 if both bits are 1
		|				OR					 	Sets each bit to 1 if one of two bits is 1
		^				XOR						Sets each bit to 1 if only one of two bits is 1
		<<				Zero fill left shift	Shift left by pushing zeros in from the right and let the leftmost bits fall off
		>>				Signed right shift		Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off

	*/

	x := 3
	y := 4
	result := 0

	result = (x & y)
	fmt.Println(x, "&", y, "=", result)

	result = (x | y)
	fmt.Println(x, "&", y, "=", result)

	result = (x ^ y)
	fmt.Println(x, "&", y, "=", result)

	result = (x << y)
	fmt.Println(x, "&", y, "=", result)

	result = (x >> y)
	fmt.Println(x, "&", y, "=", result)

}
