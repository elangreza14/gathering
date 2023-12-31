package dto

import (
	"time"

	"github.com/elangreza14/gathering/internal/domain"
)

// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/

// CreateGatheringReq is ...
type CreateGatheringReq struct {
	Creator    string               `json:"creator" binding:"required"`
	Type       domain.GatheringType `json:"type" binding:"oneof=FREE INVITATION"`
	ScheduleAt time.Time            `json:"schedule_at" binding:"required"`
	Name       string               `json:"name" binding:"required"`
	Location   string               `json:"location" binding:"required"`
	Attendees  []int64              `json:"attendees"`
}

// CreateGatheringRes is ...
type CreateGatheringRes struct {
	ID int64 `json:"id"`
}
