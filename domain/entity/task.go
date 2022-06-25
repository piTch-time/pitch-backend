package entity

import "time"

// Task ...
type Task struct {
	ID          uint
	RoomID      uint
	CreatedBy   string // user nickname temporary
	Description string
	IsDone      bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// Tasks ...
type Tasks []Task

// NewTask ...
func NewTask(roomID uint, createdBy, description string) (*Task, error) {
	return &Task{
		RoomID:      roomID,
		CreatedBy:   createdBy,
		Description: description,
		IsDone:      false,
	}, nil
}

// IsEqual guarantees Entity's identity
func (t *Task) IsEqual(other *Task) bool {
	return other.ID == t.ID
}
