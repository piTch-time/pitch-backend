package repository

import "github.com/piTch-time/pitch-backend/domain/entity"

// RoomRepository ...
type RoomRepository interface {
	Create(room *entity.Room) (uint, error)
	GetByID(id uint) (*entity.Room, error)
	GetAll() (*entity.Rooms, error)
}
