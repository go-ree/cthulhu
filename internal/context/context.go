package context

import (
	"context"
	"cthulhu/pkg/logger"
)

// RequestContext 请求上下文，包含所有请求相关的信息
type RequestContext struct {
	context.Context
	Logger *logger.Logger
	// 可以添加更多上下文信息，如：
	// UserID string
	// TraceID string
	// RequestID string
	// Metadata map[string]interface{}
}

// New 创建新的请求上下文
func New(ctx context.Context, log *logger.Logger) *RequestContext {
	return &RequestContext{
		Context: ctx,
		Logger:  log,
	}
}

// WithValue 添加值到上下文
func (c *RequestContext) WithValue(key, val interface{}) *RequestContext {
	return &RequestContext{
		Context: context.WithValue(c.Context, key, val),
		Logger:  c.Logger,
	}
}

// GetValue 从上下文获取值
func (c *RequestContext) GetValue(key interface{}) interface{} {
	return c.Context.Value(key)
}
