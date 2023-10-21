package domain

import "time"

type (
	// GatheringType is ...
	GatheringType string

	// Gathering is ...
	Gathering struct {
		ID         int64
		Creator    string
		Type       GatheringType
		ScheduleAt time.Time
		Name       string
		Location   string
	}
)

const (
	// GatheringTypeFREE ...
	GatheringTypeFREE GatheringType = "FREE"
	// GatheringTypeINVITATION ...
	GatheringTypeINVITATION GatheringType = "INVITATION"
)
