package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) GetQuestionByLevel(level int) (*models.Question, error) {
	var question models.Question
	err := r.db.Where("level = ?", level).Preload("Options").First(&question).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}

func (r *QuestionRepository) CreateQuestion(question *models.Question) error {
	return r.db.Create(question).Error
}
