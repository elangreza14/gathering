// Package dto is
package dto

// CreateAttendeeReq ...
type CreateAttendeeReq struct {
	MemberID    int64 `json:"member_id" binding:"required,gte=0"`
	GatheringID int64 `json:"gathering_id" binding:"required,gte=0"`
}

// CreateAttendeeRes ...
type CreateAttendeeRes struct {
	ID int64 `json:"id" binding:"required"`
}
