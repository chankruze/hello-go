/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 17:35:49 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for loop
	for i := 0; i < len(days); i++ {
		fmt.Println("📆", days[i])
	}

	// for loop with range (index)
	for i := range days {
		fmt.Println("📆", days[i])
	}

	// for loop with range (value)
	for _, day := range days {
		fmt.Println("📆", day)
	}

	rougueValue := 1

	for rougueValue <= 10 {
		if rougueValue == 4 {
			goto label
		}

		if rougueValue == 6 {
			break
		}

		fmt.Println("🎲", rougueValue)
		rougueValue++
	}

label:
	fmt.Println("goto example")
}
