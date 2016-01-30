package faker

import "testing"

func TestFakerAMPM(t *testing.T) {
	f := NewDefault()
	l := f.AMPM()
	assertElementInSlice(t, l, ampm)
}

func TestFakerCentury(t *testing.T) {
	f := NewDefault()
	l := f.Century()
	assertElementInSlice(t, l, centuries)
}

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

func TestFakerDayOfWeek(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.DayOfWeek()
		if n < 1 || n > 7 {
			t.Errorf("DayOfWeek must return numbers from 1 inclusive to 7 inclusive: %d", n)
			return
		}
	}
}

func TestFakerDayOfMonth(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := f.DayOfMonth()
		if n < 1 || n > 31 {
			t.Errorf("DayOfMonth must return numbers from 1 inclusive to 31 inclusive: %d", n)
			return
		}
	}
}

func TestFakerTimeZone(t *testing.T) {
	f := NewDefault()
	l := f.TimeZone()
	assertElementInSlice(t, l, timezones)
}
