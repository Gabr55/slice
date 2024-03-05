package main

import (
	"fmt"
	"sort"
	"sync"
)

// Реализация функции Filter с использованием мьютекса
func Filter[T any](s []T, predicate func(T) (ok bool)) (result []T) {
	var wg sync.WaitGroup
	result = make([]T, 0, len(s)) // Предварительное выделение памяти может ускорить работу
	mux := &sync.Mutex{}          // Мьютекс для синхронизации доступа к output

	for _, element := range s {
		wg.Add(1)
		go func(e T) {
			defer wg.Done()
			if predicate(e) {
				mux.Lock() // Блокировка мьютекса для записи
				result = append(result, e)
				mux.Unlock() // Разблокировка мьютекса
			}
		}(element)
	}
	wg.Wait() // Ожидание завершения всех горутин
	return result
}

func main() {
	// Пример использования
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(n int) bool {
		return n%2 == 0
	}

	filtered := Filter(numbers, isEven)
	sort.Slice(filtered, func(i, j int) bool {
		return filtered[i] < filtered[j]
	})
	fmt.Println(filtered) // Вывод: 10 2 4 6 8
}
