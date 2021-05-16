package main

import "fmt"

func main() {
    names := []string{"nico","shubham","dal"}
    names = append(names,"foo")
    fmt.Println(names)
}
