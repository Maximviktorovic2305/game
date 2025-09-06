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
