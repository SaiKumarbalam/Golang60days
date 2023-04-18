package main

import "fmt"

func main() {
	/*
		Operator	Name				Description										    Example
		&&	     Logical And	Returns true if both statements are true		           x < y && x > z
		||	     Logical Or	    Returns true if one of the statements is true			   x < y || x > z
		!	    Logical Not	    Reverse the result, returns false if the result is true	   !(x == y && x > z)

	*/
	var x, y, z = 10, 20, 30
	fmt.Println(x < y && x > z)
	fmt.Println(x < y || x > z)
	fmt.Println(!(x == y && x > z))

}
