package tutil

import (
	"math"
	"time"
)

// nanosec since unix epoch. Commonly: as returned by time.UnixNano()
type Ntm int64

// msec since midnight
type Mm int64

// human readable
type Htm int64

func Millisec(a int) Ntm { return Ntm(a * 1e6) }

func DefaultDay() Ntm {
	// HmToTm() doesn't have a date by default, so default to 2013-11-15 04:05:47.655169476 -0500 EST
	return Ntm(1384506347655169476)
}

// HmToTm converts from human readable time, such as 930 (for 9:30am)
// to tm_t, which is Milliseconds since midnight.
// For example human readable, 24-hour clock time:
//     1559 for 3:59pm -> 15*60*60*1000 + 59*60*1000 in tm_t
func HmToTm(hm Htm) Ntm {
	// convert to msec since midnight
	hrs := hm / 100
	minutes := hm % 100
	msecMidnt := int64((minutes + hrs*60) * 60 * 1000)

	//return Mm(msecMidnt)
	return MsecMidntToEpocNanosec(Mm(msecMidnt), DefaultDay())
}

func HmToMm(hm Htm) Mm {
	// convert to msec since midnight
	hrs := hm / 100
	minutes := hm % 100
	msecMidnt := int64((minutes + hrs*60) * 60 * 1000)

	return Mm(msecMidnt)
}

func MmToHm(mm Mm) Htm {
	// convert from msec since midnight to human readable
	// e.g. 34200000 -> 930

	totsec := mm / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60

	return Htm(hrs*100 + minutes)
}

// merge a msec since midnight time and a nano-seconds since epoch time. Use the nanosec since epoch grabDateFrom for the day.
// modifies grabDateFrom to have the time of day given by msecMidnt
func MsecMidntToEpocGoTime(msecMidnt Mm, grabDateFrom Ntm) time.Time {

	// parse msecMidnt for hrs, minutes, totsec, msec
	msec := int64(msecMidnt) % 1000
	totsec := msecMidnt / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60
	totsec -= minutes * 60

	unano := time.Unix(0, int64(grabDateFrom))
	unano = unano.UTC()
	year, month, day := unano.Date()

	gotime := time.Date(year, month, day, int(hrs), int(minutes), int(totsec), int(msec)*1000000, LondonLocation)
	//	return Ntm(gotime.UnixNano())
	return gotime
}

func MsecMidntToEpocNanosec(msecMidnt Mm, grabDateFrom Ntm) Ntm {
	gotime := MsecMidntToEpocGoTime(msecMidnt, grabDateFrom)
	return Ntm(gotime.UnixNano())
}

// strip out the date part, leaving a time since midnight of 1 Jan 1970, in nanoseconds
func StripDayNtm(tm Ntm) Ntm {
	unano := time.Unix(0, int64(tm))
	return StripDayGoTime(unano)
}

func StripDayGoTime(unano time.Time) Ntm {

	hrs, minutes, sec := unano.Clock()
	nsec := unano.Nanosecond()

	ntm := Ntm(time.Date(1970, time.Month(1), 1, hrs, minutes, sec, nsec, time.UTC).UnixNano())

	//fmt.Printf("\n ntm = %v", ntm)
	return ntm
}

// re-parent tm to be on the date given by grabDateFrom, returning the reparented Ntm
func ReplaceDayOnlyNtm(tm Ntm, grabDateFrom Ntm) Ntm {

	parent := time.Unix(0, int64(grabDateFrom))
	parent = parent.UTC()
	year, month, day := parent.Date()

	forTm := time.Unix(0, int64(tm))
	forTm = forTm.UTC()
	hr, minutes, sec := forTm.Clock()
	nsec := forTm.Nanosecond()

	ntm := Ntm(time.Date(year, time.Month(month), day, hr, minutes, sec, nsec, LondonLocation).UnixNano())

	//	fmt.Printf("\n ntm = %v", ntm)
	return ntm

}

func NtmToMsecMidnt(tm Ntm) Mm {
	unano := time.Unix(0, int64(tm))
	unano = unano.UTC()
	hr, minutes, sec := unano.Clock()
	nsec := unano.Nanosecond()

	msec := int64(nsec) / int64(1000000)
	totsec := int64((hr * 60 * 60) + (minutes * 60) + sec)
	return Mm(totsec*1000 + msec)
}

func MmToNtmLondon(msecMidnt Mm, grabDateFrom Ntm) Ntm {
	gotm := MsecMidntToEpocGoTime(msecMidnt, grabDateFrom)
	return StripDayGoTime(gotm)
}

func MmToNtmSimple(x Mm) Ntm {
	return Ntm(x * 1e6)
}

func StartNtm(hm Htm, datestring string) Ntm {
	msecMidnt := HmToMm(hm)
	d := NewDate(datestring)

	// parse msecMidnt for hrs, minutes, totsec, msec
	msec := int64(msecMidnt) % 1000
	totsec := msecMidnt / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60
	totsec -= minutes * 60

	gotime := time.Date(d.Year, time.Month(d.Month), d.Day, int(hrs), int(minutes), int(totsec), int(msec)*1000000, LondonLocation)

	return Ntm(gotime.UnixNano())
}

func MmDate(msecMidnt Mm, d Date) Ntm {

	// parse msecMidnt for hrs, minutes, totsec, msec
	msec := int64(msecMidnt) % 1000
	totsec := msecMidnt / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60
	totsec -= minutes * 60

	gotime := time.Date(d.Year, time.Month(d.Month), d.Day, int(hrs), int(minutes), int(totsec), int(msec)*1000000, LondonLocation)

	return Ntm(gotime.UnixNano())
}

func StartMm(msecMidnt Mm, datestring string) Ntm {
	d := NewDate(datestring)
	return MmDate(msecMidnt, d)
}

/*
func TmToHm(msecMidnt Tm) Htm {
	//msec := msecMidnt % 1000
	totsec := msecMidnt / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60
	//totsec -= minutes * 60

	return Htm(hrs*100 + minutes)
}
*/

func TmToFloat64(tm Ntm) float64 {
	msecMidnt := NtmToMsecMidnt(tm)

	msec := msecMidnt % 1000
	totsec := msecMidnt / 1000

	hrs := totsec / (60 * 60)
	totsec -= hrs * 60 * 60

	minutes := totsec / 60
	totsec -= minutes * 60

	f := float64(hrs*10000 + minutes*100 + totsec)
	f += float64(msec) / 1000.0
	return f
}

func NtmToDate(tm Ntm) Date {
	parent := time.Unix(0, int64(tm))
	parent = parent.UTC()
	year, month, day := parent.Date()

	return Date{Year: year, Month: int(month), Day: day}
}

func NtmToGoTime(tm Ntm) time.Time {
	parent := time.Unix(0, int64(tm))
	parent = parent.UTC()
	return parent
}

func GoTimeToNtm(gotime time.Time) Ntm {
	return Ntm(gotime.UnixNano())
}

const TmTZFmt = "2006-01-02 15:04:05.000 -0700 MST"

const TmFmt = "2006-01-02 15:04:05.000"

func (tm Ntm) String() string {
	return NtmToGoTime(tm).Format(TmFmt)
}

func (tm Ntm) StringWithTZ() string {
	return NtmToGoTime(tm).Format(TmTZFmt)
}

func MaxT(a Ntm, b Ntm) Ntm {
	if a > b {
		return a
	}
	return b
}

const MaxNtm = math.MaxInt64

func DeepCopyNtmSlice(a []Ntm) []Ntm {
	r := make([]Ntm, len(a))
	copy(r, a)
	return r
}
