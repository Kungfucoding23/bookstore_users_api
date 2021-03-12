package date

import "time"

const (
	apiDateLayout = "02-01-2006T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC() //standard time zone
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout) //02-01-2006  2006-01-02
}
