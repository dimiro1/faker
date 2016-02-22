package faker

import "time"

func (f Faker) DateTime() time.Time {
	return time.Now()
}

func (f Faker) DateTimeBetween(d1, d2 time.Time) time.Time {
	return time.Now()
}

func (f Faker) DateTimeBefore(d time.Time) time.Time {
	return time.Now()
}

// DateTimeAfter returns a date after the specified date
func (f Faker) DateTimeAfter(d time.Time) time.Time {
	days := time.Hour * 24 * time.Duration(randomIntBetween(1, 1000))

	return d.Add(days)
}

// AMPM returns a string AM or PM
func (f Faker) AMPM() string {
	return randomElement(ampm)
}

// RFC3339 returns a random date formatted as RFC3339
func (f Faker) RFC3339() string {
	return f.DateTime().Format(time.RFC3339)
}

// Year returns a year between 1500 and 2100
func (f Faker) Year() int {
	return randomIntBetween(1500, 2100)
}

// Month returns a month between 1 and 12
func (f Faker) Month() int {
	return randomIntBetween(1, 12)
}

// DayOfWeek returns a day of week between 1 and 7 where 1 is sunday
func (f Faker) DayOfWeek() int {
	return randomIntBetween(1, 7)
}

// DayOfMonth returns a day of the month between 1 and 31
func (f Faker) DayOfMonth() int {
	return randomIntBetween(1, 31)
}

// Century returns a century from I to XXI
func (f Faker) Century() string {
	return randomElement(centuries)
}

// TimeZone returns TZ name
func (f Faker) TimeZone() string {
	return randomElement(timezones)
}
