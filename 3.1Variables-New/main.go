package main

import (
	"fmt"
)

func main() {

	// Int type check - 1
	// Integer overflow. This occurs when an integer is incremented or decremented beyond its maximum or minimum value
	var overflow int32 = 2147483647
	overflow++
	fmt.Println(overflow)

	// Float type check - 1
	// Integer underflow. This occurs when an integer is decremented or incremented below its minimum or maximum value.
	var underflow int32 = -2147483648
	underflow--
	fmt.Println(underflow)

	// Float type check - 2
	// Precision errors. Floating point numbers are approximations of real numbers, and as such, they can lose precision.
	// This means that two floating point numbers that are equal to each other mathematically may not be equal to each other in Golang.
	// This is because the value of 0.1 cannot be represented exactly as a floating point number.
	var f float32 = 0.1
	var f1 float32 = 0.10000000000000001
	fmt.Println(f == f1)

	// Float type check - 3
	// NaN (Not a Number) values. NaN values are floating point numbers that represent an undefined value.
	// This can occur when a floating point operation is performed on invalid inputs, such as dividing by zero.
	// NaN values are often difficult to debug, because they can cause unexpected results.
	var x float32 = 1.0
	fmt.Println(x / 0)

}
