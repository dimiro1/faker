package faker

import "time"

func (f Faker) DateTimeBetween(d1, d2 time.Time) time.Time {
	return time.Now()
}

func (f Faker) DateTimeBefore(d time.Time) time.Time {
	return time.Now()
}

func (f Faker) DateTimeAfter(d time.Time) time.Time {
	return time.Now()
}

func (f Faker) AMPM() string {
	return "AM"
}

func (f Faker) ISO8601() string {
	return "ISO8601"
}

// Year returns a year between 1500 and 2100
func (f Faker) Year() int {
	return randomIntBetween(1500, 2100)
}

func (f Faker) Month() int {
	return 3
}

func (f Faker) DayOfWeek() int {
	return 1
}

func (f Faker) DayOfMonth() int {
	return 20
}

func (f Faker) Century() string {
	return "XX"
}

func (f Faker) TimeZone() string {
	return "America/Sao_Paulo"
}
