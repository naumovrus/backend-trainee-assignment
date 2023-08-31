package service

import (
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	"github.com/naumovrus/backend-trainee-asignment/interntal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type User interface {
	CreateUser(user ent.User) (int, error)
	GetUserSegments(userId int) ([]ent.Segment, error)
	// GetUser(username, password string) (ent.User, error)
}

type Segment interface {
	CreateSegment(segment ent.Segment) (int, error)
	DeleteSegment(segment ent.Segment) error
	AddUserSegment(userId int, segments repository.SegmentRequest) ([]int, error)
	DeleteUserSegment(userId int, segments repository.SegmentRequest) error
	// GetUserSegments(userId int) ([]ent.Segment, error)

}

type Service struct {
	User
	Segment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:    NewAuthService(repos.User),
		Segment: NewSegmentService(repos.Segment),
	}

}
