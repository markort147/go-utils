package slicex

func Find[T any](items []T, predicate func(T) bool) (T, error) {
	for _, item := range items {
		if predicate(item) {
			return item, nil
		}
	}

	var zero T
	return zero, ErrElementNotFound
}

func FindAll[T any](items []T, predicate func(x T) bool) []T {
	filtered := make([]T, 0)
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
