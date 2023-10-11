package repoPostgres

import (
	"database/sql"

	"github.com/elangreza14/gathering/internal/domain"
)

// https://github.com/DATA-DOG/go-sqlmock

type RepoPostgres struct{ db *sql.DB }

func New(db *sql.DB) *RepoPostgres {
	return &RepoPostgres{
		db: db,
	}
}

func (r *RepoPostgres) FindMemberByID(ID int64) (*domain.Member, error) {
	return nil, nil
}

func (r *RepoPostgres) FindInvitationByID(ID int64) (*domain.Invitation, error) {
	return nil, nil
}

func (r *RepoPostgres) FindGatheringByID(ID int64) (*domain.Gathering, error) {
	return nil, nil
}

func (r *RepoPostgres) FindInvitationByGatheringIDAndMemberID(gatheringID, memberID int64) (*domain.Invitation, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateMember(domain.Member) (*domain.Member, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateGathering(domain.Gathering) (*domain.Gathering, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateAttendee(domain.Attendee) (*domain.Attendee, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateInvitations(gatheringID int64, status string, memberID ...int64) error {
	return nil
}

func (r *RepoPostgres) UpdateInvitation(domain.Invitation) error {
	return nil
}
