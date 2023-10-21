// Package controller is ...
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

// MemberController is ..
type MemberController struct {
	memberService memberService
}

// NewMemberController is ...
func NewMemberController(service memberService) *MemberController {
	return &MemberController{
		memberService: service,
	}
}

// CreateMember is ...
// CreateMember godoc
//
//	@Summary		create member
//	@Description	create member
//	@Tags			member
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.CreateMemberReq								true	"body"
//	@success		200		{object}	dto.SuccessResponse{data=dto.CreateMemberRes}	"success"
//	@Failure		400		{object}	dto.ErrorResponse{error=[]dto.ErrorField}		"error validation"
//	@Failure		500		{object}	dto.ErrorResponse{error=string}					"error internal"
//	@Router			/member [post]
func (mc *MemberController) CreateMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.CreateMemberReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		res, err := mc.memberService.CreateMember(c.Request.Context(), json)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"result": "error", "cause": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"result": "ok", "data": res})
	}
}

// RespondInvitation is ...
// RespondInvitation godoc
//
//	@Summary		create member
//	@Description	create member
//	@Tags			member
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RespondInvitationReq					true	"body"
//	@success		200		{object}	dto.SuccessResponsePlain{}					"success"
//	@Failure		400		{object}	dto.ErrorResponse{error=[]dto.ErrorField}	"error validation"
//	@Failure		500		{object}	dto.ErrorResponse{error=string}				"error internal"
//	@Router			/member/invitation [post]
func (mc *MemberController) RespondInvitation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var json dto.RespondInvitationReq
		if err := c.ShouldBindJSON(&json); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		err := mc.memberService.RespondInvitation(c.Request.Context(), json)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewBaseResponse(nil, err))
			return
		}

		c.JSON(http.StatusCreated, dto.NewBaseResponse(nil, nil))
	}
}
