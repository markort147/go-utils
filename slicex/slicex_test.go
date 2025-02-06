package slicex

import "testing"

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
