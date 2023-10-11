package domain

import "time"

type Gathering struct {
	ID         int64
	Creator    string
	Type       string
	ScheduleAt time.Time
	Name       string
	Location   string
}
