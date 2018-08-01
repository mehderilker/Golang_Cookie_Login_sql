package convert

import (
	"time"
	"strconv"
)

func ToDateString(t time.Time) string {
	if (t == time.Time{}) {
		return "--"
	}
	return t.UTC().Format("02.01.2006")
}

func ToString(v int) string{
	var new= strconv.Itoa(v)
	return new
}

func ToInt(v string) int{
	var new,_ = strconv.Atoi(v)
	return new
}

func ToFloat64(v string ) float64{
	var new,_ = strconv.ParseFloat(v,64)
	return new
}