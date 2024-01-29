package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"restfule-api-generic/internal/easywalk/entity"
	"restfule-api-generic/internal/easywalk/service"
)

type simplyHandler[T entity.SimplyEntityInterface] struct {
	svc service.SimplyServiceInterface[T]
}

func (h simplyHandler[T]) Create(c *gin.Context) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := h.svc.Create(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}

func (h simplyHandler[T]) Update(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.Update(uuidID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (h simplyHandler[T]) Delete(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.Delete(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "success"})
}

func (h simplyHandler[T]) Read(c *gin.Context) {
	var id string = c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		return
	}

	uuidID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.svc.Read(uuidID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if res == nil {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h simplyHandler[T]) FindAll(c *gin.Context) {
	var req map[string]any
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entities, err := h.svc.FindAll(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if entities == nil || len(entities) == 0 {
		c.JSON(http.StatusNoContent, gin.H{"result": "no content"})
		return
	}

	c.JSON(http.StatusOK, entities)
}
