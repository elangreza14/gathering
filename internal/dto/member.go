package dto

type CreateMemberReq struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type CreateMemberRes struct {
	ID int64 `json:"id"`
}

type RespondInvitationReq struct {
	MemberID     int64 `json:"member_id" binding:"required"`
	InvitationID int64 `json:"invitation_id" binding:"required"`
	Attend       bool  `json:"attend" binding:"required"`
}
