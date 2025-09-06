package database

import (
	"server/internal/models"

	"gorm.io/gorm"
)

func SeedQuestions(db *gorm.DB) error {
	// Check if questions already exist
	var count int64
	db.Model(&models.Question{}).Count(&count)
	if count > 0 {
		return nil // Questions already seeded
	}

	// Define prize amounts for each level
	prizes := []int{100, 200, 300, 500, 1000, 2000, 4000, 8000, 16000, 32000, 64000, 125000, 250000, 500000, 1000000}

	// Define questions data
	questionsData := []struct {
		Text    string
		Level   int
		Options []struct {
			Letter    string
			Text      string
			IsCorrect bool
		}
	}{
		{
			Text:  "What is the capital of France?",
			Level: 1,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "London", IsCorrect: false},
				{Letter: "B", Text: "Berlin", IsCorrect: false},
				{Letter: "C", Text: "Paris", IsCorrect: true},
				{Letter: "D", Text: "Madrid", IsCorrect: false},
			},
		},
		{
			Text:  "Which planet is known as the Red Planet?",
			Level: 2,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Venus", IsCorrect: false},
				{Letter: "B", Text: "Mars", IsCorrect: true},
				{Letter: "C", Text: "Jupiter", IsCorrect: false},
				{Letter: "D", Text: "Saturn", IsCorrect: false},
			},
		},
		{
			Text:  "What is the largest mammal in the world?",
			Level: 3,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Elephant", IsCorrect: false},
				{Letter: "B", Text: "Blue Whale", IsCorrect: true},
				{Letter: "C", Text: "Giraffe", IsCorrect: false},
				{Letter: "D", Text: "Hippopotamus", IsCorrect: false},
			},
		},
		{
			Text:  "Which element has the chemical symbol 'O'?",
			Level: 4,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Gold", IsCorrect: false},
				{Letter: "B", Text: "Oxygen", IsCorrect: true},
				{Letter: "C", Text: "Osmium", IsCorrect: false},
				{Letter: "D", Text: "Oganesson", IsCorrect: false},
			},
		},
		{
			Text:  "Who painted the Mona Lisa?",
			Level: 5,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Vincent van Gogh", IsCorrect: false},
				{Letter: "B", Text: "Pablo Picasso", IsCorrect: false},
				{Letter: "C", Text: "Leonardo da Vinci", IsCorrect: true},
				{Letter: "D", Text: "Michelangelo", IsCorrect: false},
			},
		},
		{
			Text:  "What is the smallest country in the world?",
			Level: 6,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Monaco", IsCorrect: false},
				{Letter: "B", Text: "Vatican City", IsCorrect: true},
				{Letter: "C", Text: "San Marino", IsCorrect: false},
				{Letter: "D", Text: "Liechtenstein", IsCorrect: false},
			},
		},
		{
			Text:  "Which ocean is the largest?",
			Level: 7,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Atlantic Ocean", IsCorrect: false},
				{Letter: "B", Text: "Indian Ocean", IsCorrect: false},
				{Letter: "C", Text: "Arctic Ocean", IsCorrect: false},
				{Letter: "D", Text: "Pacific Ocean", IsCorrect: true},
			},
		},
		{
			Text:  "What is the hardest natural substance on Earth?",
			Level: 8,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Gold", IsCorrect: false},
				{Letter: "B", Text: "Iron", IsCorrect: false},
				{Letter: "C", Text: "Diamond", IsCorrect: true},
				{Letter: "D", Text: "Platinum", IsCorrect: false},
			},
		},
		{
			Text:  "Which bird is known for its ability to mimic human speech?",
			Level: 9,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Eagle", IsCorrect: false},
				{Letter: "B", Text: "Parrot", IsCorrect: true},
				{Letter: "C", Text: "Owl", IsCorrect: false},
				{Letter: "D", Text: "Penguin", IsCorrect: false},
			},
		},
		{
			Text:  "What is the currency of Japan?",
			Level: 10,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Yuan", IsCorrect: false},
				{Letter: "B", Text: "Won", IsCorrect: false},
				{Letter: "C", Text: "Yen", IsCorrect: true},
				{Letter: "D", Text: "Ringgit", IsCorrect: false},
			},
		},
		{
			Text:  "Who wrote 'Romeo and Juliet'?",
			Level: 11,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Charles Dickens", IsCorrect: false},
				{Letter: "B", Text: "William Shakespeare", IsCorrect: true},
				{Letter: "C", Text: "Jane Austen", IsCorrect: false},
				{Letter: "D", Text: "Mark Twain", IsCorrect: false},
			},
		},
		{
			Text:  "What is the largest organ in the human body?",
			Level: 12,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Liver", IsCorrect: false},
				{Letter: "B", Text: "Brain", IsCorrect: false},
				{Letter: "C", Text: "Skin", IsCorrect: true},
				{Letter: "D", Text: "Heart", IsCorrect: false},
			},
		},
		{
			Text:  "Which gas makes up about 78% of Earth's atmosphere?",
			Level: 13,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Oxygen", IsCorrect: false},
				{Letter: "B", Text: "Carbon Dioxide", IsCorrect: false},
				{Letter: "C", Text: "Nitrogen", IsCorrect: true},
				{Letter: "D", Text: "Argon", IsCorrect: false},
			},
		},
		{
			Text:  "What is the tallest mammal?",
			Level: 14,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Elephant", IsCorrect: false},
				{Letter: "B", Text: "Giraffe", IsCorrect: true},
				{Letter: "C", Text: "Hippopotamus", IsCorrect: false},
				{Letter: "D", Text: "Rhinoceros", IsCorrect: false},
			},
		},
		{
			Text:  "What is the chemical symbol for gold?",
			Level: 15,
			Options: []struct {
				Letter    string
				Text      string
				IsCorrect bool
			}{
				{Letter: "A", Text: "Go", IsCorrect: false},
				{Letter: "B", Text: "Gd", IsCorrect: false},
				{Letter: "C", Text: "Au", IsCorrect: true},
				{Letter: "D", Text: "Ag", IsCorrect: false},
			},
		},
	}

	// Create questions and options
	for i, qd := range questionsData {
		question := models.Question{
			Text:  qd.Text,
			Level: qd.Level,
			Prize: prizes[i],
		}

		// Create the question first
		if err := db.Create(&question).Error; err != nil {
			return err
		}

		// Create options for this question
		for _, od := range qd.Options {
			option := models.Option{
				QuestionID: question.ID,
				Letter:     od.Letter,
				Text:       od.Text,
				IsCorrect:  od.IsCorrect,
			}
			if err := db.Create(&option).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
