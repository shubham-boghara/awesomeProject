package main

import "fmt"

func main() {
	shubham := map[string]string{"name": "shubham", "age": "12"}
	fmt.Println(shubham)
	for _, value := range shubham {
		fmt.Println(value)
	}

}
