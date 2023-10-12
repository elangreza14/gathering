package controller

import (
	"net/http"

	"github.com/elangreza14/gathering/internal/dto"
	service "github.com/elangreza14/gathering/internal/service"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
	memberService *service.MemberService
}

func NewMemberController(service *service.MemberService) *MemberController {
	return &MemberController{
		memberService: service,
	}
}

func (mc *MemberController) CreateMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateMemberReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		res, err := mc.memberService.CreateMember(json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok", "data": res})
	}
}

func (mc *MemberController) RespondInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.RespondInvitationReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		err := mc.memberService.RespondInvitation(json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok"})
	}
}
