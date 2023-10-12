package domain

type (
	InvitationStatus string
	Invitation       struct {
		ID          int64
		MemberID    int64
		GatheringID int64
		Status      InvitationStatus
	}
)

const (
	InvitationStatusWAITING InvitationStatus = "WAITING"
	InvitationStatusATTEND  InvitationStatus = "ATTEND"
	InvitationStatusABSENT  InvitationStatus = "ABSENT"
)
