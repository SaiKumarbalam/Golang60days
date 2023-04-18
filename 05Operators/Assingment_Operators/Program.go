package main

import "fmt"

func main() {
	/*
		Add and Assignment (+=)
		It is the combination of '+' and '=' operators,
		it adds the given value with the current value of the variable and assigns it to the variable.
		x +=y is equivalent to x = x+y
		*       *      *
		Subtract and Assignment (-=)	It is the combination of '-' and '=' operators,
		it subtracts the given value with the current value of the variable and assigns it to the variable.
		x -=y is equivalent to x = x-y
		*       *      *
		Multiply and Assignment (*=)	It is the combination of '*' and '=' operators,
		it multiplies the given value with the current value of the variable and assigns it to the variable.
		x *=y is equivalent to x = x*y
		*       *      *
		Divide and Assignment (/=)	It is the combination of '/' and '=' operators,
		it divides the given value with the current value of the variable and assigns it to the variable.
		x /=y is equivalent to x = x/y
		*       *      *
		Modulus and Assignment (%=)	It is the combination of '%' and '=' operators,
		it finds the remainder of the current value of the variable with the given value and assigns it to the variable.
		x %=y is equivalent to x = x%y
		*       *      *
		Bitwise AND Assignment (&=)	It is the combination of '&' and '=' operators,
		it performs the Bitwise AND operation with the current value of the variable and the given value and assigns the result to the variable.
		x &=y is equivalent to x = x&y
		*       *      *
		Bitwise OR Assignment (|=)	It is the combination of '|' and '=' operators,
		it performs the Bitwise OR operation with the current value of the variable and the given value and assigns the result to the variable.
		x |=y is equivalent to x = x|y
		*       *      *      *       *      *     *       *      *
		Bitwise XOR Assignment (^=)	It is the combination of '^' and '=' operators,
		it performs the Bitwise XOR operation with the current value of the variable and the given value and assigns the result to the variable.
		x ^=y is equivalent to x = x^y
		*       *      *     *       *      *     *       *      *
		Bitwise Left Shift Assignment (<<=)	It is the combination of '<<' and '=' operators,
		it performs the Bitwise Left-shift operation with the current value of the variable and the given value and assigns the result to the variable.
		x <<=y is equivalent to x = x<<y
		*       *      *    *       *      *     *       *      *
		Bitwise Right Shift Assignment (>>=)	It is the combination of '>>' and '=' operators,
		it performs the Bitwise Left-shift operation with the current value of the variable and the given value and assigns the result to the variable.
		x >>=y is equivalent to x = x>>y

	*/
	x := 5
	y := 3

	x += y
	fmt.Println("x:", x)

	x -= y
	fmt.Println("x:", x)

	x *= y
	fmt.Println("x:", x)

	x /= y
	fmt.Println("x:", x)

	x %= y
	fmt.Println("x:", x)

	x &= y
	fmt.Println("x:", x)

	x |= y
	fmt.Println("x:", x)

	x <<= y
	fmt.Println("x:", x)

	x >>= y
	fmt.Println("x:", x)
}
