/*
Author: chankruze (chankruze@gmail.com)
Created: Mon Mar 14 2022 19:10:53 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Pizza Hut")
	fmt.Println("Please rate your experience (1-5):")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for your rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	if numRating < 1 || numRating > 5 {
		fmt.Println("Rating should be between 1 and 5")
		return
	}

	fmt.Println("Added 1 to your rating:", numRating+1)
}
