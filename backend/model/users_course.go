package model

import "time"

type UsersCourse struct {
	Course          Course    `dynamodbav:"course"`
	LastTimeWatched time.Time `dynamodbav:"last_time_watched"`
}
