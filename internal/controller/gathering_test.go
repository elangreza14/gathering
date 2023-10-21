package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/elangreza14/gathering/internal/controller"
	"github.com/elangreza14/gathering/internal/dto"
	GatheringController "github.com/elangreza14/gathering/mock/controller"

	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/suite"
)

type TestGatheringControllerSuite struct {
	suite.Suite

	Ctrl                 *gomock.Controller
	MockGatheringService *GatheringController.MockgatheringService
}

func (suite *TestGatheringControllerSuite) SetupSuite() {
	suite.Ctrl = gomock.NewController(suite.T())
	suite.MockGatheringService = GatheringController.NewMockgatheringService(suite.Ctrl)
}

func (suite *TestGatheringControllerSuite) TearDownSuite() {
	suite.Ctrl.Finish()
}

func TestGatheringControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TestGatheringControllerSuite))
}

func (suite *TestGatheringControllerSuite) TestGatheringController_CreateGathering() {
	r := SetUpRouter()
	gatheringController := controller.NewGatheringController(suite.MockGatheringService)
	gathering := r.Group("/v1")
	gathering.PUT("/gathering", gatheringController.CreateGathering())

	requestBody := dto.CreateGatheringReq{
		Creator:    "1",
		Type:       "FREE",
		ScheduleAt: time.Now(),
		Name:       "1",
		Location:   "1",
		Attendees:  []int64{1},
	}
	payload, _ := json.Marshal(requestBody)

	suite.Run("error from validation", func() {
		errRequestBody := dto.CreateGatheringReq{
			Creator: "",
			Type:    "",
		}
		errPayload, _ := json.Marshal(errRequestBody)

		bodyReader := bytes.NewReader(errPayload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"result":"errors","error":[{"field":"Creator","message":"This field is required"},{"field":"Type","message":"Should be FREE or INVITATION"},{"field":"ScheduleAt","message":"This field is required"},{"field":"Name","message":"This field is required"},{"field":"Location","message":"This field is required"}]}`, string(responseData))
		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("error from service", func() {
		suite.MockGatheringService.EXPECT().
			CreateGathering(gomock.Any(), gomock.Any()).Return(nil, errors.New("error from db"))

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"result":"error","error":"error from db"}`, string(responseData))
		suite.Equal(http.StatusInternalServerError, w.Code)
	})

	suite.Run("success", func() {
		suite.MockGatheringService.EXPECT().CreateGathering(gomock.Any(), gomock.Any()).Return(
			&dto.CreateGatheringRes{
				ID: 1,
			}, nil)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"data":{"id":1},"result":"ok"}`, string(responseData))
		suite.Equal(http.StatusCreated, w.Code)
	})
}

func (suite *TestGatheringControllerSuite) TestGatheringController_AttendGathering() {
	r := SetUpRouter()
	gatheringController := controller.NewGatheringController(suite.MockGatheringService)
	gathering := r.Group("/v1")
	gathering.PUT("/gathering/invitation", gatheringController.AttendGathering())

	requestBody := dto.CreateAttendeeReq{
		MemberID:    1,
		GatheringID: 1,
	}
	payload, _ := json.Marshal(requestBody)

	suite.Run("error from validation", func() {
		errRequestBody := dto.CreateAttendeeReq{
			MemberID: 1,
		}
		errPayload, _ := json.Marshal(errRequestBody)

		bodyReader := bytes.NewReader(errPayload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering/invitation", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"result":"errors","error":[{"field":"GatheringID","message":"This field is required"}]}`, string(responseData))
		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("error from service", func() {
		suite.MockGatheringService.EXPECT().AttendGathering(gomock.Any(), gomock.Any(), gomock.Any()).Return(
			nil, errors.New("error from db"),
		)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering/invitation", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"result":"error","error":"error from db"}`, string(responseData))
		suite.Equal(http.StatusInternalServerError, w.Code)
	})

	suite.Run("success", func() {
		suite.MockGatheringService.EXPECT().AttendGathering(gomock.Any(), gomock.Any(), gomock.Any()).Return(
			&dto.CreateAttendeeRes{
				ID: 1,
			}, nil,
		)

		bodyReader := bytes.NewReader(payload)
		req, _ := http.NewRequest(http.MethodPut, "/v1/gathering/invitation", bodyReader)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responseData, _ := io.ReadAll(w.Body)
		suite.Equal(`{"data":{"id":1},"result":"ok"}`, string(responseData))
		suite.Equal(http.StatusCreated, w.Code)
	})
}
