package slicex

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("find a even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := 2

		got, err := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("find a short word", func(t *testing.T) {
		words := []string{"hello", "go", "slicex"}
		want := "go"

		got, err := Find(words, func(x string) bool {
			return len(x) < 3
		})

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("element not found", func(t *testing.T) {
		words := []string{"hello", "go", "slicex"}

		_, err := Find(words, func(x string) bool {
			return len(x) > 10
		})

		if err != ErrElementNotFound {
			t.Errorf("expected error %q but got %q", ErrElementNotFound, err)
		}
	})
}

func TestFindAll(t *testing.T) {
	t.Run("find even numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := []int{2, 4}

		got := FindAll(numbers, func(x int) bool {
			return x%2 == 0
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("find short words", func(t *testing.T) {
		words := []string{"hello", "go", "slicex"}
		want := []string{"go"}

		got := FindAll(words, func(x string) bool {
			return len(x) < 3
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})

	t.Run("elements not found", func(t *testing.T) {
		words := []string{"hello", "go", "slicex"}
		var want []string

		got := FindAll(words, func(x string) bool {
			return len(x) > 10
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("mapping items", func(t *testing.T) {
		words := []string{"hello", "go", "slicex"}
		want := []int{5, 2, 6}

		got := Map(words, func(w string) int {
			return len(w)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}
