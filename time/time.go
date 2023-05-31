package time

import (
	string2 "github.com/ntminh611/mutils/string"
	"time"
)

var local, _ = time.LoadLocation("Asia/Ho_Chi_Minh")

func UnixMilli2String(t int64) string {
	tTime := time.UnixMilli(t)
	return tTime.In(local).Format(time.RFC3339)
}

func UnixMilliString2String(t string) string {
	tTime := time.UnixMilli(string2.String2Int64(t))
	return tTime.In(local).Format(time.RFC3339)
}

func UnixMilli2StringInLocAndFormat(t int64, loc *time.Location, format string) string {
	tTime := time.UnixMilli(t)
	return tTime.In(loc).Format(format)
}
