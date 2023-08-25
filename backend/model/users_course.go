package model

import "time"

type UsersCourse struct {
	Course          int       `dynamodbav:"course"`
	LastTimeWatched time.Time `dynamodbav:"last_time_watched"`
}
