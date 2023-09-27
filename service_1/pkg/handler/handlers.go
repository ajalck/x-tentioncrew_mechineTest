package handler

import (
	"net/http"
	"strconv"

	"github.com/ajalck/service_1/pkg/models"
	usecase "github.com/ajalck/service_1/pkg/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service usecase.UserUseCase
}

func NewUserHandler(usecase usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	registerReq := &models.User{}
	if err := ctx.Bind(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.service.CreateUser(registerReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}
func (h *UserHandler) GetUserByID(ctx *gin.Context) {
	userId,_ :=strconv.Atoi(ctx.Query("id"))
	user, err := h.service.GetUserByID(uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, gin.H{"status": "success",
		"data": user})
}
func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	registerReq := &models.User{}
	userId,_ :=strconv.Atoi(ctx.Query("id"))
	if err := ctx.Bind(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := h.service.UpdateUser(uint(userId),registerReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	userId,_:=strconv.Atoi(ctx.Query("id"))
	err := h.service.DeleteUser(uint(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(200, gin.H{"status": "success"})
}
