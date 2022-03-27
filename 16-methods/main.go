/*
Author: chankruze (chankruze@gmail.com)
Created: Mon Mar 28 2022 00:52:29 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// methods are functions that are associated with a struct
	// methods are defined using the `func` keyword
	chanak := Person{Name: "Chanakya"}
	chanak.Greet()
	chanak.NewAge(26)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("Hello", p.Name)
}

func (p Person) NewAge(age int) {
	p.Age = age
}
