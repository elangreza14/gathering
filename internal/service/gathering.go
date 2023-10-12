package service

import (
	"errors"

	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
)

type gatheringRepo interface {
	FindMemberByID(ID int64) (*domain.Member, error)
	FindGatheringByID(ID int64) (*domain.Gathering, error)
	FindInvitationByGatheringIDAndMemberID(gatheringID, memberID int64) (*domain.Invitation, error)

	CreateGathering(domain.Gathering) (*domain.Gathering, error)
	CreateInvitations(gatheringID int64, status string, memberID ...int64) error
	CreateAttendee(domain.Attendee) (*domain.Attendee, error)
}

type GatheringService struct {
	gatheringRepo gatheringRepo
}

func NewGatheringService(repo gatheringRepo) *GatheringService {
	return &GatheringService{
		gatheringRepo: repo,
	}
}

func (gs *GatheringService) CreateGathering(req dto.CreateGatheringReq) (*dto.CreateGatheringRes, error) {
	res, err := gs.gatheringRepo.CreateGathering(domain.Gathering{
		Creator:    req.Creator,
		Type:       req.Type,
		ScheduleAt: req.ScheduleAt,
		Name:       req.Name,
		Location:   req.Location,
	})
	if err != nil {
		return nil, err
	}

	if req.Type == "INVITATION" {
		if err := gs.gatheringRepo.CreateInvitations(res.ID, "WAITING", req.WithAttendees...); err != nil {
			return nil, err
		}
	}

	return &dto.CreateGatheringRes{
		ID: res.ID,
	}, nil
}

func (gs *GatheringService) AttendGathering(req dto.CreateAttendeeReq) (*dto.CreateAttendeeRes, error) {
	member, err := gs.gatheringRepo.FindMemberByID(req.MemberID)
	if err != nil {
		return nil, err
	}

	gathering, err := gs.gatheringRepo.FindGatheringByID(req.GatheringID)
	if err != nil {
		return nil, err
	}

	if gathering.Type == "INVITATION" {
		invt, err := gs.gatheringRepo.FindInvitationByGatheringIDAndMemberID(gathering.ID, member.ID)
		if err != nil {
			return nil, err
		}

		if invt.MemberID != member.ID {
			return nil, errors.New("unauthorized")
		}
	}

	res, err := gs.gatheringRepo.CreateAttendee(domain.Attendee{
		MemberID:    member.ID,
		GatheringID: gathering.ID,
	})
	if err != nil {
		return nil, err
	}

	return &dto.CreateAttendeeRes{
		ID: res.ID,
	}, nil
}
