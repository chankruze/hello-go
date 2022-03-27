/*
Author: chankruze (chankruze@gmail.com)
Created: Mon Mar 28 2022 00:27:21 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// print the sum of two numbers
	fmt.Println(sum(10, 20))

	// print the sum of three numbers
	fmt.Println(sum(10, 20, 30))

	// print the sum of 4 numbers
	fmt.Println(sum(10, 20, 30, 40))

	marks := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	// print the sum of marks
	fmt.Println("Sum of marks:", sum(marks...))

	// multiple return values
	str, total := sum2(10, 20, 30)
	fmt.Println(str, total)

}

func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func sum2(nums ...int) (string, int) {
	total := 0

	for _, num := range nums {
		total += num
	}

	return "sum =", total
}
