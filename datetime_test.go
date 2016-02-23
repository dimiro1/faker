package faker

import (
	"testing"
	"time"
)

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

func TestFakerDateTimeBetween(t *testing.T) {
	f := NewDefault()
	begin := time.Now()
	end := time.Now().Add(time.Hour * 6)
	c := f.DateTimeBetween(begin, end)

	if c.Before(begin) || c.After(end) {
		t.Errorf("DateTimeBetween must be beetween the specified dates: %s", c)
	}
}

func TestFakerDateTimeAfter(t *testing.T) {
	f := NewDefault()
	now := time.Now()
	c := f.DateTimeAfter(now)

	if !c.After(now) {
		t.Error("DateTimeAfter must be after the specified date")
	}
}

func TestFakerDateTimeBefore(t *testing.T) {
	f := NewDefault()
	now := time.Now()
	c := f.DateTimeBefore(now)

	if !c.Before(now) {
		t.Error("DateTimeBefore must be before the specified date")
	}
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
