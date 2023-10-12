package repoPostgres

import (
	"context"
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

func (r *RepoPostgres) FindMemberByID(ctx context.Context, ID int64) (*domain.Member, error) {
	return nil, nil
}

func (r *RepoPostgres) FindInvitationByID(ctx context.Context, ID int64) (*domain.Invitation, error) {
	return nil, nil
}

func (r *RepoPostgres) FindGatheringByID(ctx context.Context, ID int64) (*domain.Gathering, error) {
	return nil, nil
}

func (r *RepoPostgres) FindInvitationByGatheringIDAndMemberID(ctx context.Context, gatheringID, memberID int64) (*domain.Invitation, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateMember(ctx context.Context, arg domain.Member) (*domain.Member, error) {
	const createAuthor = `
	INSERT INTO members (
	  first_name, last_name, email
	) VALUES (
	  $1, $2, $3
	)
	`

	row := r.db.QueryRowContext(ctx, createAuthor, arg.FirstName, arg.LastName, arg.Email)
	i := &domain.Member{}
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Email)
	return i, err
}

func (r *RepoPostgres) CreateGathering(ctx context.Context, arg domain.Gathering) (*domain.Gathering, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateAttendee(ctx context.Context, arg domain.Attendee) (*domain.Attendee, error) {
	return nil, nil
}

func (r *RepoPostgres) CreateInvitations(ctx context.Context, gatheringID int64, status string, memberID ...int64) error {
	return nil
}

func (r *RepoPostgres) UpdateInvitation(ctx context.Context, arg domain.Invitation) error {
	return nil
}
