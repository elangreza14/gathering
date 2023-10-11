package dto

type CreateAttendeeReq struct {
	MemberID    int64
	GatheringID int64
}

type CreateAttendeeRes struct {
	ID int64
}
