package faker

import "testing"

func TestFakerYear(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.Year()
		if n < 1500 || n > 2100 {
			t.Errorf("Year must return numbers from 1500 inclusive to 2100 inclusive: %d", n)
			return
		}
	}
}

func TestFakerMonth(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.Month()
		if n < 1 || n > 12 {
			t.Errorf("Month must return numbers from 1 inclusive to 12 inclusive: %d", n)
			return
		}
	}
}
