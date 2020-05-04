// This file demonstrates struct embeddeding in Go
package main

import (
	"fmt"
)

type iAbc interface {
	helloword(string) string
}

func (p Bird) helloword(name string) string {
	return "Bird: " + name
}
func (p Pigeon) helloword(name string) string {
	return "Pigeon: " + name
}

// Bird is a sample 'super class' to demonstrate 'inheritance' via embedding
type Bird struct {
	featherLength  int
	classification string
}

// Pigeon is the derived struct
type Pigeon struct {
	Bird
	featherLength float64
	Name          string
}

func main() {
	p := Pigeon{Name: "Tweety"}
	p.featherLength = 3.14
	p.Bird.featherLength = 123

	fmt.Printf("%+v", p.Bird.helloword("12"))
}
