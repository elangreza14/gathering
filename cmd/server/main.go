package main

import (
	"database/sql"
	"log"

	controller "github.com/elangreza14/gathering/internal/controller"
	repoPostgres "github.com/elangreza14/gathering/internal/repo"
	service "github.com/elangreza14/gathering/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	repo := repoPostgres.New(db)

	router := gin.Default()
	v1 := router.Group("/v1")

	memberController := controller.NewMemberController(service.NewMemberService(repo))
	member := v1.Group("/member")
	member.POST("/", memberController.CreateMember())
	member.PUT("/accept-invitation", memberController.RespondInvitation())

	gatheringController := controller.NewGatheringController(service.NewGatheringService(repo))
	gathering := v1.Group("/gathering")
	gathering.POST("/", gatheringController.CreateGathering())
	gathering.PUT("/attend-gathering", gatheringController.AttendGathering())

	router.Run(":5000")
}
