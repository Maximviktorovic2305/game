package routes

import (
	"server/internal/handlers"
	"server/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:3002", "http://localhost:3003", "http://localhost:3004"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Initialize services
	questionService := services.NewQuestionService(db)
	gameService := services.NewGameService(db)

	// Initialize handlers
	questionHandler := handlers.NewQuestionHandler(questionService)
	gameHandler := handlers.NewGameHandler(gameService)

	// Routes
	e.POST("/api/game/start", gameHandler.StartGame)
	e.GET("/api/questions/next", questionHandler.GetNextQuestion)
	e.POST("/api/game/answer", gameHandler.AnswerQuestion)
	e.POST("/api/game/lifeline/fifty-fifty", gameHandler.UseFiftyFifty)
	e.POST("/api/game/lifeline/audience", gameHandler.UseAudience)
	e.POST("/api/game/lifeline/call", gameHandler.UseCall)
	e.POST("/api/game/quit/:sessionID", gameHandler.QuitGame)
	e.GET("/api/leaderboard", gameHandler.GetLeaderboard)

	return e
}
