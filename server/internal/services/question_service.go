package services

import (
	"server/internal/models"
	"server/internal/repository"

	"gorm.io/gorm"
)

type QuestionService struct {
	repo *repository.QuestionRepository
}

func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{
		repo: repository.NewQuestionRepository(db),
	}
}

func (s *QuestionService) GetQuestionByLevel(level int) (*models.Question, error) {
	return s.repo.GetQuestionByLevel(level)
}
