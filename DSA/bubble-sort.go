package main

import (
	"fmt"
	"math/rand"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}

// with boolean flag to stop sorting when needed
func bubbleSort(list []int) []int {
	defer duration(track("bubbleSort"))
	var i, j int
	len := len(list)
	for i = 0; i < len; i++ {
		for j = 0; j < len-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

// the less efficient version
func bubbleSortImp(list []int) []int {
	defer duration(track("bubbleSortImp"))
	var i, j int
	len := len(list)
	for i = 0; i < len; i++ {
		swap := false
		for j = 0; j < len-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	return list
}

func main2() {
	list := []int{}
	for i := 0; i < 10; i++ {
		list = append(list, rand.Intn(10000))

	}
	//fmt.Println(bubbleSort(list))
	bubbleSortImp(list)
}
