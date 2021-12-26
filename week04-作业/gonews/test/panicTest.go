package main

import "fmt"

func main() {

	defer func() {
		fmt.Println("a")
	}()

	var b interface{}

	fmt.Println(b.(int))

}
