package domain

type (
	// InvitationStatus is ...
	InvitationStatus string

	// Invitation is ...
	Invitation struct {
		ID          int64
		MemberID    int64
		GatheringID int64
		Status      InvitationStatus
	}
)

const (
	// InvitationStatusWAITING is ...
	InvitationStatusWAITING InvitationStatus = "WAITING"
	// InvitationStatusATTEND is ...
	InvitationStatusATTEND InvitationStatus = "ATTEND"
	// InvitationStatusABSENT is ...
	InvitationStatusABSENT InvitationStatus = "ABSENT"
)
