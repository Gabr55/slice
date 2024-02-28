package main

import "fmt"

// Go - predicate based functions
// https://knowlib.app/p/functional-programming-with-go-predicate-based-functions

// Функция вернет только те элементы, где predicate == true
func Filter[T any](s []T, predicate func(T) (ok bool)) (result []T) {
	output := []T{}
	for _, element := range s {
		if predicate(element) {
			output = append(output, element)
		}
	}
	return output
}

func main() {
	// Функция проверяющая четное число
	isEven := func(n int) bool { return n%2 == 0 }
	// Слайс чисел 1-10
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Запуск фунции Filter
	evenNums := Filter(nums, isEven)
	fmt.Println(evenNums)
}
