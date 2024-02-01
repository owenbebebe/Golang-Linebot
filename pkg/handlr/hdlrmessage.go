package handlr

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/owenbebebe/Golang-Linebot/pkg/model"
)

func CreateMessage(c *gin.Context) {
	var msg model.Message
	if err := c.BindJSON(&msg); err != nil {
		fmt.Println("Bind JSON error", err)
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		if err := model.CreateMessage(&msg); err != nil {
			fmt.Println("Create message error", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusCreated, msg)
		}
	}
}
