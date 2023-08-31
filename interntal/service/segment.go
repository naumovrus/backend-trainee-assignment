package service

import (
	ent "github.com/naumovrus/backend-trainee-asignment/interntal/entities"
	"github.com/naumovrus/backend-trainee-asignment/interntal/repository"
)

type SegmentService struct {
	repo repository.Segment
}

func NewSegmentService(repo repository.Segment) *SegmentService {
	return &SegmentService{repo: repo}
}

func (s *SegmentService) CreateSegment(segment ent.Segment) (int, error) {
	return s.repo.CreateSegment(segment)
}

func (s *SegmentService) AddUserSegment(userId int, segments repository.SegmentRequest) ([]int, error) {
	return s.repo.AddUserSegment(userId, segments)
}

func (s *SegmentService) DeleteUserSegment(userId int, segments repository.SegmentRequest) error {
	return s.repo.DeleteUserSegment(userId, segments)
}

func (s *SegmentService) DeleteSegment(segment ent.Segment) error {
	return s.repo.DeleteSegment(segment)
}
