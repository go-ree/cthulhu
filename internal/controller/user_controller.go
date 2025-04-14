package controller

import (
	"net/http"
	"strconv"

	"cthulhu/internal/provider"
	"cthulhu/pkg/logger"
	"github.com/gin-gonic/gin"
)

// UserController 处理用户相关请求
type UserController struct {
	userProvider *provider.UserProvider
	logger       *logger.Logger
}

// NewUserController 创建用户控制器实例
func NewUserController(userProvider *provider.UserProvider, logger *logger.Logger) *UserController {
	return &UserController{
		userProvider: userProvider,
		logger:       logger,
	}
}

// GetUserByID 根据ID获取用户
func (c *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		c.logger.Error("Invalid user ID: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.userProvider.GetUserByID(id)
	if err != nil {
		c.logger.Error("Failed to get user: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// GetAllUsers 获取所有用户
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userProvider.GetAllUsers()
	if err != nil {
		c.logger.Error("Failed to get users: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// RegisterRoutes 注册控制器路由
func (c *UserController) RegisterRoutes(router *gin.Engine) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/:id", c.GetUserByID)
		userGroup.GET("/", c.GetAllUsers)
	}
}
