package controller_test

import (
	"bytes"
	"encoding/json"
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

	suite.Run("success", func() {
		requestBody := dto.CreateMemberReq{
			FirstName: "test",
			LastName:  "test",
			Email:     "test",
		}
		payload, _ := json.Marshal(requestBody)

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
