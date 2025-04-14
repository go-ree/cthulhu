package ai

import (
	"context"
	"time"

	"cthulhu/pkg/logger"
)

// Message 消息结构
type Message struct {
	Role      string    `json:"role"`      // 消息角色：user, assistant, system
	Content   string    `json:"content"`   // 消息内容
	Timestamp time.Time `json:"timestamp"` // 消息时间戳
}

// Conversation 对话历史
type Conversation struct {
	ID       string    `json:"id"`       // 会话ID
	Messages []Message `json:"messages"` // 消息历史
	Created  time.Time `json:"created"`  // 创建时间
}

// AIContext AI上下文
type AIContext struct {
	context.Context                        // 基础上下文
	Logger          *logger.Logger         // 日志记录器
	Conversation    *Conversation          // 当前对话
	UserInfo        map[string]string      // 用户信息
	Parameters      map[string]interface{} // 参数信息
	TraceID         string                 // 追踪ID
}

// New 创建新的AI上下文
func New(ctx context.Context, log *logger.Logger) *AIContext {
	return &AIContext{
		Context: ctx,
		Logger:  log,
		Conversation: &Conversation{
			ID:       generateID(),
			Messages: []Message{},
			Created:  time.Now(),
		},
		UserInfo:   make(map[string]string),
		Parameters: make(map[string]interface{}),
		TraceID:    generateID(),
	}
}

// AddMessage 添加消息到对话
func (c *AIContext) AddMessage(role, content string) {
	msg := Message{
		Role:      role,
		Content:   content,
		Timestamp: time.Now(),
	}
	c.Conversation.Messages = append(c.Conversation.Messages, msg)
}

// GetHistory 获取对话历史
func (c *AIContext) GetHistory() []Message {
	return c.Conversation.Messages
}

// SetParameter 设置参数
func (c *AIContext) SetParameter(key string, value interface{}) {
	c.Parameters[key] = value
}

// GetParameter 获取参数
func (c *AIContext) GetParameter(key string) interface{} {
	return c.Parameters[key]
}

// SetUserInfo 设置用户信息
func (c *AIContext) SetUserInfo(key, value string) {
	c.UserInfo[key] = value
}

// GetUserInfo 获取用户信息
func (c *AIContext) GetUserInfo(key string) string {
	return c.UserInfo[key]
}

// generateID 生成唯一ID
func generateID() string {
	return time.Now().Format("20060102150405") + "-" + randomString(8)
}

// randomString 生成随机字符串
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[time.Now().UnixNano()%int64(len(charset))]
		time.Sleep(1 * time.Nanosecond)
	}
	return string(result)
}
