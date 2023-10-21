// Package postgres is ...
package postgres

import (
	"context"
	"database/sql"

	"github.com/elangreza14/gathering/internal/domain"
)

// https://github.com/DATA-DOG/go-sqlmock

// RepoPostgres is ...
type RepoPostgres struct{ db *sql.DB }

// NewRepoPostgres is ...
func NewRepoPostgres(db *sql.DB) *RepoPostgres {
	return &RepoPostgres{
		db: db,
	}
}

// FindMemberByID is ...
func (r *RepoPostgres) FindMemberByID(ctx context.Context, id int64) (*domain.Member, error) {
	const getMember = `
	SELECT id, email, first_name, last_name FROM members
	WHERE id = $1 LIMIT 1`

	row := r.db.QueryRowContext(ctx, getMember, id)
	i := &domain.Member{}
	err := row.Scan(&i.ID, &i.Email, &i.FirstName, &i.LastName)
	return i, err
}

// FindInvitationByID is ...
func (r *RepoPostgres) FindInvitationByID(ctx context.Context, id int64) (*domain.Invitation, error) {
	const getInvitation = `
	SELECT id, member_id, gathering_id, status FROM invitations
	WHERE id = $1 LIMIT 1`

	row := r.db.QueryRowContext(ctx, getInvitation, id)
	i := &domain.Invitation{}
	err := row.Scan(&i.ID, &i.MemberID, &i.GatheringID, &i.Status)
	return i, err
}

// FindGatheringByID is ...
func (r *RepoPostgres) FindGatheringByID(ctx context.Context, id int64) (*domain.Gathering, error) {
	const getGathering = `
	SELECT id, creator, type, schedule_at, name, location FROM gatherings
	WHERE id = $1 LIMIT 1`

	row := r.db.QueryRowContext(ctx, getGathering, id)
	i := &domain.Gathering{}
	err := row.Scan(&i.ID, &i.Creator, &i.Type, &i.ScheduleAt, &i.Name, &i.Location)
	return i, err
}

// FindInvitationByGatheringIDAndMemberID is ...
func (r *RepoPostgres) FindInvitationByGatheringIDAndMemberID(
	ctx context.Context,
	gatheringID,
	memberID int64,
) (*domain.Invitation, error) {
	const getInvitation = `
	SELECT id, member_id, gathering_id, status FROM invitations
	WHERE gathering_id = $1 AND member_id=$2 LIMIT 1`

	row := r.db.QueryRowContext(ctx, getInvitation, gatheringID, memberID)
	i := &domain.Invitation{}
	err := row.Scan(&i.ID, &i.MemberID, &i.GatheringID, &i.Status)
	return i, err
}

// CreateMember is ...
func (r *RepoPostgres) CreateMember(ctx context.Context, arg domain.Member) (*domain.Member, error) {
	const createAuthor = `
	INSERT INTO members (
	  first_name, last_name, email
	) VALUES (
	  $1, $2, $3
	) RETURNING id, first_name, last_name, email
	`

	row := r.db.QueryRowContext(ctx, createAuthor, arg.FirstName, arg.LastName, arg.Email)
	i := &domain.Member{}
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Email)
	return i, err
}

// CreateGathering is ...
func (r *RepoPostgres) CreateGathering(ctx context.Context, arg domain.Gathering) (*domain.Gathering, error) {
	const createGathering = `
	INSERT INTO gatherings (
	  creator, type, schedule_at, name, location
	) VALUES (
	  $1, $2, $3, $4, $5
	) RETURNING id, creator, type, schedule_at, name, location
	`

	row := r.db.QueryRowContext(ctx, createGathering, arg.Creator, arg.Type, arg.ScheduleAt, arg.Name, arg.Location)
	i := &domain.Gathering{}
	err := row.Scan(&i.ID, &i.Creator, &i.Type, &i.ScheduleAt, &i.Name, &i.Location)
	return i, err
}

// CreateAttendee is ...
func (r *RepoPostgres) CreateAttendee(ctx context.Context, arg domain.Attendee) (*domain.Attendee, error) {
	const createAttendee = `
	INSERT INTO attendees (
	  member_id, gathering_id
	) VALUES (
	  $1, $2
	) RETURNING id, member_id, gathering_id
	`

	row := r.db.QueryRowContext(ctx, createAttendee, arg.MemberID, arg.GatheringID)
	i := &domain.Attendee{}
	err := row.Scan(&i.ID, &i.MemberID, &i.GatheringID)
	return i, err
}

// CreateInvitations is ...
func (r *RepoPostgres) CreateInvitations(
	ctx context.Context,
	gatheringID int64,
	status domain.InvitationStatus,
	memberIDs ...int64,
) error {
	const createInvitation = `
	INSERT INTO invitations (
	  member_id, gathering_id, status
	) VALUES (
	  $1, $2, $3
	) 
	`

	stmt, err := r.db.Prepare(createInvitation)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for i := 0; i < len(memberIDs); i++ {
		if _, err = stmt.ExecContext(ctx, memberIDs[i], gatheringID, status); err != nil {
			return err
		}
	}

	return nil
}

// UpdateInvitation is ...
func (r *RepoPostgres) UpdateInvitation(ctx context.Context, arg domain.Invitation) error {
	const updateInvitation = `
	UPDATE invitations SET status = $2
	WHERE id = $1`

	_, err := r.db.ExecContext(ctx, updateInvitation, arg.ID, arg.Status)
	return err
}
