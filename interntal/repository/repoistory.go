package repository

import (
	"github.com/jmoiron/sqlx"

	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
)

type SegmentRequest struct {
	Segments []string `json:"segments"`
}

type User interface {
	CreateUser(user ent.User) (int, error)
	GetUserSegments(userId int) ([]ent.Segment, error)
	// GetUser(username, password string) (ent.User, error)
}

type Segment interface {
	CreateSegment(segment ent.Segment) (int, error)
	DeleteSegment(segment ent.Segment) error
	AddUserSegment(userId int, segments SegmentRequest) ([]int, error)
	DeleteUserSegment(userId int, segments SegmentRequest) error
}

type Repository struct {
	User
	Segment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserPostgres(db),
		Segment: NewSegmentPostgres(db),
	}

}
