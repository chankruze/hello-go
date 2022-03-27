/*
Author: chankruze (chankruze@gmail.com)
Created: Sun Mar 27 2022 15:25:22 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import "fmt"

func main() {
	// maps are reference types
	// maps are not value types
	// maps are not thread-safe
	// maps are not safe for concurrent use
	// maps are declared using the built-in make function

	// make(map[key-type]value-type)
	languages := make(map[string]string)

	// add elements to the map
	languages["en"] = "English"
	languages["hi"] = "Hindi"
	languages["es"] = "Spanish"
	languages["fr"] = "French"
	languages["de"] = "German"
	languages["it"] = "Italian"
	languages["pt"] = "Portuguese"

	// print the map
	fmt.Println(languages)

	// get the value of a key
	fmt.Println(languages["en"])

	// delete a key-value pair
	delete(languages, "en")

	// check if a key exists
	if _, ok := languages["en"]; ok {
		fmt.Println("key exists")
	} else {
		fmt.Println("key does not exist")
	}

	// iterate over the map
	fmt.Println("\n# iterating over the map")
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// iterate over the keys
	fmt.Println("\n# iterating over the keys")
	for key := range languages {
		fmt.Println(key)
	}

	// iterate over the values
	fmt.Println("\n# iterating over the values")
	for _, value := range languages {
		fmt.Println(value)
	}
}
