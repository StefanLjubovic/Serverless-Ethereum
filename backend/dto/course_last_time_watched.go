package dto

import (
	"backend/model"
	"time"
)

type CourseLastTimeWatched struct {
	Course          model.Course `json:"course"`
	LastTimeWatched time.Time    `json:"last_time_watched"`
}
