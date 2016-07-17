package helper

import "time"

/**
	获取下当天的开始时间的时间戳
 */
func GetDayBeginTime() int64 {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	td := tm.Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	tm2,_ := time.ParseInLocation("2006-01-02", td, loc)
	return tm2.Unix()
}