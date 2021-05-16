package main

import (
	"fmt"
	"strings"
)

func addition(a,b int) int  {
	return a*b
}
func lenAndUpper(name string) (int,string)  {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string)  {
	fmt.Println(words)
}

func main() {
   fmt.Println(addition(4,5))

   totalLength, upperName := lenAndUpper("Shubham")
   //totalLength,_ := ____
   fmt.Println(totalLength,upperName)

   repeatMe("Shubham","jaimin","jyot","niv")
}
