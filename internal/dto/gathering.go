package dto

import "time"

type CreateGatheringReq struct {
	Creator       string
	Type          string
	ScheduleAt    time.Time
	Name          string
	Location      string
	WithAttendees []int64
}

type CreateGatheringRes struct {
	ID int64
}
