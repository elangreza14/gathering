package domain

import "time"

type (
	GatheringType string
	Gathering     struct {
		ID         int64
		Creator    string
		Type       GatheringType
		ScheduleAt time.Time
		Name       string
		Location   string
	}
)

const (
	GatheringTypeFREE       GatheringType = "FREE"
	GatheringTypeINVITATION GatheringType = "INVITATION"
)
