package main

import "fmt"

func canIDrink(age int) bool {
	if indiaAge := age +2; indiaAge>=18 {
		return true
	}
	return false
}
func main() {
  fmt.Println(canIDrink(5))
}
