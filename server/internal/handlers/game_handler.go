package handlers

import (
	"net/http"
	"strconv"

	"server/internal/models"
	"server/internal/services"

	"github.com/labstack/echo/v4"
)

type GameHandler struct {
	service *services.GameService
}

func NewGameHandler(service *services.GameService) *GameHandler {
	return &GameHandler{service: service}
}

type StartGameResponse struct {
	Session *models.GameSession `json:"session"`
}

func (h *GameHandler) StartGame(c echo.Context) error {
	session, err := h.service.StartNewGame()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to start game"})
	}

	return c.JSON(http.StatusOK, StartGameResponse{Session: session})
}

type AnswerRequest struct {
	SessionID  uint `json:"session_id"`
	QuestionID uint `json:"question_id"`
	OptionID   uint `json:"option_id"`
}

type AnswerResponse struct {
	Correct bool                `json:"correct"`
	Session *models.GameSession `json:"session"`
}

func (h *GameHandler) AnswerQuestion(c echo.Context) error {
	var req AnswerRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	correct, session, err := h.service.AnswerQuestion(req.SessionID, req.QuestionID, req.OptionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to process answer: " + err.Error()})
	}

	return c.JSON(http.StatusOK, AnswerResponse{
		Correct: correct,
		Session: session,
	})
}

type LifelineRequest struct {
	SessionID  uint `json:"session_id"`
	QuestionID uint `json:"question_id"`
}

type FiftyFiftyResponse struct {
	RemovedOptions []models.Option `json:"removed_options"`
}

func (h *GameHandler) UseFiftyFifty(c echo.Context) error {
	var req LifelineRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	removedOptions, err := h.service.UseFiftyFifty(req.SessionID, req.QuestionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to use lifeline: " + err.Error()})
	}

	return c.JSON(http.StatusOK, FiftyFiftyResponse{RemovedOptions: removedOptions})
}

type AudienceResponse struct {
	Percentages map[string]int `json:"percentages"`
}

func (h *GameHandler) UseAudience(c echo.Context) error {
	var req LifelineRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	percentages, err := h.service.UseAudience(req.SessionID, req.QuestionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to use lifeline: " + err.Error()})
	}

	return c.JSON(http.StatusOK, AudienceResponse{Percentages: percentages})
}

type CallResponse struct {
	Advice string `json:"advice"`
}

func (h *GameHandler) UseCall(c echo.Context) error {
	var req LifelineRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	advice, err := h.service.UseCall(req.SessionID, req.QuestionID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to use lifeline: " + err.Error()})
	}

	return c.JSON(http.StatusOK, CallResponse{Advice: advice})
}

func (h *GameHandler) QuitGame(c echo.Context) error {
	sessionID, err := strconv.ParseUint(c.Param("sessionID"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid session ID"})
	}

	session, err := h.service.QuitGame(uint(sessionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to quit game: " + err.Error()})
	}

	return c.JSON(http.StatusOK, session)
}

func (h *GameHandler) GetLeaderboard(c echo.Context) error {
	sessions, err := h.service.GetLeaderboard()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get leaderboard: " + err.Error()})
	}

	return c.JSON(http.StatusOK, sessions)
}
