package handler

import (
	"github.com/gin-gonic/gin"
	"restfule-api-generic/internal/easywalk/entity"
	"restfule-api-generic/internal/easywalk/service"
)

type GenericHandlerInterface[T any] interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)

	Read(c *gin.Context)
	FindAll(c *gin.Context)
}

func NewHandler[T entity.SimplyEntityInterface](group *gin.RouterGroup, svc service.SimplyServiceInterface[T]) GenericHandlerInterface[T] {
	handlers := &simplyHandler[T]{svc: svc}

	group.POST("", handlers.Create)
	group.PATCH(":id", handlers.Update)
	group.DELETE(":id", handlers.Delete)

	group.GET(":id", handlers.Read)
	group.GET("", handlers.FindAll)

	return handlers
}
