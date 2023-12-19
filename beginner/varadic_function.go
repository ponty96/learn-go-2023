package main

import "fmt"


func calcSum(nums ...int) int {
	totalSum := 0
	for _index, num := range nums {
		fmt.Println("index: ", _index, " num: ", num)
		totalSum += num
	}

	return totalSum
}

func main() {
	firstInt := 0
	secondInt := 1
	thirdInt := 2

	sum := calcSum(firstInt, secondInt, thirdInt)
	fmt.Println("sum: ", sum)
}
