package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var (
		max_count = -1
		max_item  string
		list      = []int{}
	)
	for i := 0; i < 1000000; i++ {
		list = append(list, rand.Intn(10))
	}
	max_item, max_count = mostFrequent(list)
	fmt.Println(list)
	fmt.Println(max_item, max_count)
}

func mostFrequent(list []int) (string, int) {
	var (
		maxCount = -1
		maxItem  string
		count    = make(map[string]int)
	)
	for _, item := range list {
		itemStr := fmt.Sprintf("%v", item)
		if _, ok := count[itemStr]; ok {
			count[itemStr]++

		} else {
			count[itemStr] = 1
		}
		if count[itemStr] > maxCount {
			maxCount = count[itemStr]
			maxItem = itemStr
		}
	}
	return maxItem, maxCount

}
