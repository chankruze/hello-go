/*
Author: chankruze (chankruze@gmail.com)
Created: Mon Mar 14 2022 23:35:07 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to Math and Random")

	// bad math
	// var a int = 10
	// var b float64 = 20.5
	// fmt.Println("a + b = ", a+int(b))

	// good math
	var a int = 10
	var b float64 = 20.5
	fmt.Println("a + b = ", a+int(b))

	// random number (math/rand)
	// rand.Seed(time.Now().Unix())
	// fmt.Println("Random number: ", rand.Intn(5))

	// random number (crypto/rand)
	randNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number: ", randNum)
}
