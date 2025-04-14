package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	aiContext "cthulhu/internal/context/ai"
	aiModel "cthulhu/internal/model/ai"
	aiProtocol "cthulhu/internal/protocol/ai"
	"cthulhu/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化日志
	log := logger.New()
	log.Info("Starting AI MCP server...")

	// 创建AI模型
	modelConfig := aiModel.ModelConfig{
		Temperature: 0.7,
		MaxTokens:   1000,
		Parameters: map[string]interface{}{
			"top_p": 0.9,
		},
	}

	model, err := aiModel.NewModel(aiModel.GPT3, modelConfig)
	if err != nil {
		log.Error("Failed to create AI model: ", err)
		os.Exit(1)
	}

	// 创建AI上下文
	ctx := aiContext.New(context.Background(), log)

	// 初始化Gin
	router := gin.Default()

	// 创建AI协议处理器
	protocol := aiProtocol.New(model)

	// 注册路由
	protocol.RegisterRoutes(router, ctx)

	// 健康检查路由
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"model":  string(model.GetType()),
		})
	})

	// 启动服务器
	go func() {
		addr := ":8080"
		log.Info("AI Service listening on ", addr)
		if err := router.Run(addr); err != nil {
			log.Error("Failed to start server: ", err)
			os.Exit(1)
		}
	}()

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down AI MCP server...")
}
