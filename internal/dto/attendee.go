package dto

type CreateAttendeeReq struct {
	MemberID    int64 `json:"member_id" binding:"required"`
	GatheringID int64 `json:"gathering_id" binding:"required"`
}

type CreateAttendeeRes struct {
	ID int64 `json:"id" binding:"required"`
}
