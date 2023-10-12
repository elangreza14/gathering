package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	controller "github.com/elangreza14/gathering/internal/controller"
	repoPostgres "github.com/elangreza14/gathering/internal/repo"
	service "github.com/elangreza14/gathering/internal/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type (
	// Env is list all of env
	Env struct {
		PostgresHostname string `mapstructure:"POSTGRES_HOSTNAME"`
		PostgresSsl      string `mapstructure:"POSTGRES_SSL"`
		PostgresUser     string `mapstructure:"POSTGRES_USER"`
		PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
		PostgresDB       string `mapstructure:"POSTGRES_DB"`
		PostgresPort     int32  `mapstructure:"POSTGRES_PORT"`
	}
)

func main() {
	env := &Env{}
	envBase := "local"
	mode := os.Getenv("MODE")
	if mode != "" {
		envBase = mode
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(envBase)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresHostname,
		env.PostgresPort,
		env.PostgresDB,
		env.PostgresSsl)
	db, err := sql.Open("postgres", dbUrl)
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
