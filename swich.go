package main

import "fmt"

func canIDrinkAgain(age int) bool  {
	switch age {
	case 10: return false
	case 20:return true
	}
	return false
}

func main() {
   fmt.Println(canIDrinkAgain(20))
}
