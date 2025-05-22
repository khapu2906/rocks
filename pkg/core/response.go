package core


import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error,omitempty"`
}

func NewResponse(status int, message string, data interface{}, err interface{}) Response {
	return Response{
		Meta: Meta{
			Status:  status,
			Message: message,
		},
		Data:  data,
		Error: err,
	}
}

func NewSuccessResponse(c *gin.Context, status int, message string, data interface{}) interface{} {
	if message == "" {
		message = http.StatusText(status)
	}
	c.JSON(status, NewResponse(status, message, data, nil))
	return nil
}

func NewErrorResponse(c *gin.Context, status int, message string, err interface{}) interface{} {
	if message == "" {
		message = http.StatusText(status)
	}
	c.JSON(status, NewResponse(status, message, nil, err))
	return nil
}
