package main

import "fmt"

func Sort(arr []int) {
	if len(arr) == 0 {
		return
	}
	key := arr[0]
	left := 0
	right := len(arr) - 1
	for left < right {
		for left < right && arr[right] > key {
			right--
		}
		if left < right {
			arr[left] = arr[right]
			left++
		}
		for left < right && arr[left] < key {
			left++
		}
		if left < right {
			arr[right] = arr[left]
			right--
		}
	}
	arr[left] = key
	Sort(arr[:left])
	Sort(arr[left+1:])
}

func main() {
	arr := []int{2, 3, 5, 0, -2, 23, -4}
	Sort(arr)
	fmt.Println(arr)
}
