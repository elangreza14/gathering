// Code generated by MockGen. DO NOT EDIT.
// Source: member.go
//
// Generated by this command:
//
//	mockgen -source member.go -destination ../../mock/controller/mock_member.go -package controller
//
// Package controller is a generated GoMock package.
package controller

import (
	context "context"
	reflect "reflect"

	dto "github.com/elangreza14/gathering/internal/dto"
	gomock "go.uber.org/mock/gomock"
)

// MockmemberService is a mock of memberService interface.
type MockmemberService struct {
	ctrl     *gomock.Controller
	recorder *MockmemberServiceMockRecorder
}

// MockmemberServiceMockRecorder is the mock recorder for MockmemberService.
type MockmemberServiceMockRecorder struct {
	mock *MockmemberService
}

// NewMockmemberService creates a new mock instance.
func NewMockmemberService(ctrl *gomock.Controller) *MockmemberService {
	mock := &MockmemberService{ctrl: ctrl}
	mock.recorder = &MockmemberServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmemberService) EXPECT() *MockmemberServiceMockRecorder {
	return m.recorder
}

// CreateMember mocks base method.
func (m *MockmemberService) CreateMember(ctx context.Context, req dto.CreateMemberReq) (*dto.CreateMemberRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", ctx, req)
	ret0, _ := ret[0].(*dto.CreateMemberRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockmemberServiceMockRecorder) CreateMember(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockmemberService)(nil).CreateMember), ctx, req)
}

// RespondInvitation mocks base method.
func (m *MockmemberService) RespondInvitation(ctx context.Context, req dto.RespondInvitationReq) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RespondInvitation", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondInvitation indicates an expected call of RespondInvitation.
func (mr *MockmemberServiceMockRecorder) RespondInvitation(ctx, req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondInvitation", reflect.TypeOf((*MockmemberService)(nil).RespondInvitation), ctx, req)
}
