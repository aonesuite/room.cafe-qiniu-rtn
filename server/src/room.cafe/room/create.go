package room

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/rtc"

	"components/config"
	"components/db"
	"components/log"

	"providers/white"

	"room.cafe/models"
)

// CreateArgs create room args
type CreateArgs struct {
	Name    string `json:"name"`    // 自定义房间名称
	Private bool   `json:"private"` // 是否为私密房间
}

// Create 创建房间
// POST	/room
func Create(c *gin.Context) {
	log := log.New(c)
	currentUser := c.MustGet("currentUser").(*models.User)
	args := CreateArgs{}

	if c.Request.Body != nil {
		if err := c.Bind(&args); err != nil {
			log.Error("bind create room args failed", err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid args", "code": "INVALID_ARGS"})
			return
		}
	}

	uuid := bson.NewObjectId().Hex()

	// 获取 RTN token
	rtcMgr := rtc.NewManager(&qbox.Mac{
		AccessKey: config.GetString("qiniu.access_key"),
		SecretKey: []byte(config.GetString("qiniu.secret_key")),
	})

	roomAccess := rtc.RoomAccess{
		AppID:      config.GetString("qiniu.rtn_appid"),
		RoomName:   uuid,
		UserID:     currentUser.RoomUserID(),
		ExpireAt:   time.Now().Unix() + 60*60*12,
		Permission: "admin",
	}

	roomToken, err := rtcMgr.GetRoomToken(roomAccess)
	if err != nil {
		log.Error("get rtn room token failed", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "create room failed", "code": "INTERNAL_SERVER_ERROR"})
		return
	}

	// 获取白板 token
	whiteClient := white.NewClient(config.GetString("herewhite.mini_token"), config.GetString("herewhite.host"))
	whiteArgs := white.ReqCreateWhite{Name: uuid, Limit: 100}
	whiteRet, err := whiteClient.CreateWhite(log, whiteArgs)
	if err != nil {
		log.Error("create white room failed", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "create room failed", "code": "INTERNAL_SERVER_ERROR"})
		return
	}

	database := db.Get(log.ReqID())
	database = database.Begin()

	room := models.Room{
		UUID:            uuid,
		Name:            args.Name,
		Private:         args.Private,
		Owner:           currentUser.ID,
		RTC:             roomAccess.RoomName,
		RTCToken:        roomToken,
		Whiteboard:      whiteRet.Room.UUID,
		WhiteboardToken: whiteRet.RoomToken,
	}

	// 创建房间
	if err := database.Create(&room).Error; err != nil {
		database.Rollback()
		log.Error("create room failed", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "create room failed", "code": "INTERNAL_SERVER_ERROR"})
		return
	}

	attendee := models.Attendee{
		UserID: currentUser.ID,
		RoomID: room.ID,
	}

	if err := database.Create(&attendee).Error; err != nil {
		database.Rollback()
		log.Error("room add attendee failed", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "create room failed", "code": "INTERNAL_SERVER_ERROR"})
		return
	}

	database.Commit()

	c.JSON(http.StatusCreated, room)

}
