/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 15:43:13 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// structs are reference types
	// structs are not value types
	// structs are not thread-safe
	// structs are not safe for concurrent use
	// structs are declared using the built-in make function
	// no iheritance in go

	// var person struct {
	// 	age  int
	// 	name string
	// }

	// person.age = 10
	// person.name = "chankruze"

	// fmt.Println(person)

	// create an user
	chankruze := User{"Chandan Kumar Mandal", "chankruze@gmail.com", 21, true}

	fmt.Println(chankruze)
	// %+v -> {Key: value}
	fmt.Printf("%+v\n", chankruze)
	fmt.Printf("Name: %+v\n", chankruze.Name)
}

// uppercase struct name is a convention
// uppercase field names is a convention
type User struct {
	Name     string
	Email    string
	Age      int
	Verified bool
}
