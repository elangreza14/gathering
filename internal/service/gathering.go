package service

//go:generate mockgen -source $GOFILE -destination ../../mock/service/mock_$GOFILE -package $GOPACKAGE

import (
	"context"
	"errors"
	"fmt"

	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
)

type gatheringRepo interface {
	FindMemberByID(ctx context.Context, ID int64) (*domain.Member, error)
	FindGatheringByID(ctx context.Context, ID int64) (*domain.Gathering, error)
	FindInvitationByGatheringIDAndMemberID(ctx context.Context, gatheringID, memberID int64) (*domain.Invitation, error)

	CreateGathering(ctx context.Context, arg domain.Gathering) (*domain.Gathering, error)
	CreateInvitations(ctx context.Context, gatheringID int64, status domain.InvitationStatus, memberID ...int64) error
	CreateAttendee(ctx context.Context, arg domain.Attendee) (*domain.Attendee, error)
}

type GatheringService struct {
	gatheringRepo gatheringRepo
}

func NewGatheringService(repo gatheringRepo) *GatheringService {
	return &GatheringService{
		gatheringRepo: repo,
	}
}

func (gs *GatheringService) CreateGathering(ctx context.Context, req dto.CreateGatheringReq) (*dto.CreateGatheringRes, error) {
	if req.Type == domain.GatheringTypeINVITATION && len(req.Attendees) < 1 {
		return nil, errors.New("attendees must be more than 0 when type is invitation")
	}

	res, err := gs.gatheringRepo.CreateGathering(ctx, domain.Gathering{
		Creator:    req.Creator,
		Type:       req.Type,
		ScheduleAt: req.ScheduleAt,
		Name:       req.Name,
		Location:   req.Location,
	})
	if err != nil {
		return nil, err
	}

	if req.Type == domain.GatheringTypeINVITATION {
		if err := gs.gatheringRepo.CreateInvitations(ctx, res.ID, domain.InvitationStatusWAITING, req.Attendees...); err != nil {
			return nil, err
		}
	}

	return &dto.CreateGatheringRes{
		ID: res.ID,
	}, nil
}

func (gs *GatheringService) AttendGathering(ctx context.Context, req dto.CreateAttendeeReq) (*dto.CreateAttendeeRes, error) {
	member, err := gs.gatheringRepo.FindMemberByID(ctx, req.MemberID)
	if err != nil {
		fmt.Println("cek")
		return nil, err
	}

	gathering, err := gs.gatheringRepo.FindGatheringByID(ctx, req.GatheringID)
	if err != nil {
		fmt.Println("cek 2")
		return nil, err
	}

	if gathering.Type == domain.GatheringTypeINVITATION {
		invt, err := gs.gatheringRepo.FindInvitationByGatheringIDAndMemberID(ctx, gathering.ID, member.ID)
		if err != nil {
			fmt.Println("cek 3")
			return nil, err
		}

		if invt.MemberID != member.ID {
			fmt.Println("cek 4")
			return nil, errors.New("unauthorized")
		}
	}

	res, err := gs.gatheringRepo.CreateAttendee(ctx, domain.Attendee{
		MemberID:    member.ID,
		GatheringID: gathering.ID,
	})
	if err != nil {
		fmt.Println("cek 5")
		return nil, err
	}

	return &dto.CreateAttendeeRes{
		ID: res.ID,
	}, nil
}
