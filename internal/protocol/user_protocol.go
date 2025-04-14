package protocol

import (
	"errors"
	"net/http"
	"strconv"

	"cthulhu/internal/context"
	"cthulhu/internal/model"
	"github.com/gin-gonic/gin"
)

// UserProtocol 用户协议处理器
type UserProtocol struct{}

// UserRequest 用户请求协议
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserResponse 用户响应协议
type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// New 创建新的协议处理器
func New() *UserProtocol {
	return &UserProtocol{}
}

// DecodeUserRequest 解码用户请求
func (p *UserProtocol) DecodeUserRequest(c *gin.Context) (*UserRequest, error) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

// EncodeResponse 编码响应
func (p *UserProtocol) EncodeResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, UserResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// GetUserID 从请求中获取用户ID
func (p *UserProtocol) GetUserID(c *gin.Context) (int, error) {
	id := c.Param("id")
	if id == "" {
		return 0, errors.New("missing user id")
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid user id")
	}

	return userID, nil
}

// HandleGetUser 处理获取用户请求
func (p *UserProtocol) HandleGetUser(ctx *context.RequestContext, c *gin.Context) {
	userID, err := p.GetUserID(c)
	if err != nil {
		p.EncodeResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	user, err := model.NewUser(userID, "testuser", "test@example.com")
	if err != nil {
		p.EncodeResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	p.EncodeResponse(c, http.StatusOK, "success", user)
}

// RegisterRoutes 注册路由
func (p *UserProtocol) RegisterRoutes(router *gin.Engine, ctx *context.RequestContext) {
	userGroup := router.Group("/api/users")
	{
		userGroup.GET("/:id", func(c *gin.Context) {
			p.HandleGetUser(ctx, c)
		})
	}
}
