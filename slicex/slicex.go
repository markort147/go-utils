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

func FindAll[T any](items []T, predicate func(x T) bool) (filtered []T) {
	for _, item := range items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return
}

func Map[T, R any](items []T, mapping func(T) R) (mapped []R) {
	for _, item := range items {
		mapped = append(mapped, mapping(item))
	}
	return
}
