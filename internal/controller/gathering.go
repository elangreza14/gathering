package controller

import (
	"net/http"

	"github.com/elangreza14/gathering/internal/dto"
	service "github.com/elangreza14/gathering/internal/service"
	"github.com/gin-gonic/gin"
)

type GatheringController struct {
	gatheringService *service.GatheringService
}

func NewGatheringController(service *service.GatheringService) *GatheringController {
	return &GatheringController{
		gatheringService: service,
	}
}

func (mc *GatheringController) CreateGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateGatheringReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		res, err := mc.gatheringService.CreateGathering(json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": res})
	}
}

func (mc *GatheringController) AttendGathering() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateAttendeeReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		res, err := mc.gatheringService.AttendGathering(json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": res})
	}
}
