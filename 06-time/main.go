/*
Author: chankruze (chankruze@gmail.com)
Created: Tue Mar 15 2022 00:15:58 GMT+0530 (India Standard Time)

Copyright (c) geekofia 2022 and beyond
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time and Math")

	t := time.Now()
	fmt.Println("Current time: ", t)
	// dd-mm-yyy -> 02-01-2006
	fmt.Println("Current time: ", t.Format("02-01-2006 Mon"))
	// mm-dd-yyy -> 01-02-2006
	fmt.Println("Current time: ", t.Format("01-02-2006 Mon"))
	// yyyy-mm-dd -> 2006-01-02
	fmt.Println("Current time: ", t.Format("2006-01-02 Mon"))
	// dd-mm-yyy HH:mm:ss -> 02-01-2006 15:04:05
	fmt.Println("Current time: ", t.Format("02-01-2006 Mon 15:04:05"))
	// micro seconds
	fmt.Println("Current time: ", t.Format("02-01-2006 Mon 15:04:05.000000"))
	// nano seconds
	fmt.Println("Current time: ", t.Format("02-01-2006 Mon 15:04:05.000000000"))
	// mm.dd.yyy -> 01.02.2006
	fmt.Println("Current time: ", t.Format("01.02.2006 Mon"))

	time.Sleep(100 * time.Millisecond)

	currentTime := time.Now()
	timeElapsed := currentTime.Sub(t)
	// time.Duration -> int64
	// time.Since(t) shorthand for time.Now().Sub(t).
	fmt.Println("Current time: ", timeElapsed, time.Since(t))

	// create a date
	createdDate := time.Date(2020, time.January, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println("Created date: ", createdDate)
	fmt.Println("Created date: ", createdDate.Format("01-02-2006 Mon"))
}
