package dto

type CreateMemberReq struct {
	FirstName string
	LastName  string
	Email     string
}

type CreateMemberRes struct {
	ID int64
}

type RespondInvitationReq struct {
	MemberID     int64
	InvitationID int64
	Attend       bool
}
