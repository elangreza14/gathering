// Package Main is ...
//
//nolint:errcheck
package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	controller "github.com/elangreza14/gathering/internal/controller"
	repo "github.com/elangreza14/gathering/internal/postgres"
	service "github.com/elangreza14/gathering/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type (
	// Env is list all of env.
	Env struct {
		PostgresHostname string `mapstructure:"POSTGRES_HOSTNAME"`
		PostgresSsl      string `mapstructure:"POSTGRES_SSL"`
		PostgresUser     string `mapstructure:"POSTGRES_USER"`
		PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
		PostgresDB       string `mapstructure:"POSTGRES_DB"`
		PostgresPort     int32  `mapstructure:"POSTGRES_PORT"`
		MigrationFolder  string `mapstructure:"MIGRATION_FOLDER"`
	}
)

func main() {
	db, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.NewRepoPostgres(db)

	router := gin.Default()
	v1 := router.Group("/v1")

	memberController := controller.NewMemberController(service.NewMemberService(repo))
	member := v1.Group("/member")
	member.POST("/", memberController.CreateMember())
	member.PUT("/invitation", memberController.RespondInvitation())

	gatheringController := controller.NewGatheringController(service.NewGatheringService(repo))
	gathering := v1.Group("/gathering")
	gathering.POST("/", gatheringController.CreateGathering())
	gathering.PUT("/attends", gatheringController.AttendGathering())

	router.Run(":5000")
}

func setup() (*sql.DB, error) {
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
		return nil, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		return nil, err
	}

	dbURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		env.PostgresUser,
		env.PostgresPassword,
		env.PostgresHostname,
		env.PostgresPort,
		env.PostgresDB,
		env.PostgresSsl)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	migration, err := migrate.New(
		fmt.Sprintf("file://%v", env.MigrationFolder), dbURL)
	if err != nil {
		return nil, err
	}

	// apply migration to DB
	if err = migration.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			return nil, err
		}
	}

	const maxConn = 25
	const maxLifeTime = 5 * time.Minute

	db.SetMaxOpenConns(maxConn)
	db.SetMaxIdleConns(maxConn)
	db.SetConnMaxLifetime(maxLifeTime)

	return db, nil
}
