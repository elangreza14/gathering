package service

//go:generate mockgen -source $GOFILE -destination ../../mock/service/mock_$GOFILE -package $GOPACKAGE

import (
	"context"
	"errors"

	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
)

type memberRepo interface {
	FindMemberByID(ctx context.Context, ID int64) (*domain.Member, error)
	FindInvitationByID(ctx context.Context, ID int64) (*domain.Invitation, error)

	CreateMember(ctx context.Context, arg domain.Member) (*domain.Member, error)

	UpdateInvitation(ctx context.Context, arg domain.Invitation) error
}

type MemberService struct {
	memberRepo memberRepo
}

func NewMemberService(repo memberRepo) *MemberService {
	return &MemberService{
		memberRepo: repo,
	}
}

func (is *MemberService) CreateMember(ctx context.Context, req dto.CreateMemberReq) (*dto.CreateMemberRes, error) {
	res, err := is.memberRepo.CreateMember(ctx, domain.Member{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateMemberRes{
		ID: res.ID,
	}, nil
}

func (is *MemberService) RespondInvitation(ctx context.Context, req dto.RespondInvitationReq) error {
	member, err := is.memberRepo.FindMemberByID(ctx, req.MemberID)
	if err != nil {
		return err
	}

	invitation, err := is.memberRepo.FindInvitationByID(ctx, req.InvitationID)
	if err != nil {
		return err
	}

	if member.ID != invitation.MemberID {
		return errors.New("unauthorized")
	}

	if req.Attend {
		invitation.Status = "ATTEND"
	} else {
		invitation.Status = "ABSENT"
	}

	return is.memberRepo.UpdateInvitation(ctx, *invitation)
}
