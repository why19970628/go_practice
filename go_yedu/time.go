package main

import (
	"strconv"
	"time"
)

var (
	TIME_LAYOUT     = "2006-01-02 15:04:05"
	DATE_LAYOUT     = "2006-01-02"
	TIME_INT_LAYOUT = "20060102150405"
	DATE_INT_LAYOUT = "20060102"
	TIME_ZONE       = "Asia/Shanghai"
)

/**
 *  2006-01-02 15:04:05 => 1136185445
 */
func DateTimeToUnixTime(dateTime string) (int64, error) {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t, err := time.ParseInLocation(TIME_LAYOUT, dateTime, loc)
	return t.Unix(), err
}

/**
 * 2006-01-02 00:00:00 => 1136131200
 */
func DateToZeroUnixTime(date string) (int64, error) {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t, err := time.ParseInLocation(DATE_LAYOUT, date[0:10], loc)
	return t.Unix(), err
}

/**
 * 1136131200 => 2006-01-02
 */
func UnixTimeToDate(unixTime int64) string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Unix(unixTime, 0)
	return t.In(loc).Format(DATE_LAYOUT)
}

/**
 * 1136185445 => 2006-01-02 15:04:05
 */
func UnixTimeToDateTime(unixTime int64) string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Unix(unixTime, 0)
	return t.In(loc).Format(TIME_LAYOUT)
}

/**
 * Now unixtime
 */
func GetCurrentUnixTime() int64 {
	loc, _ := time.LoadLocation(TIME_ZONE)
	return time.Now().In(loc).Unix()
}

/**
 * today date
 */
func GetCurrentDate() string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Now().In(loc)
	return t.In(loc).Format(DATE_LAYOUT)
}

func GetCurrentDateTimeStr() string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Now().In(loc)
	return t.In(loc).Format(TIME_INT_LAYOUT)
}

func GetCurrentDateStr() string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Now().In(loc)
	return t.In(loc).Format(DATE_INT_LAYOUT)
}

/**
 * today datetime
 */
func GetCurrentDateTime() string {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t := time.Now().In(loc)
	return t.Format(TIME_LAYOUT)
}

/**
 * 2006-01-02 15:04:05 => 20060102150405
 */
func DateTimeToInt(dateTime string) (int64, error) {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t, _ := time.ParseInLocation(TIME_LAYOUT, dateTime, loc)
	dtStr := t.In(loc).Format(TIME_INT_LAYOUT)
	return strconv.ParseInt(dtStr, 10, 64)
}

/**
 * 2006-01-02 00:00:00 => 20060102000000
 */
func DateToInt(date string) (int64, error) {
	loc, _ := time.LoadLocation(TIME_ZONE)
	t, _ := time.ParseInLocation(TIME_LAYOUT, date, loc)
	dtStr := t.In(loc).Format(DATE_INT_LAYOUT)
	return strconv.ParseInt(dtStr, 10, 64)
}

func Tomorrow(unixTime int64) int64 {
	return unixTime + 86400
}

func AfterDays(unixTime int64, days int) int64 {
	return unixTime + int64(86400*days)
}

func Current() time.Time {
	loc, _ := time.LoadLocation(TIME_ZONE)
	return time.Now().In(loc)
}

func GetUnixNano() int64 {
	return time.Now().UnixNano()
}
