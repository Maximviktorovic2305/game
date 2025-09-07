package services

import (
	"math/rand"
	"server/internal/models"
	"server/internal/repository"
	"time"

	"gorm.io/gorm"
)

type GameService struct {
	sessionRepo  *repository.GameSessionRepository
	questionRepo *repository.QuestionRepository
}

func NewGameService(db *gorm.DB) *GameService {
	return &GameService{
		sessionRepo:  repository.NewGameSessionRepository(db),
		questionRepo: repository.NewQuestionRepository(db),
	}
}

func (s *GameService) StartNewGame() (*models.GameSession, error) {
	session := &models.GameSession{
		CurrentLevel: 1,
		Score:        0,
		Lifelines: models.Lifelines{
			FiftyFifty: true,
			Audience:   true,
			Call:       true,
		},
		Status:    "in_progress",
		StartTime: time.Now(),
	}

	err := s.sessionRepo.CreateSession(session)
	return session, err
}

func (s *GameService) GetSessionByID(id uint) (*models.GameSession, error) {
	return s.sessionRepo.GetSessionByID(id)
}

func (s *GameService) AnswerQuestion(sessionID uint, questionID uint, selectedOptionID uint) (bool, *models.GameSession, error) {
	session, err := s.sessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return false, nil, err
	}

	// Instead of validating against a specific question, we just need to verify
	// that the selected option exists and belongs to a question from the current level
	selectedOption, err := s.sessionRepo.ValidateOption(selectedOptionID, int(session.CurrentLevel))
	if err != nil {
		return false, nil, gorm.ErrRecordNotFound
	}

	isCorrect := selectedOption.IsCorrect

	// Get the question for scoring purposes
	question, err := s.sessionRepo.GetQuestionByID(selectedOption.QuestionID)
	if err != nil {
		return false, nil, err
	}

	// Update session based on answer
	if isCorrect {
		session.CurrentLevel++
		session.Score = question.Prize

		// Check if game is won
		if session.CurrentLevel > 15 {
			session.Status = "won"
			endTime := time.Now()
			session.EndTime = &endTime
		}
	} else {
		session.Status = "lost"
		endTime := time.Now()
		session.EndTime = &endTime

		// Apply safety net rule
		if session.CurrentLevel > 5 {
			session.Score = 1000
		} else {
			session.Score = 0
		}
	}

	err = s.sessionRepo.UpdateSession(session)
	return isCorrect, session, err
}

func (s *GameService) UseFiftyFifty(sessionID uint, questionID uint) ([]models.Option, error) {
	session, err := s.sessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	if !session.Lifelines.FiftyFifty {
		return nil, gorm.ErrRecordNotFound
	}

	// Get the specific question by ID to ensure consistency with what the user sees
	question, err := s.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		return nil, err
	}

	// Verify that the question is from the current level
	if question.Level != int(session.CurrentLevel) {
		return nil, gorm.ErrRecordNotFound
	}

	// Mark lifeline as used
	session.Lifelines.FiftyFifty = false
	err = s.sessionRepo.UpdateSession(session)
	if err != nil {
		return nil, err
	}

	// Find incorrect options
	var incorrectOptions []models.Option
	for _, option := range question.Options {
		if !option.IsCorrect {
			incorrectOptions = append(incorrectOptions, option)
		}
	}

	// Randomly remove 2 incorrect options
	if len(incorrectOptions) >= 2 {
		// Shuffle and take first 2
		for i := range incorrectOptions {
			j := rand.Intn(i + 1)
			incorrectOptions[i], incorrectOptions[j] = incorrectOptions[j], incorrectOptions[i]
		}
		incorrectOptions = incorrectOptions[:2]
	}

	return incorrectOptions, nil
}

func (s *GameService) UseAudience(sessionID uint, questionID uint) (map[string]int, error) {
	session, err := s.sessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	if !session.Lifelines.Audience {
		return nil, gorm.ErrRecordNotFound
	}

	// Get the specific question by ID to ensure consistency with what the user sees
	question, err := s.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		return nil, err
	}

	// Verify that the question is from the current level
	if question.Level != int(session.CurrentLevel) {
		return nil, gorm.ErrRecordNotFound
	}

	// Mark lifeline as used
	session.Lifelines.Audience = false
	err = s.sessionRepo.UpdateSession(session)
	if err != nil {
		return nil, err
	}

	// Generate audience percentages
	percentages := make(map[string]int)
	total := 0

	for _, option := range question.Options {
		if option.IsCorrect {
			// Higher percentage for correct answer
			percent := 70 + rand.Intn(21) // 70-90%
			percentages[option.Letter] = percent
			total += percent
		} else {
			// Lower percentages for incorrect answers
			percent := rand.Intn(16) // 0-15%
			percentages[option.Letter] = percent
			total += percent
		}
	}

	// Adjust to ensure total is 100%
	if total != 100 {
		for letter := range percentages {
			if question.Options[0].Letter == letter { // Adjust the first option
				percentages[letter] += 100 - total
				break
			}
		}
	}

	return percentages, nil
}

func (s *GameService) UseCall(sessionID uint, questionID uint) (string, error) {
	session, err := s.sessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return "", err
	}

	if !session.Lifelines.Call {
		return "", gorm.ErrRecordNotFound
	}

	// Get the specific question by ID to ensure consistency with what the user sees
	question, err := s.questionRepo.GetQuestionByID(questionID)
	if err != nil {
		return "", err
	}

	// Verify that the question is from the current level
	if question.Level != int(session.CurrentLevel) {
		return "", gorm.ErrRecordNotFound
	}

	// Mark lifeline as used
	session.Lifelines.Call = false
	err = s.sessionRepo.UpdateSession(session)
	if err != nil {
		return "", err
	}

	// Generate friend's advice
	var correctOption models.Option
	for _, option := range question.Options {
		if option.IsCorrect {
			correctOption = option
			break
		}
	}

	// High probability of correct answer
	if rand.Intn(100) < 80 { // 80% chance of correct advice
		return "Я почти уверен, что правильный ответ " + correctOption.Letter + ": " + correctOption.Text, nil
	} else {
		// 20% chance of incorrect advice
		for _, option := range question.Options {
			if !option.IsCorrect {
				return "Я думаю, ответ может быть " + option.Letter + ": " + option.Text, nil
			}
		}
	}

	return "Я не совсем уверен, но я бы отгадал " + correctOption.Letter, nil
}

func (s *GameService) QuitGame(sessionID uint) (*models.GameSession, error) {
	session, err := s.sessionRepo.GetSessionByID(sessionID)
	if err != nil {
		return nil, err
	}

	session.Status = "quit"
	endTime := time.Now()
	session.EndTime = &endTime

	err = s.sessionRepo.UpdateSession(session)
	return session, err
}

func (s *GameService) GetLeaderboard() ([]models.GameSession, error) {
	return s.sessionRepo.GetLeaderboard()
}
