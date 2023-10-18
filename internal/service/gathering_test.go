package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
	. "github.com/elangreza14/gathering/internal/service"
	gomockService "github.com/elangreza14/gathering/mock/service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestGatheringServiceSuite struct {
	suite.Suite

	MockGatheringRepo *gomockService.MockgatheringRepo

	MockCreateGatheringReq dto.CreateGatheringReq
	MockCreateAttendeeReq  dto.CreateAttendeeReq
	MockTimePassed         time.Time
	Cs                     *GatheringService
	Ctrl                   *gomock.Controller
}

func (suite *TestGatheringServiceSuite) SetupSuite() {
	suite.Ctrl = gomock.NewController(suite.T())
	suite.MockGatheringRepo = gomockService.NewMockgatheringRepo(suite.Ctrl)
	suite.MockCreateGatheringReq = dto.CreateGatheringReq{
		Type:      domain.GatheringTypeINVITATION,
		Attendees: []int64{},
	}
	suite.MockCreateAttendeeReq = dto.CreateAttendeeReq{
		MemberID:    1,
		GatheringID: 1,
	}
	suite.Cs = NewGatheringService(suite.MockGatheringRepo)

	layoutFormat := "2006-01-02 15:04:05"
	value := "2023-10-12 08:04:00"
	v, _ := time.Parse(layoutFormat, value)
	suite.MockTimePassed = v
}

func (suite *TestGatheringServiceSuite) TearDownSuite() {
	suite.Ctrl.Finish()
}

func TestGatheringServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TestGatheringServiceSuite))
}

func (suite *TestGatheringServiceSuite) TestGatheringService_CreateGathering() {
	suite.Run("error when type is invitation attendees must be more than 0", func() {
		ctx := context.Background()
		_, err := suite.Cs.CreateGathering(ctx, suite.MockCreateGatheringReq)
		suite.Error(err)
		suite.Equal(err.Error(), "attendees must be more than 0 when type is invitation")
	})

	suite.Run("error when type is invitation and create gathering is error", func() {
		ctx := context.Background()
		suite.MockCreateGatheringReq.Attendees = []int64{1}
		suite.MockGatheringRepo.EXPECT().CreateGathering(ctx, gomock.Any()).Return(nil, errors.New("err from db"))

		_, err := suite.Cs.CreateGathering(ctx, suite.MockCreateGatheringReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("error when type is invitation and no invitation data", func() {
		ctx := context.Background()
		suite.MockCreateGatheringReq.Attendees = []int64{1}
		suite.MockGatheringRepo.EXPECT().CreateGathering(ctx, gomock.Any()).Return(&domain.Gathering{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().CreateInvitations(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("err from db"))

		_, err := suite.Cs.CreateGathering(ctx, suite.MockCreateGatheringReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("create success when type is invitation", func() {
		ctx := context.Background()
		suite.MockCreateGatheringReq.Attendees = []int64{1}
		suite.MockGatheringRepo.EXPECT().CreateGathering(ctx, gomock.Any()).Return(&domain.Gathering{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().CreateInvitations(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

		res, err := suite.Cs.CreateGathering(ctx, suite.MockCreateGatheringReq)
		suite.NoError(err)
		suite.Equal(res.ID, int64(1))
	})

	suite.Run("create success when type is free", func() {
		ctx := context.Background()
		suite.MockCreateGatheringReq.Type = domain.GatheringTypeFREE
		suite.MockGatheringRepo.EXPECT().CreateGathering(ctx, gomock.Any()).Return(&domain.Gathering{
			ID: 1,
		}, nil)

		res, err := suite.Cs.CreateGathering(ctx, suite.MockCreateGatheringReq)
		suite.NoError(err)
		suite.Equal(res.ID, int64(1))
	})
}

func (suite *TestGatheringServiceSuite) TestGatheringService_AttendGathering() {
	suite.Run("error when getting member from db", func() {
		ctx := context.Background()

		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(nil, errors.New("err from db"))

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("error when getting gathering from db", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(nil, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(nil, errors.New("err from db"))

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("error time not yet passed", func() {
		ctx := context.Background()

		layoutFormat := "2006-01-02 15:04:05"
		value := "2023-10-12 10:04:00"
		v, _ := time.Parse(layoutFormat, value)
		gathering := &domain.Gathering{
			ID:         1,
			Creator:    "",
			Type:       "INVITATION",
			ScheduleAt: v,
		}
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(gathering, nil)

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "gathering not yet started")
	})

	suite.Run("error when getting member from db", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(&domain.Gathering{
			ID:   1,
			Type: "INVITATION",
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindInvitationByGatheringIDAndMemberID(ctx, gomock.Any(), gomock.Any()).Return(
			&domain.Invitation{
				MemberID: 2,
			}, nil)

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "unauthorized")
	})

	suite.Run("error when getting member from db", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(&domain.Gathering{
			ID:   1,
			Type: "INVITATION",
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindInvitationByGatheringIDAndMemberID(ctx, gomock.Any(), gomock.Any()).Return(
			&domain.Invitation{
				MemberID: 1,
			}, nil)
		suite.MockGatheringRepo.EXPECT().CreateAttendee(ctx, gomock.Any()).Return(
			nil, errors.New("err from db"))

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("error when find invitation", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(&domain.Gathering{
			ID:   1,
			Type: "INVITATION",
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindInvitationByGatheringIDAndMemberID(ctx, gomock.Any(), gomock.Any()).Return(
			nil, errors.New("err from db"))

		_, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("success with type is invitation", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(&domain.Gathering{
			ID:   1,
			Type: "INVITATION",
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindInvitationByGatheringIDAndMemberID(ctx, gomock.Any(), gomock.Any()).Return(
			&domain.Invitation{
				MemberID: 1,
			}, nil)
		suite.MockGatheringRepo.EXPECT().CreateAttendee(ctx, gomock.Any()).Return(
			&domain.Attendee{
				ID: 1,
			}, nil)

		res, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.NoError(err)
		suite.Equal(res.ID, int64(1))
	})

	suite.Run("success with type is free", func() {
		ctx := context.Background()
		suite.MockGatheringRepo.EXPECT().FindMemberByID(ctx, gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockGatheringRepo.EXPECT().FindGatheringByID(ctx, gomock.Any()).Return(&domain.Gathering{
			ID:   1,
			Type: domain.GatheringTypeFREE,
		}, nil)
		suite.MockGatheringRepo.EXPECT().CreateAttendee(ctx, gomock.Any()).Return(
			&domain.Attendee{
				ID: 1,
			}, nil)

		res, err := suite.Cs.AttendGathering(ctx, suite.MockTimePassed, suite.MockCreateAttendeeReq)
		suite.NoError(err)
		suite.Equal(res.ID, int64(1))
	})
}
