package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterModule(rg *gin.RouterGroup, db *gorm.DB) {
    repo := NewRepository(db)
    usecase := NewUsecase(repo)
    handler := NewHandler(usecase)

    users := rg.Group("/users")
    {
		users.GET("/", handler.Get)
        users.POST("/register", handler.Register)
        users.POST("/login", handler.Login)
    }
}
