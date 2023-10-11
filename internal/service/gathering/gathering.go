package gathering

import (
	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
)

type repo interface {
	FindMemberByID(ID int64) (*domain.Member, error)
	FindGatheringByID(ID int64) (*domain.Gathering, error)
	FindInvitationByGatheringIDAndMemberID(gatheringID, memberID int64) (*domain.Gathering, error)

	CreateGathering(domain.Gathering) (*domain.Gathering, error)
	CreateInvitation(gatheringID int64, status string, memberID ...int64) error
	CreateAttendee(domain.Attendee) (*domain.Attendee, error)
}

type GatheringService struct {
	gatheringRepo repo
}

func NewGatheringService(GatheringRepo repo) *GatheringService {
	return &GatheringService{
		gatheringRepo: GatheringRepo,
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
		if err := gs.gatheringRepo.CreateInvitation(res.ID, "WAITING", req.WithAttendees...); err != nil {
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
		if _, err := gs.gatheringRepo.FindInvitationByGatheringIDAndMemberID(gathering.ID, member.ID); err != nil {
			return nil, err
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
