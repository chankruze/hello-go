/*
Author: chankruze (chankruze@gmail.com)
Created: Tue Mar 08 2022 17:59:03 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

const Token string = "adssdrewrfdsvxegv" // public (first letter cap)

func main() {
	var username string = "chandan"
	fmt.Println("username =", username)
	fmt.Printf("username is of type: %T\n", username)
	fmt.Println("")

	var age int = 21
	fmt.Println("age =", age)
	fmt.Printf("age is of type: %T\n", age)
	fmt.Println("")

	var isMarried bool = true
	fmt.Println("isMarried =", isMarried)
	fmt.Printf("username is of type: %T\n", isMarried)
	fmt.Println("")

	var gender byte = 'M' // byte == uint8
	fmt.Println("gender =", gender)
	fmt.Printf("gender is of type: %T\n", gender)
	fmt.Println("")

	var height float32 = 5.12345678
	fmt.Println("height =", height)
	fmt.Printf("height is of type: %T\n", height)
	fmt.Println("")

	var height2 float64 = 5.12345678
	fmt.Println("height2 =", height2)
	fmt.Printf("height2 is of type: %T\n", height2)
	fmt.Println("")

	// default values
	var defaultInt int
	fmt.Println("defaultInt =", defaultInt)
	fmt.Println("")

	var defaultFloat32 float32
	fmt.Println("defaultFloat32 =", defaultFloat32)
	fmt.Println("")

	var defaultFloat64 float64
	fmt.Println("defaultFloat64 =", defaultFloat64)
	fmt.Println("")

	var defaultString string
	fmt.Println("defaultString =", defaultString)
	fmt.Println("")

	// implicit value
	var github = "https://github.com/chankruze"
	fmt.Println("github =", github)
	// github = 3  // invalid
	fmt.Println("")

	// no var style (only inside a method)
	numOfFans := 465
	fmt.Println("numOfFans =", numOfFans)
	fmt.Println("")

	fmt.Println("Token =", Token)
}
