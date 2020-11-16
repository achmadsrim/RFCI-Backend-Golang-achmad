package main

import "fmt"

func insertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		j := i
		for j > 0 && arr[j-1] > tmp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = tmp
	}
}
func main() {
	arr := []int{4, 9, 7, 5, 8, 9, 3}
	fmt.Print("Isi Element Array: ", arr, "\n", "Hasil Pengurutan: [ ")
	insertionSort(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Print("]\n")
}
