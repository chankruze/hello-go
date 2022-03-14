/*
Author: chankruze (chankruze@gmail.com)
Created: Mon Mar 14 2022 18:54:49 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give us a rating: ")

	// comma ok || ok err
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T\n", input)
}
