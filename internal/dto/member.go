// Package dto is ...
package dto

import "github.com/elangreza14/gathering/internal/domain"

// CreateMemberReq ...
type CreateMemberReq struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

// CreateMemberRes ...
type CreateMemberRes struct {
	ID int64 `json:"id"`
}

// RespondInvitationReq ...
type RespondInvitationReq struct {
	MemberID     int64                   `json:"member_id" binding:"required,gte=0"`
	InvitationID int64                   `json:"invitation_id" binding:"required,gte=0"`
	Attend       domain.InvitationStatus `json:"attend" binding:"required,oneof=ATTEND ABSENT"`
}
