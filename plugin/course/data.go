package course

import (
	"time"
)

func getWeek(time2 time.Time) (int, int) {
	oneWeek, err := time.Parse("2006-01-02 15:04:05", "2022-02-28 00:00:00")
	if err != nil {
		return 0, 0
	}
	day := time2.Sub(oneWeek).Hours() / 24 / 7
	args := time2.Sub(oneWeek).Hours() / 24
	return int(day) + 1, int(args)%7 + 1
}

func getDay() {
	println(time.Now().Weekday().String())
}
