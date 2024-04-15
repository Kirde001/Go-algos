package main

import (
	"fmt"
	"math"
)

func huntingSearch(list []int, key int) (count int) {
	var i,temp float64
	high := len(list) - 1
	low := 0 // нижняя граница
	i = 0 // для степени в двойке, сложность алгоритма логарифмическая
	count = 0 // кол-во непосредственных вхождений
	temp = 0 // временная переменная
	for high != low {
		if list[high] == key {
			count += 1 
			list = append(list[:high], list[high+1:]...)
			high = len(list) - 1 // границы переопределяем те, с которыми работаем
			fmt.Println("I am here")
			continue
		}
		for key < list[high] && low != high {
			high -= int(math.Pow(2,i))
			i += 1 
			temp = i - 1
			if high-int(math.Pow(2,i)) <= low { 
				i = 0
			}
			if list[high] == key {
				count += 1
				list = append(list[:high], list[high+1:]...)
				high = len(list) - 1
			}
			if high == low {
				break
			}
		}
		if !(key < list[high] && low != high) {
			i = 0
			low = high
			if high + int(math.Pow(2,temp)) > len(list) - 1 {
				break 
			} else {
				high += int(math.Pow(2,temp))
			}
		}
		for key > list[low] && low != high {
			low += int(math.Pow(2,i))
			i += 1
			temp = i - 1
			if high <= low + int(math.Pow(2,i)) {
				i = 0
			}
			if key == list[low] {
				count += 1
				list = append(list[:low], list[low+1:]...)
			}
			if high == low {
				break
			}
		}
		if !(key > list[low] && low != high) {
			i = 0
			high = low
			if low - int(math.Pow(2,temp)) < 0 {
				break
			}else {
				low -= int(math.Pow(2,temp))
			}
		}
	}
	return
}
func main1() {
	var key int
	fmt.Println("")
	list := []int{}
	for i:=0; i<10; i++ {
		list = append(list, i)
    
	}
	for j := range list {
		fmt.Println(list[j])
		key = list[j]
		fmt.Println(huntingSearch(list,key))
	}
	//fmt.Println(huntingSearch(list,key))
}