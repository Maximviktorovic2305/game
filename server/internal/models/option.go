package models

type Option struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	QuestionID uint   `json:"question_id"`
	Letter     string `json:"letter"` // A, B, C, D
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"`
}
