package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elangreza14/gathering/internal/controller"
	"github.com/elangreza14/gathering/internal/dto"
	MemberController "github.com/elangreza14/gathering/mock/controller"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/suite"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

type TestMemberControllerSuite struct {
	suite.Suite

	Ctrl              *gomock.Controller
	MockMemberService *MemberController.MockmemberService
}

func (suite *TestMemberControllerSuite) SetupSuite() {
	suite.Ctrl = gomock.NewController(suite.T())
	suite.MockMemberService = MemberController.NewMockmemberService(suite.Ctrl)
}

func (suite *TestMemberControllerSuite) TearDownSuite() {
	suite.Ctrl.Finish()
}

func TestMemberControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TestMemberControllerSuite))
}

func (suite *TestMemberControllerSuite) TestMemberController_Register() {
	r := SetUpRouter()
	memberController := controller.NewMemberController(suite.MockMemberService)
	member := r.Group("/v1")
	member.POST("/member", memberController.CreateMember())

	requestBody := dto.CreateMemberReq{
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
	}
	payload, _ := json.Marshal(requestBody)

	suite.Run("error validation", func() {
		requestBody := dto.CreateMemberReq{
			FirstName: "test",
			LastName:  "test",
		}
		payload, _ := json.Marshal(requestBody)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/member", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal("{\"cause\":\"Key: 'CreateMemberReq.Email' Error:Field validation for 'Email' failed on the 'required' tag\",\"status\":\"error\"}", string(responseData))
		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("error from service", func() {
		suite.MockMemberService.EXPECT().CreateMember(gomock.Any(), gomock.Any()).Return(
			nil, errors.New("errors from db"),
		)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/member", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal("{\"cause\":\"errors from db\",\"status\":\"error\"}", string(responseData))
		suite.Equal(http.StatusInternalServerError, w.Code)
	})

	suite.Run("success", func() {
		suite.MockMemberService.EXPECT().CreateMember(gomock.Any(), gomock.Any()).Return(
			&dto.CreateMemberRes{
				ID: 1,
			}, nil,
		)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPost, "/v1/member", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal("{\"data\":{\"id\":1},\"status\":\"ok\"}", string(responseData))
		suite.Equal(http.StatusCreated, w.Code)
	})
}
