package domain

type Invitation struct {
	ID          int64
	MemberID    int64
	GatheringID int64
	Status      string
}
