package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
)

const (
	usersTable        = "users"
	segmentsTable     = "segments"
	userSegmentsTable = "users_segments"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUser(user ent.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// delete user

func (r *UserPostgres) GetUserSegments(userId int) ([]ent.Segment, error) {
	var segments []ent.Segment
	query := fmt.Sprintf("SELECT s.name FROM %s s INNER JOIN %s us on s.id = us.segment_id WHERE us.user_id = $1", segmentsTable, userSegmentsTable)
	err := r.db.Select(&segments, query, userId)
	return segments, err
}
