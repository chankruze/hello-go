/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 17:23:05 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Luddo Kingo")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("You can move 1 spot")
	case 2:
		fmt.Println("You can move 2 spots")
	case 3:
		fmt.Println("You can move 3 spots")
		// fallthrough // is used to skip the next case
	case 4:
		fmt.Println("You can move 4 spots")
	case 5:
		fmt.Println("You can move 5 spots")
	case 6:
		fmt.Println("You can move 6 spots, you can also open/draw again")
	}
}
