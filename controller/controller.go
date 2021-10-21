package controller

import (
	"todo/model"
	"todo/service"
	"todo/utils"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) Create(ctx *gin.Context) {
	var createRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Username    string `json:"username"`
	}

	_, err := utils.UnmarshalInStruct(ctx.Request, &createRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	name := createRequest.Name
	description := createRequest.Description
	username := createRequest.Username
	created, err := c.service.Create(ctx, name, description, username)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{"success": created})
}

func (c *controller) GetAll(ctx *gin.Context) {
	allTodos, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"todos": allTodos})
}

func (c *controller) GetByID(ctx *gin.Context) {
	var getByIDRequest struct {
		ID string `json:"id"`
	}

	_, err := utils.UnmarshalInStruct(ctx.Request, &getByIDRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todo, err := c.service.GetByID(ctx, getByIDRequest.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"todos": todo})
}

func (c *controller) Update(ctx *gin.Context) {
	var updateRequest struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Username    string `json:"username"`
	}

	_, err := utils.UnmarshalInStruct(ctx.Request, &updateRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var updateTodo = model.Todo{
		Name:        updateRequest.Name,
		Description: updateRequest.Description,
		Username:    updateRequest.Username,
	}

	updated, err := c.service.Update(ctx, updateTodo, updateRequest.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": updated})
}

func (c *controller) Delete(ctx *gin.Context) {
	var deleteRequest struct {
		ID string `json:"id"`
	}

	_, err := utils.UnmarshalInStruct(ctx.Request, &deleteRequest)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	deleted, err := c.service.Delete(ctx, deleteRequest.ID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": deleted})
}
