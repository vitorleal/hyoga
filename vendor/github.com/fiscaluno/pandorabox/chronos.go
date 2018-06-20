package pandorabox

import "time"

// DateAddDays return a date Unix format
func DateAddDays(daysToAdd float64) int64 {
	myDate := time.Now()
	expDate := myDate.Add(time.Hour * 24 * time.Duration(daysToAdd)).Unix()
	return expDate
}
