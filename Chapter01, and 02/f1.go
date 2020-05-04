package main

import "fmt"

type Stringer12 interface {
	String() string
}

type fakeString struct {
	content string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value string) {

	switch value {
	case "1":
		fmt.Println(value)
	case "2":
		fmt.Println("--===--", value)
	}
}

func main() {
	s := &fakeString{content: "abc"}
	printString(s)
	printString("Hello, Gophers")

}
