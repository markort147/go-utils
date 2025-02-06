package slicex

type SlicexErr string

const (
	ErrElementNotFound = SlicexErr("element not found")
)

func (s SlicexErr) Error() string {
	return string(s)
}

func Find[T any](numbers []T, predicate func(T) bool) (T, error) {
	for _, n := range numbers {
		if predicate(n) {
			return n, nil
		}
	}

	var zero T
	return zero, ErrElementNotFound
}
