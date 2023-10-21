// Package controller is ...
package controller

//go:generate mockgen -source $GOFILE -destination ../../mock/controller/mock_$GOFILE -package $GOPACKAGE

import (
	"context"
	"net/http"
	"time"

	"github.com/elangreza14/gathering/internal/dto"
	"github.com/gin-gonic/gin"
)

// GatheringController is ...
type GatheringController struct {
	gatheringService gatheringService
}

type gatheringService interface {
	AttendGathering(context.Context, time.Time, dto.CreateAttendeeReq) (*dto.CreateAttendeeRes, error)
	CreateGathering(context.Context, dto.CreateGatheringReq) (*dto.CreateGatheringRes, error)
}

// NewGatheringController is ...
func NewGatheringController(service gatheringService) *GatheringController {
	return &GatheringController{
		gatheringService: service,
	}
}

// CreateGathering is ...
// CreateGathering godoc
//
//	@Summary		create gathering
//	@Description	create gathering
//	@Tags			gathering
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateGatheringReq								true	"test"
//	@success		200		{object}	dto.SuccessResponse{data=dto.CreateGatheringRes}	"success"
//	@Failure		400		{object}	dto.ErrorResponse{error=[]dto.ErrorField}			"error validation"
//	@Failure		500		{object}	dto.ErrorResponse{error=string}						"error internal"
//	@Router			/gathering [post]
func (mc *GatheringController) CreateGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateGatheringReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		res, err := mc.gatheringService.CreateGathering(c.Request.Context(), json)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusCreated, dto.NewBaseResponse(res, nil))
	}
}

// AttendGathering is ...
// AttendGathering godoc
//
//	@Summary		create gathering invitation
//	@Description	create gathering invitation
//	@Tags			gathering
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateAttendeeReq							true	"body"
//	@success		200		{object}	dto.SuccessResponse{data=dto.CreateAttendeeRes}	"success"
//	@Failure		400		{object}	dto.ErrorResponse{error=[]dto.ErrorField}		"error validation"
//	@Failure		500		{object}	dto.ErrorResponse{error=string}					"error internal"
//	@Router			/gathering/invitation [post]
func (mc *GatheringController) AttendGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateAttendeeReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		res, err := mc.gatheringService.AttendGathering(c.Request.Context(), time.Now(), json)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusCreated, dto.NewBaseResponse(res, nil))
	}
}
