package main

import (
	"database/sql"
	"log"

	repoPostgres "github.com/elangreza14/gathering/internal/repo"
	gathering "github.com/elangreza14/gathering/internal/service/gathering"
	member "github.com/elangreza14/gathering/internal/service/member"
)

func main() {
	db, err := sql.Open("driver-name", "database=test1")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	repo := repoPostgres.New(db)
	gathering.NewGatheringService(repo)
	member.NewMemberService(repo)
}
