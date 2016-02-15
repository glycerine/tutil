package tutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

var WestCoastUSLocation *time.Location
var EastCoastUSLocation *time.Location
var LondonLocation *time.Location
var UTCLocation = time.UTC

func init() {
	var err error
	WestCoastUSLocation, err = time.LoadLocation("America/Los_Angeles")
	panicOn(err)
	EastCoastUSLocation, err = time.LoadLocation("America/New_York")
	panicOn(err)
	LondonLocation, err = time.LoadLocation("Europe/London")
	panicOn(err)
}

func panicOn(err error) {
	if err != nil {
		panic(err)
	}
}

func (d Date) Unix() int64 {
	return d.ToGoTime().Unix()
}

func (d Date) ToGoTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

func (d Date) String() string {
	return fmt.Sprintf("%d/%d/%d", d.Year, d.Month, d.Day)
}

// return true if a < b
func Before(a Date, b Date) bool {
	if a.Year < b.Year {
		return true
	} else if a.Year > b.Year {
		return false
	}

	if a.Month < b.Month {
		return true
	} else if a.Month > b.Month {
		return false
	}

	if a.Day < b.Day {
		return true
	} else if a.Day > b.Day {
		return false
	}

	return false
}

func NewDate(s string) Date {
	slc := strings.Split(s, "/")
	if len(slc) != 3 {
		panic(fmt.Sprintf("NewDate error on input: '%s', did not find two '/', expected format: 2013/11/29", s))
	}
	y, err := strconv.Atoi(slc[0])
	if err != nil {
		panic(err)
	}
	m, err := strconv.Atoi(slc[1])
	if err != nil {
		panic(err)
	}
	d, err := strconv.Atoi(slc[2])
	if err != nil {
		panic(err)
	}

	if y < 1970 || y > 2030 {
		panic(fmt.Sprintf("year out of bounds: %v", y))
	}
	if m < 1 || m > 12 {
		panic(fmt.Sprintf("month out of bounds: %v", m))
	}
	if d < 1 || d > 31 {
		panic(fmt.Sprintf("day out of bounds: %v", d))
	}
	return Date{Year: y, Month: m, Day: d}
}

func DatesEqual(a Date, b Date) bool {
	if a.Year == b.Year {
		if a.Month == b.Month {
			if a.Day == b.Day {
				return true
			}
		}
	}
	return false
}
