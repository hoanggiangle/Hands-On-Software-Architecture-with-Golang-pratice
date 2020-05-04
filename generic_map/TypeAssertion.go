package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	result := ReturnData()

	// int(result) wouldn't work at this stage:

	// Type assert, by placing brackets and type after variable name.
	// Note that we need to assign to a new variable.
	//myInt := result.(int)

	// Now we can work with this, should print '10'
	myInt, ok := result.(int)
	if !ok {
		log.Printf("got data of type %T but wanted int", result)
		os.Exit(1)
	}

	fmt.Println(myInt)
}

func ReturnData() interface{} {
	return "b"
}
