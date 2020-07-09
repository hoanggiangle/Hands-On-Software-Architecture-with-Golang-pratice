package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("i: ", i, " - ", s)
	}
}
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
func main() {
	fmt.Println("main")
	c := make(chan int, 2)
	c <- 2
	c <- 3
	c3 := func() {
		c <- 41
		c <- 5
	}
	go c3()

	fmt.Println(<- c, <- c, <- c)
	time.Sleep(1 * time.Second)
}
