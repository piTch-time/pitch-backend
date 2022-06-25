package service

import (
	"github.com/piTch-time/pitch-backend/domain"
	"github.com/piTch-time/pitch-backend/domain/entity"
	"github.com/piTch-time/pitch-backend/domain/repository"
)

// RoomService ...
type RoomService interface {
	Create(goal, name, password, musicURL string, start, end string) (uint, error)
	Get(id uint) (*entity.Room, error)
	GetAll() (*entity.Rooms, error)
	Delete(room *entity.Room) error
}

type roomService struct {
	roomRepository repository.RoomRepository
}

// NewRoomService ...
func NewRoomService(rr repository.RoomRepository) RoomService {
	return &roomService{
		roomRepository: rr,
	}
}

func (rs *roomService) Create(goal, name, password, musicURL string, start, end string) (uint, error) {
	startAt := domain.StrToTime(start)
	endAt := domain.StrToTime(end)
	room, err := entity.NewRoom(goal, name, password, musicURL, &startAt, &endAt)
	if err != nil {
		return domain.NilID, err
	}
	roomID, err := rs.roomRepository.Create(room)
	if err != nil {
		return domain.NilID, err
	}
	return roomID, nil
}

func (rs *roomService) Get(id uint) (room *entity.Room, err error) {
	if room, err = rs.roomRepository.GetByID(id); err != nil {
		return nil, err
	}
	// TODO: populate tasks
	return room, nil
}

func (rs *roomService) GetAll() (*entity.Rooms, error) {
	rooms, err := rs.roomRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (rs *roomService) Delete(room *entity.Room) error {
	err := rs.roomRepository.Delete(room)
	if err != nil {
		return err
	}
	return nil
}
