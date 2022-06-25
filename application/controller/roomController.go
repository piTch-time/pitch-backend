package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/piTch-time/pitch-backend/application"
	"github.com/piTch-time/pitch-backend/domain/service"
	"github.com/piTch-time/pitch-backend/infrastructure/logger"
)

// RoomController handles /v1/rooms api
type RoomController interface {
	Get() gin.HandlerFunc
	GetAll() gin.HandlerFunc
	Post() gin.HandlerFunc
	// VerifyPassword() gin.HandlerFunc
}

type roomController struct {
	roomService service.RoomService
}

// NewRoomController is a roomController's constructor
func NewRoomController(rs service.RoomService) RoomController {
	return &roomController{
		roomService: rs,
	}
}

type responseRoom struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"createdAt"`
}

type listResponseRoom struct {
	Rooms []responseRoom `json:"rooms"`
}

// @Summary      List rooms
// @Description  방 리스트
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Success      200  {object}   listResponseRoom
// @Failure      400
// @Router       /rooms [get]
func (rc *roomController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		rooms, err := rc.roomService.GetAll()
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		roomsResponse := []responseRoom{}
		for _, room := range *rooms {
			roomsResponse = append(roomsResponse, responseRoom{
				ID:        room.ID,
				Name:      room.Name,
				CreatedAt: room.CreatedAt,
			})
		}
		c.JSON(http.StatusOK, listResponseRoom{Rooms: roomsResponse})
	}
}

// TODO: task add
type detailResponseRoom struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	Goal     string     `json:"goal"`
	MusicURL string     `json:"musicUrl"`
	StartAt  *time.Time `json:"startAt"`
	EndAt    *time.Time `json:"endAt"`
}

// @Summary      get a room
// @Description  룸 상세
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        id    path     int  true  "방 ID"  Format(uint)
// @Success      200  {object}   detailResponseRoom
// @Failure      400
// @Router       /rooms/{id} [get]
func (rc *roomController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID, err := application.ParseUint(c.Param("room_id"))
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		room, err := rc.roomService.Get(roomID)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := detailResponseRoom{
			ID:       room.ID,
			Name:     room.Name,
			Goal:     room.Goal,
			MusicURL: room.MusicURL,
			StartAt:  room.StartAt,
			EndAt:    room.EndAt,
		}
		c.JSON(http.StatusOK, res)
	}
}

type postRequestRoom struct {
	Name     string `json:"name"`
	Goal     string `json:"goal"`
	Password string `json:"password"`
	MusicURL string `json:"musicUrl"`
	StartAt  string `json:"startAt"`
	EndAt    string `json:"endAt"`
}

type postResponseRoom struct {
	RoomID uint `json:"roomId"`
}

// @Summary      create a room
// @Description  방 생성
// @Tags         rooms
// @Accept       json
// @Produce      json
// @Param        room  body     postRequestRoom  true  "방 생성요청 body"
// @Success      200  {object}   postResponseRoom
// @Failure      400
// @Router       /rooms [post]
func (rc *roomController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postRequestRoom
		if err := c.BindJSON(&req); err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		roomID, err := rc.roomService.Create(req.Goal, req.Name, req.Password, req.MusicURL, req.StartAt, req.EndAt)
		if err != nil {
			logger.Error(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		res := postResponseRoom{RoomID: roomID}
		c.JSON(http.StatusOK, res)
	}
}
