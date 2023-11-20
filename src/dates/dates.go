package dates

import (
	"fmt"
	"time"
)

type MonthSet struct {
	Prev time.Time
	Curr time.Time
	Next time.Time
}

func GetCurrentMonthString() string {
	now := time.Now()
	month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return month.Format("2006-01")
}

func GetCurrentMonth(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), 1, 0, 0, 0, 0, d.Location())
}

func GetPreviousMonth(d time.Time) time.Time {
	m := d.Month() - 1
	y := d.Year()
	if m == 12 {
		y -= 1
	}
	prev := time.Date(y, m, 1, 0, 0, 0, 0, d.Location())
	return prev
}

func GetNextMonth(d time.Time) time.Time {
	m := d.Month() + 1
	y := d.Year()
	if m == 1 {
		y += 1
	}
	prev := time.Date(y, m, 1, 0, 0, 0, 0, d.Location())
	return prev
}

func GetMonthSet(date time.Time) MonthSet {
	return MonthSet{Prev: GetPreviousMonth(date), Curr: GetCurrentMonth(date), Next: GetNextMonth(date)}

}

func MustDateFromString(s string) time.Time {
	t, err := time.Parse("2006-01", s)
	print(t.GoString())
	if err != nil {
		fmt.Print(err)
		return time.Now()
	}
	return t
}
