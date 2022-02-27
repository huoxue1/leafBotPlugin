package course

import (
	"fmt"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	getDay()
}

func TestDraw(t *testing.T) {

}

func TestGet(t *testing.T) {
	parseXlsx("course.xlsx")
}

func TestTime(t *testing.T) {
	oneWeek, _ := time.Parse("2006-01-02 15:04:05", "2021-10-24 23:59:00")
	week, day := getWeek(oneWeek)
	print(week, day)
}

func TestC(t *testing.T) {
	xlsx, err := parseXlsx("./course.xlsx")
	if err != nil {
		return
	}

	for i, courses := range xlsx {
		fmt.Println(i, "   ", courses)
	}
}

func TestName(t *testing.T) {
	courses := parse(`[173079]JavaEE程序设计 
曾勇 [1-2, 4, 6-17周][4-5节] 实验楼计算机通用实验室-407
`)
	fmt.Println(courses)
}
