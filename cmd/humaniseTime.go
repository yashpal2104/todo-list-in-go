package cmd

import (
	"strconv"
	"time"
)

func HumanizeTimeSince(t time.Time) string {
	diff := time.Since(t)
	// diff := time.Duration.Abs(time.Duration(t.Second()))
	// print(diff)

	if (diff < 60 * time.Second) {
		return "few seconds ago"
	} 
	if (diff < 60 * time.Minute) {
		minute := int(diff/time.Hour)
		return strconv.Itoa(int(minute)) + " minutes ago"
	}
	if (diff < 60 * time.Hour) {
		minute := int(diff/time.Duration(time.Now().Day()))
		return  strconv.Itoa(int(minute)) + " minutes ago" 
	} else{
		return diff.String()
	}

}
