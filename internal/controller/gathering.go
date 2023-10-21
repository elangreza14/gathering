// Package controller is ...
package controller

//go:generate mockgen -source $GOFILE -destination ../../mock/controller/mock_$GOFILE -package $GOPACKAGE

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/elangreza14/gathering/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

// CreateGathering ...
func (mc *GatheringController) CreateGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateGatheringReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		res, err := mc.gatheringService.CreateGathering(c.Request.Context(), json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": res})
	}
}

// AttendGathering is ...
func (mc *GatheringController) AttendGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateAttendeeReq
		if err := c.ShouldBindJSON(&json); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				out := make([]ErrorMsg, len(ve))
				for i, fe := range ve {
					out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
				}
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "cause": out})
			}
			return
		}

		res, err := mc.gatheringService.AttendGathering(c.Request.Context(), time.Now(), json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": res})
	}
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}
