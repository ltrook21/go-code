package main

import (
	"fmt"
)

var i int = 42 // compiler := trick cannot be used outside of a function.

// variable blocks that allow me to group declared variables. Just for organization.
var (
	Username string = "Poopscoopa69"
	Email    string = "example@gmail.com"
	ID       int    = 3151234
)

var (
	databaseName string = "db"
	User         string = "user"
	PW           string = "pw"
)

func main() {
	// var i int = 42 // var name type
	//i := 42 // Same as line above, but := tells Go to figure var type on its own.

	var i int
	i = 42
	//var j int = 27
	//k := 99

	fmt.Println(i)
}
