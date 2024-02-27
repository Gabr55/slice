package slice

func Filter[T any](s []T, predicate func(T) (ok bool)) (result []T) {
	output := []T{}
	for _, element := range s {
		if predicate(element) {
			output = append(output, element)
		}
	}
	return output
}
