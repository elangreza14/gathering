package dto

import "time"

type CreateGatheringReq struct {
	Creator    string    `json:"creator" binding:"required"`
	Type       string    `json:"type" binding:"required"`
	ScheduleAt time.Time `json:"schedule_at" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Location   string    `json:"location" binding:"required"`
	Attendees  []int64   `json:"attendees"`
}

type CreateGatheringRes struct {
	ID int64
}
