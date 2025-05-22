package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, statusCode int, message string) {
	log.Println(err)
}
