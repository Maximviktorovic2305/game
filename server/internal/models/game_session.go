package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type Lifelines struct {
	FiftyFifty bool `json:"fifty_fifty"`
	Audience   bool `json:"audience"`
	Call       bool `json:"call"`
}

// Implement the Scanner and Valuer interfaces for Lifelines
func (l *Lifelines) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, l)
}

func (l Lifelines) Value() (driver.Value, error) {
	return json.Marshal(l)
}

type GameSession struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	CurrentLevel int        `json:"current_level"`
	Score        int        `json:"score"`
	Lifelines    Lifelines  `json:"lifelines" gorm:"type:jsonb"`
	Status       string     `json:"status"` // in_progress, won, lost, quit
	StartTime    time.Time  `json:"start_time"`
	EndTime      *time.Time `json:"end_time,omitempty"`
}
