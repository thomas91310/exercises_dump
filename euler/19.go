//You are given the following information, but you may prefer to do some research for yourself.

// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package main

var (
	yearStart = 1900
	yearEnd   = 2001
)

var (
	days = map[int]string{
		1: "Monday",
		2: "Tuesday",
		3: "Wednesday",
		4: "Thursday",
		5: "Friday",
		6: "Saturday",
		7: "Sunday",
	}
)

func amazingSundays() int {
	startDay := 1
	amazingSundays := 0
	for year := yearStart; year < yearEnd; year++ {
		months := detectLeapYear(year)
		for _, numDays := range months {
			counterDays := 1
			for {
				if startDay == 6 && counterDays == 1 {
					amazingSundays++
				}
				if startDay == 7 {
					startDay = 1
				} else {
					startDay++
				}
				if counterDays == numDays {
					break
				}
				counterDays++
			}
		}
	}
	return amazingSundays
}

func detectLeapYear(year int) map[int]int {
	months := map[int]int{
		1:  31,
		2:  28,
		3:  30,
		4:  31,
		5:  30,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	if year%100 == 0 && year%400 == 0 {
		months[2] = 29
	}
	if year%4 == 0 && year%100 != 0 && year%400 != 0 {
		months[2] = 29
	}
	return months
}
