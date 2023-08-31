package service

import (
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	"github.com/naumovrus/backend-trainee-asignment/interntal/repository"
)

type UserService struct {
	repo repository.User
}

func NewAuthService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user ent.User) (int, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserSegments(userId int) ([]ent.Segment, error) {
	return s.repo.GetUserSegments(userId)
}
