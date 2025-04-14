package provider

import (
	"cthulhu/internal/model"
	"cthulhu/pkg/logger"
)

// UserProvider 处理用户数据访问
type UserProvider struct {
	logger *logger.Logger
}

// NewUserProvider 创建用户提供者实例
func NewUserProvider(logger *logger.Logger) *UserProvider {
	return &UserProvider{
		logger: logger,
	}
}

// GetUserByID 通过ID获取用户
func (p *UserProvider) GetUserByID(id int) (*model.User, error) {
	// 这里通常会有数据库操作，这里简化为直接返回模拟数据
	p.logger.Info("Getting user by ID: ", id)
	return model.NewUser(id, "testuser", "test@example.com")
}

// GetAllUsers 获取所有用户
func (p *UserProvider) GetAllUsers() ([]*model.User, error) {
	p.logger.Info("Getting all users")
	// 模拟返回一些用户数据
	user1, err := model.NewUser(1, "user1", "user1@example.com")
	if err != nil {
		return nil, err
	}

	user2, err := model.NewUser(2, "user2", "user2@example.com")
	if err != nil {
		return nil, err
	}

	users := []*model.User{user1, user2}
	return users, nil
}
