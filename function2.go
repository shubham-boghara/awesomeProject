package main

import (
	"fmt"
	"strings"
)

func lenAndUppercase(name string) (length int, uppercase string)  {
	defer fmt.Println("Done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
   totalLength,name := lenAndUppercase("Shubham")
   fmt.Println(totalLength,name)
}
