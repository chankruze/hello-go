/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 12:47:48 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// 1. declare a pointer
	var p1 *int
	// print the value to the pointer
	fmt.Printf("Value at %p is %#v\n", p1, p1)

	// 2. reference to a variable
	num := 12
	// declare a pointer
	p2 := &num
	// print the value to the pointer
	fmt.Printf("Value at %p is %#v\n", p2, *p2)

	// 3. new: declare pointer to a variable
	p3 := new(int)
	// print the value of the pointer
	fmt.Printf("Value at %p is %#v\n", p3, *p3)
	// assign a value to the pointer
	*p3 = 10
	// print the value of the pointer
	fmt.Printf("Value at %p is %#v\n", p3, *p3)
}
