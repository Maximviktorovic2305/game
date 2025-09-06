package handlers

import (
	"net/http"
	"strconv"

	"server/internal/services"

	"github.com/labstack/echo/v4"
)

type QuestionHandler struct {
	service *services.QuestionService
}

func NewQuestionHandler(service *services.QuestionService) *QuestionHandler {
	return &QuestionHandler{service: service}
}

func (h *QuestionHandler) GetNextQuestion(c echo.Context) error {
	levelStr := c.QueryParam("level")
	if levelStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Level parameter is required"})
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil || level < 1 || level > 15 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid level parameter"})
	}

	question, err := h.service.GetQuestionByLevel(level)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Question not found"})
	}

	return c.JSON(http.StatusOK, question)
}
