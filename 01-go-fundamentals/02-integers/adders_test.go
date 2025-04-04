package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {
	t.Run("with 2 and 2", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
	t.Run("with 3 and 5", func(t *testing.T) {
		sum := Add(3, 5)
		expected := 8

		if sum != expected {
			t.Errorf("expected '%d', but got '%d'", expected, sum)
		}
	})
}
