package repository

import "github.com/piTch-time/pitch-backend/domain/entity"

// TaskRepository ...
type TaskRepository interface {
	Create(task *entity.Task) (uint, error)
	Update(task *entity.Task) (*entity.Task, error)
	GetByID(id uint) (*entity.Task, error)
	GetAll(roomID uint) (*entity.Tasks, error)
	GetAllByNickName(roomID uint, createdBy string) (*entity.Tasks, error)
}
