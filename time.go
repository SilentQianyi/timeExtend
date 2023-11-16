package timeExtend

import "time"

func GetCurZero() time.Time {
	curTime := time.Now()
	return time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, curTime.Location())
}

func GetZeroByTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
