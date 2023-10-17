package controller

//go:generate mockgen -source $GOFILE -destination ../../mock/controller/mock_$GOFILE -package $GOPACKAGE

import (
	"context"
	"net/http"

	"github.com/elangreza14/gathering/internal/dto"
	"github.com/gin-gonic/gin"
)

type memberService interface {
	CreateMember(ctx context.Context, req dto.CreateMemberReq) (*dto.CreateMemberRes, error)
	RespondInvitation(ctx context.Context, req dto.RespondInvitationReq) error
}

type MemberController struct {
	memberService memberService
}

func NewMemberController(service memberService) *MemberController {
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

		res, err := mc.memberService.CreateMember(c.Request.Context(), json)
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

		err := mc.memberService.RespondInvitation(c.Request.Context(), json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": "ok"})
	}
}
