package slice

import "sync"

// Реализация функции Filter с использованием мьютекса
func Filter[T any](s []T, predicate func(T) (ok bool)) (result []T) {
	var wg sync.WaitGroup
	result = make([]T, 0, len(s)) // Предварительное выделение памяти может ускорить работу
	mux := &sync.Mutex{}          // Мьютекс для синхронизации доступа к output

	// Для каждого элемента слайса запускаем горутину
	for _, element := range s {
		wg.Add(1)
		go func(e T) {
			// Освобождаем горутину после завершения работы
			defer wg.Done()
			// Если predicate вернет true, то добавляем элемент в result
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
