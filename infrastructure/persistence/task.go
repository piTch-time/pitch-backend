package persistence

import (
	"github.com/jinzhu/copier"
	"github.com/piTch-time/pitch-backend/domain"
	"github.com/piTch-time/pitch-backend/domain/entity"
	"github.com/piTch-time/pitch-backend/domain/repository"
	"gorm.io/gorm"
)

// TaskGorm is a db representation of entity.Task
type TaskGorm struct {
	ID     uint `gorm:"primaryKey"`
	RoomID uint `gorm:"column:room_id"`
	// Room        RoomGorm `gorm:"uniqueIndex:idx_room_task;column:room_id;constraint:OnDelete:CASCADE;"`
	CreatedBy   string   `gorm:"column:created_by"`
	Room        RoomGorm `gorm:"column:room_id;constraint:OnDelete:CASCADE;"`
	Description string   `gorm:"column:description"`
	IsDone      bool     `gorm:"column:is_done"`
	BaseGormModel
}

// TaskGorms define list of TaskGorm
type TaskGorms []TaskGorm

// TaskRepository ...
type TaskRepository struct {
	db *gorm.DB
}

// TableName define gorm table name
func (TaskGorm) TableName() string {
	return "tasks"
}

// NewTaskRepository ...
func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &TaskRepository{db: db}
}

func toTaskDTO(task *entity.Task) *TaskGorm {
	dto := new(TaskGorm)
	copier.Copy(&dto, &task)
	return dto
}

func (tg *TaskGorm) toEntity() *entity.Task {
	task := new(entity.Task)
	copier.Copy(task, &tg)
	return task
}

// Create ...
func (tr *TaskRepository) Create(task *entity.Task) (uint, error) {
	dto := toTaskDTO(task)
	if err := tr.db.Create(&dto).Error; err != nil {
		return domain.NilID, err
	}
	return dto.toEntity().ID, nil
}

// Update ...
func (tr *TaskRepository) Update(task *entity.Task) (*entity.Task, error) {
	dto := toTaskDTO(task)
	if err := tr.db.Save(&dto).Error; err != nil {
		return nil, err
	}
	return dto.toEntity(), nil
}

// GetByID ...
func (tr *TaskRepository) GetByID(id uint) (*entity.Task, error) {
	dto := TaskGorm{ID: id}
	if err := tr.db.First(&dto).Error; err != nil {
		return nil, err
	}
	return dto.toEntity(), nil

}

// GetAll ...
func (tr *TaskRepository) GetAll(roomID uint) (*entity.Tasks, error) {
	dto := TaskGorms{}
	if err := tr.db.Where("room_id = ?", roomID).Find(&dto).Error; err != nil {
		return nil, err
	}
	tasks := entity.Tasks{}
	for _, t := range dto {
		tasks = append(tasks, *t.toEntity())
	}
	return &tasks, nil
}

// GetAllByNickName ...
func (tr *TaskRepository) GetAllByNickName(roomID uint, createdBy string) (*entity.Tasks, error) {
	dto := TaskGorms{}
	if err := tr.db.Where("room_id = ?  AND created_by = ?", roomID, createdBy).Order(" created_at ").Find(&dto).Error; err != nil {
		return nil, err
	}
	tasks := entity.Tasks{}
	for _, t := range dto {
		tasks = append(tasks, *t.toEntity())
	}
	return &tasks, nil
}
