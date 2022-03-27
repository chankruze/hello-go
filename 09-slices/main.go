/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 13:20:16 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// creating a slice
	var fruitList = []string{"apple", "orange", "grapes", "banana"}

	fmt.Printf("Type of fruitList: %T\n", fruitList)

	// printing the length of the slice
	fmt.Printf("Length of fruitList: %d\n", len(fruitList))

	// printing the capacity of the slice
	fmt.Printf("Capacity of fruitList: %d\n", cap(fruitList))

	// printing the value of the slice
	fmt.Printf("Value of fruitList: %v\n", fruitList)

	// add more values to the slice
	fruitList = append(fruitList, "mango", "pineapple")

	fmt.Printf("Value of fruitList: %v\n", fruitList)

	// printing the length of the slice
	fmt.Printf("Length of fruitList: %d\n", len(fruitList))

	// printing the capacity of the slice
	fmt.Printf("Capacity of fruitList: %d\n", cap(fruitList))

	// slice of a slice
	// slice[x:y] -> x inclusive, y exclusive
	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:], "kiwi")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3], "papaya")
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 100
	highScores[1] = 900
	highScores[2] = 300
	highScores[3] = 400
	// highScores[4] = 700
	highScores = append(highScores, 700, 800, 500)

	fmt.Println(highScores)

	// sorting the int slice
	sort.Ints(highScores)

	fmt.Println(highScores)

	// check if the slice is sorted ?
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from the slice based on index
	var courses = []string{"go", "python", "java", "c++", "c#"}

	fmt.Println(courses)

	// remove the value at index 2
	index := 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
