package entity

import (
	"time"
)

// Room ...
type Room struct {
	ID        uint
	Goal      string
	Name      string
	Password  string
	MusicURL  string
	StartAt   *time.Time
	EndAt     *time.Time
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Rooms ...
type Rooms []Room

// NewRoom ...
func NewRoom(goal, name, password, musicURL string, startAt, endAt *time.Time) (*Room, error) {
	return &Room{
		Goal:     goal,
		Name:     name,
		Password: password,
		MusicURL: musicURL,
		StartAt:  startAt,
		EndAt:    endAt,
	}, nil
}

// IsEqual guarantees Entity's identity
func (r *Room) IsEqual(other *Room) bool {
	return other.ID == r.ID
}

// IsWithinPeriod ...
func (r *Room) IsWithinPeriod(roobID uint) bool {
	return true // TODO
}
