package main

import "fmt"

func man() {
	var sub float32 = 30.69
	var vol *float32

	var b int = 30
	var c *int

	var str string = "subbu"
	var str1 *string

	fmt.Println(sub)
	fmt.Println(b)
	fmt.Println(str)

	fmt.Println(vol)
	fmt.Println(c)
	fmt.Println(str1)
}
