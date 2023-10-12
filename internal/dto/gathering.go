package dto

import (
	"time"

	"github.com/elangreza14/gathering/internal/domain"
)

type CreateGatheringReq struct {
	Creator    string               `json:"creator" binding:"required"`
	Type       domain.GatheringType `json:"type" binding:"oneof=FREE INVITATION"`
	ScheduleAt time.Time            `json:"schedule_at" binding:"required"`
	Name       string               `json:"name" binding:"required"`
	Location   string               `json:"location" binding:"required"`
	Attendees  []int64              `json:"attendees"`
}

type CreateGatheringRes struct {
	ID int64
}
