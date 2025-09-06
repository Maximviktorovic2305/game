package repository

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type GameSessionRepository struct {
	db *gorm.DB
}

func NewGameSessionRepository(db *gorm.DB) *GameSessionRepository {
	return &GameSessionRepository{db: db}
}

func (r *GameSessionRepository) CreateSession(session *models.GameSession) error {
	return r.db.Create(session).Error
}

func (r *GameSessionRepository) GetSessionByID(id uint) (*models.GameSession, error) {
	var session models.GameSession
	err := r.db.First(&session, id).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

func (r *GameSessionRepository) UpdateSession(session *models.GameSession) error {
	return r.db.Save(session).Error
}

func (r *GameSessionRepository) GetLeaderboard() ([]models.GameSession, error) {
	var sessions []models.GameSession
	err := r.db.Where("status IN ?", []string{"won", "lost", "quit"}).
		Order("score DESC").
		Limit(10).
		Find(&sessions).Error
	return sessions, err
}

// ValidateOption checks if an option exists and belongs to a question from the specified level
func (r *GameSessionRepository) ValidateOption(optionID uint, level int) (*models.Option, error) {
	var option models.Option
	err := r.db.Where("id = ? AND question_id IN (SELECT id FROM questions WHERE level = ?)", optionID, level).First(&option).Error
	if err != nil {
		return nil, err
	}
	return &option, nil
}

// GetQuestionByID fetches a question by its ID
func (r *GameSessionRepository) GetQuestionByID(id uint) (*models.Question, error) {
	var question models.Question
	err := r.db.First(&question, id).Error
	if err != nil {
		return nil, err
	}
	return &question, nil
}
