/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 13:11:45 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// declaring an array of integers
	var arr [5]int

	// initializing the array
	arr[0] = 10
	arr[1] = 20
	arr[4] = 50

	// printing the array
	fmt.Println("Array Length:", len(arr))
	fmt.Println("Array:", arr)

	// printing the array using range
	for i, v := range arr {
		fmt.Printf("Value at index %d is %d\n", i, v)
	}

	// declaring and initializing an array of vegitables
	var vegList = [5]string{"potato", "tomato", "onion"}

	// printing the array
	fmt.Println("Veg List Length:", len(vegList))
	fmt.Println("Veg List:", vegList)

	// printing the array using range
	for i, v := range vegList {
		fmt.Printf("Value at index %d is %s\n", i, v)
	}

	// declaring and initializing an array of fruits
	var fruitList = []string{"mango", "apple", "strawberry"}

	// printing the array
	fmt.Println("Fuit List Length:", len(fruitList))
	fmt.Println("Fuit List:", fruitList)

	// printing the array using range
	for i, v := range fruitList {
		fmt.Printf("Value at index %d is %s\n", i, v)
	}
}
