package main

import (
	"os"
	"log"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	MQTT_MASTER_DEFAULT = "mqtt-master"
)

type Message struct {
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func getMqttMaster() string {
	host := os.Getenv("MQTT_MASTER")
	if len(host) > 0 {
		return host
	}
	return MQTT_MASTER_DEFAULT
}

func main() {
	chat, err := NewChat()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.POST("/api/rooms/:name", func(c *gin.Context) {
		roomName := c.Param("name")
		msg := Message{}
		err := c.BindJSON(&msg)
		if err != nil {
			c.JSON(400,  gin.H{"msg": "invalid payload"})
			return
		}
		room := chat.GetRoom(roomName)
		err = room.Post(msg)
		if err != nil {
			c.JSON(500,  gin.H{"msg": "failed to post message"})
			return
		}
		c.JSON(200, gin.H{"msg": "OK"})
	})
	r.GET("/api/rooms/:name", func(c *gin.Context) {
		roomName := c.Param("name")
		ch := make(chan Message)
		room := chat.GetRoom(roomName)
		room.Subscribe(ch)
		defer room.Unsubscribe(ch)
		c.Stream(func(w io.Writer) bool {
			msg := <-ch
			c.SSEvent("message", msg)
			return true
		})
	})
	r.GET("/rooms/:name", func(c *gin.Context) {
		roomName := c.Param("name")
		c.HTML(http.StatusOK, "room.tmpl", gin.H{
			"title": roomName,
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.Run()
}

