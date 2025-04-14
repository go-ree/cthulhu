package ai

import (
	"net/http"

	aiContext "cthulhu/internal/context/ai"
	aiModel "cthulhu/internal/model/ai"
	"github.com/gin-gonic/gin"
)

// Protocol AI通信协议处理器
type Protocol struct {
	model aiModel.AIModel
}

// ChatRequest 聊天请求协议
type ChatRequest struct {
	Message    string                 `json:"message"`
	ModelType  string                 `json:"model_type,omitempty"`
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// ChatResponse 聊天响应协议
type ChatResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Response   string      `json:"response"`
	Model      string      `json:"model"`
	TraceID    string      `json:"trace_id"`
	Parameters interface{} `json:"parameters,omitempty"`
}

// New 创建新的协议处理器
func New(model aiModel.AIModel) *Protocol {
	return &Protocol{
		model: model,
	}
}

// Process 处理请求
func (p *Protocol) Process(ctx *aiContext.AIContext, request *ChatRequest) (*ChatResponse, error) {
	// 记录请求
	ctx.Logger.Info("Processing AI request with model: ", p.model.GetType())

	// 添加用户消息到上下文
	ctx.AddMessage("user", request.Message)

	// 处理请求
	response, err := p.model.Process(request.Message)
	if err != nil {
		return nil, err
	}

	// 添加助手消息到上下文
	ctx.AddMessage("assistant", response)

	// 返回响应
	return &ChatResponse{
		Code:       http.StatusOK,
		Message:    "success",
		Response:   response,
		Model:      string(p.model.GetType()),
		TraceID:    ctx.TraceID,
		Parameters: p.model.GetConfig(),
	}, nil
}

// HandleChat 处理聊天请求
func (p *Protocol) HandleChat(ctx *aiContext.AIContext, c *gin.Context) {
	var req ChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ChatResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request: " + err.Error(),
		})
		return
	}

	// 处理请求
	resp, err := p.Process(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ChatResponse{
			Code:    http.StatusInternalServerError,
			Message: "Processing error: " + err.Error(),
			TraceID: ctx.TraceID,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// RegisterRoutes 注册路由
func (p *Protocol) RegisterRoutes(router *gin.Engine, ctx *aiContext.AIContext) {
	aiGroup := router.Group("/api/ai")
	{
		aiGroup.POST("/chat", func(c *gin.Context) {
			p.HandleChat(ctx, c)
		})

		aiGroup.GET("/models", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"models": []string{
					string(aiModel.GPT3),
					string(aiModel.GPT4),
					string(aiModel.BERT),
				},
				"current": string(p.model.GetType()),
			})
		})
	}
}
