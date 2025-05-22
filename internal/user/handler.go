package user

import (
	"net/http"

	"rocks/pkg/core"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc Usecase
}

func NewHandler(uc Usecase) *Handler {
	return &Handler{uc}
}

func (h *Handler) Get(c *gin.Context) {
	users, err := h.uc.Get()
	if err != nil {
		core.NewErrorResponse(c, http.StatusInternalServerError, "Failed to get users", err.Error())
		return
	}
	core.NewSuccessResponse(c, http.StatusOK, "Get user successfl", gin.H{"users": users})
}
