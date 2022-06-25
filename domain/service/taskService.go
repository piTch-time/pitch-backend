package service

import (
	"github.com/piTch-time/pitch-backend/domain"
	"github.com/piTch-time/pitch-backend/domain/entity"
	"github.com/piTch-time/pitch-backend/domain/repository"
)

// TaskService ...
type TaskService interface {
	Get(id uint) (*entity.Task, error)
	GetAll(roomID uint) (*entity.Tasks, error)
	Create(roomID uint, createdBy, description string) (uint, error)
	Update(task *entity.Task) (*entity.Task, error)
}

type taskService struct {
	taskRepository repository.TaskRepository
}

// NewTaskService ...
func NewTaskService(tr repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: tr,
	}
}

func (ts *taskService) Get(id uint) (task *entity.Task, err error) {
	if task, err = ts.taskRepository.GetByID(id); err != nil {
		return nil, err
	}
	return task, nil
}

func (ts *taskService) GetAll(roomID uint) (*entity.Tasks, error) {
	return &entity.Tasks{}, nil
}

// func (ts *taskService) GetAllByUnq(roomID uint) (*entity.Tasks, error) {
// 	return &entity.Tasks{}, nil
// }

func (ts *taskService) Create(roomID uint, createdBy, description string) (uint, error) {
	task, err := entity.NewTask(roomID, createdBy, description)
	if err != nil {
		return domain.NilID, err
	}
	taskID, err := ts.taskRepository.Create(task)
	if err != nil {
		return domain.NilID, err
	}
	return taskID, nil
}

// TODO: controller에서 GET 호출
func (ts *taskService) Update(t *entity.Task) (*entity.Task, error) {
	task, err := ts.taskRepository.Update(t)
	if err != nil {
		return nil, err
	}
	return task, nil
}
