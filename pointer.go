package main

import "fmt"

func main() {
     a := 2
     b := &a //address of a
  /*address of a of value */  *b = 20
     fmt.Println(a)
}
