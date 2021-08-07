package db

import (
	"strconv"
	"strings"
	"time"
)

type InputDatetime struct {
	year   string
	month  string
	date   string
	hour   string
	minute string
	second string
}

func FormatDatetime(input *InputDatetime) string {
	// Handle more faster
	day := strings.Join([]string{input.year, input.month, input.date}, "-")
	time := strings.Join([]string{input.hour, input.minute, input.second}, ":")

	return strings.Join([]string{day, time}, " ")
}

func GetNowFormattedStr() string {
	t := time.Now()

	input := InputDatetime{
		year:   strconv.Itoa(t.Year()),
		month:  t.Month().String(),
		date:   strconv.Itoa(t.Day()),
		hour:   strconv.Itoa(t.Hour()),
		minute: strconv.Itoa(t.Minute()),
		second: strconv.Itoa(t.Second()),
	}

	return FormatDatetime(&input)
}
