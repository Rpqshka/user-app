package handler

import (
	"math"
	"time"
	userApp "user-app"
)

func getRecordingDate() int64 {
	currentTime := time.Now()
	recordingDate := currentTime.Unix()
	return recordingDate
}

func checkInput(input userApp.SearchInput) userApp.SearchInput {
	if input.EndDate == 0 {
		input.EndDate = math.MaxInt64
	}
	if input.MaxAge == 0 {
		input.MaxAge = 101
	}

	return input
}
