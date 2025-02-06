package slicex

type SlicexErr string

const (
	ErrElementNotFound = SlicexErr("element not found")
)

func (s SlicexErr) Error() string {
	return string(s)
}
