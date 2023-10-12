package service_test

import (
	"errors"
	"testing"

	"github.com/elangreza14/gathering/internal/domain"
	"github.com/elangreza14/gathering/internal/dto"
	. "github.com/elangreza14/gathering/internal/service"
	gomockService "github.com/elangreza14/gathering/mock/service"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type TestMemberServiceSuite struct {
	suite.Suite

	MockMemberRepo *gomockService.MockmemberRepo

	MockCreateMemberReq      dto.CreateMemberReq
	MockRespondInvitationReq dto.RespondInvitationReq
	Cs                       *MemberService
	Ctrl                     *gomock.Controller
}

func (suite *TestMemberServiceSuite) SetupSuite() {
	suite.Ctrl = gomock.NewController(suite.T())
	suite.MockMemberRepo = gomockService.NewMockmemberRepo(suite.Ctrl)
	suite.MockCreateMemberReq = dto.CreateMemberReq{
		FirstName: "FirstName",
		LastName:  "LastName",
		Email:     "Email",
	}
	suite.MockRespondInvitationReq = dto.RespondInvitationReq{
		MemberID:     1,
		InvitationID: 1,
		Attend:       true,
	}
	suite.Cs = NewMemberService(suite.MockMemberRepo)
}

func (suite *TestMemberServiceSuite) TearDownSuite() {
	suite.Ctrl.Finish()
}

func TestMemberServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TestMemberServiceSuite))
}

func (suite *TestMemberServiceSuite) TestMemberService_CreateMember() {
	suite.Run("err from db", func() {
		suite.MockMemberRepo.EXPECT().CreateMember(gomock.Any()).Return(nil, errors.New("err from db"))

		_, err := suite.Cs.CreateMember(suite.MockCreateMemberReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("success", func() {
		suite.MockMemberRepo.EXPECT().CreateMember(gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)

		res, err := suite.Cs.CreateMember(suite.MockCreateMemberReq)
		suite.NoError(err)
		suite.Equal(res.ID, int64(1))
	})
}

func (suite *TestMemberServiceSuite) TestMemberService_AttendMember() {
	suite.Run("FindMemberByID err from db", func() {
		suite.MockMemberRepo.EXPECT().FindMemberByID(gomock.Any()).Return(nil, errors.New("err from db"))

		err := suite.Cs.RespondInvitation(suite.MockRespondInvitationReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("FindInvitationByID err from db", func() {
		suite.MockMemberRepo.EXPECT().FindMemberByID(gomock.Any()).Return(nil, nil)
		suite.MockMemberRepo.EXPECT().FindInvitationByID(gomock.Any()).Return(nil, errors.New("err from db"))

		err := suite.Cs.RespondInvitation(suite.MockRespondInvitationReq)
		suite.Error(err)
		suite.Equal(err.Error(), "err from db")
	})

	suite.Run("unauthorized", func() {
		suite.MockMemberRepo.EXPECT().FindMemberByID(gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockMemberRepo.EXPECT().FindInvitationByID(gomock.Any()).Return(&domain.Invitation{
			MemberID: 2,
		}, nil)

		err := suite.Cs.RespondInvitation(suite.MockRespondInvitationReq)
		suite.Error(err)
		suite.Equal(err.Error(), "unauthorized")
	})

	suite.Run("error when update invitation", func() {
		suite.MockRespondInvitationReq.Attend = true
		suite.MockMemberRepo.EXPECT().FindMemberByID(gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockMemberRepo.EXPECT().FindInvitationByID(gomock.Any()).Return(&domain.Invitation{
			MemberID: 1,
		}, nil)
		suite.MockMemberRepo.EXPECT().UpdateInvitation(gomock.Any()).Return(errors.New("err from db"))

		err := suite.Cs.RespondInvitation(suite.MockRespondInvitationReq)
		suite.Error(err)
	})

	suite.Run("success update invitation", func() {
		suite.MockRespondInvitationReq.Attend = false
		suite.MockMemberRepo.EXPECT().FindMemberByID(gomock.Any()).Return(&domain.Member{
			ID: 1,
		}, nil)
		suite.MockMemberRepo.EXPECT().FindInvitationByID(gomock.Any()).Return(&domain.Invitation{
			MemberID: 1,
		}, nil)
		suite.MockMemberRepo.EXPECT().UpdateInvitation(gomock.Any()).Return(nil)

		err := suite.Cs.RespondInvitation(suite.MockRespondInvitationReq)
		suite.NoError(err)
	})
}
