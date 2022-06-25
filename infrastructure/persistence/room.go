package persistence

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/piTch-time/pitch-backend/domain"
	"github.com/piTch-time/pitch-backend/domain/entity"
	"github.com/piTch-time/pitch-backend/domain/repository"
	"gorm.io/gorm"
)

// RoomGorm is a db representation of entity.Room
type RoomGorm struct {
	ID       uint      `gorm:"primaryKey"` // pk start with 1
	Name     string    `gorm:"column:name;not null"`
	Goal     string    `gorm:"column:goal;not null"`
	Password string    `gorm:"column:password;not null"`
	MusicURL string    `gorm:"column:musicUrl;not null"`
	StartAt  time.Time `gorm:"column:startAt;not null"`
	EndAt    time.Time `gorm:"column:endAt;not null"`

	BaseGormModel
}

// RoomGorms define list of RoomGorm
type RoomGorms []RoomGorm

// RoomRepository is a impl of domain/repository/roomRepository.go RoomRepository interface
type RoomRepository struct {
	db *gorm.DB
}

// TableName define gorm table name
func (RoomGorm) TableName() string {
	return "rooms"
}

// NewRoomRepository ...
// Domain layer의 RoomRepository interface를 만족시키는 repository impl.
// gorm connection을 들고 가지고 있다.
func NewRoomRepository(db *gorm.DB) repository.RoomRepository {
	return &RoomRepository{db: db}
}

func paginate(limit, offset uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(int(offset)).Limit(int(limit))
	}
}

func toDTO(room *entity.Room) *RoomGorm {
	dto := new(RoomGorm)
	copier.Copy(&dto, &room)
	return dto
}

func (rg *RoomGorm) toEntity() *entity.Room {
	room := new(entity.Room)
	copier.Copy(room, &rg)
	return room
}

// Create ...
func (rr *RoomRepository) Create(room *entity.Room) (uint, error) {
	dto := toDTO(room)
	if err := rr.db.Create(&dto).Error; err != nil {
		return domain.NilID, err
	}
	return dto.toEntity().ID, nil
}

// GetByID func find a row by entity's ID from db
func (rr *RoomRepository) GetByID(id uint) (*entity.Room, error) {
	dto := RoomGorm{ID: id}
	if err := rr.db.First(&dto).Error; err != nil {
		return nil, err
	}
	return dto.toEntity(), nil
}

// GetAll rooms
func (rr *RoomRepository) GetAll() (*entity.Rooms, error) {
	dtos := RoomGorms{}
	if err := rr.db.Order("created_at desc").Find(&dtos).Error; err != nil {
		return nil, err
	}

	rooms := entity.Rooms{}
	for _, r := range dtos {
		rooms = append(rooms, *r.toEntity())
	}
	return &rooms, nil
}

// Delete func delete a room
func (rr *RoomRepository) Delete(room *entity.Room) error {
	dto := toDTO(room)
	if err := rr.db.Delete(&dto).Error; err != nil {
		return err
	}
	return nil
}
