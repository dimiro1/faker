package faker

import "time"

// Ask google
// https://www.google.com.br/webhp?ie=UTF-8#q=100+years+in+hours
const maximumHours int = 876000 // 100 years in hours

// DateTime returns a random datetime
// The date is maximum 100 years from now in the past or in the future
func (f Faker) DateTime() time.Time {
	if randomBoolean() {
		return f.DateTimeBefore(time.Now())
	} else {
		return f.DateTimeAfter(time.Now())
	}
}

// DateTimeBetween returns a date beetween d1 and d2
func (f Faker) DateTimeBetween(d1, d2 time.Time) time.Time {
	diff := d2.Sub(d1).Hours()
	hours := randomIntBetween(1, int(diff))

	return d1.Add(time.Hour * time.Duration(hours))
}

// DateTimeBefore returns a date before the specified date
// The date is maximum 100 years from now in the past
func (f Faker) DateTimeBefore(d time.Time) time.Time {
	hours := time.Hour * time.Duration(randomIntBetween(1, maximumHours))

	return d.Add(-hours)
}

// DateTimeAfter returns a date after the specified date
// The date is maximum 100 years from now in the future
func (f Faker) DateTimeAfter(d time.Time) time.Time {
	days := time.Hour * time.Duration(randomIntBetween(1, maximumHours))

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
