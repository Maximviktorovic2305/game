package models

type Question struct {
	ID      uint     `gorm:"primaryKey" json:"id"`
	Text    string   `json:"text"`
	Level   int      `json:"level"` // 1–15
	Prize   int      `json:"prize"` // 100, 200, ..., 1000000
	Options []Option `gorm:"foreignKey:QuestionID" json:"options"`
}
