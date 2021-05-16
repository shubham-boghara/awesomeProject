package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Iceberg", "chocolate"}
	shubham := person{name: "Shubham", age: 20, favFood: favFood}
	personName := shubham
	fmt.Println(shubham)
	fmt.Println(personName.name)
}
