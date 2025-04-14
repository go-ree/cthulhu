package model

import (
	"errors"
	"regexp"
)

// User 用户模型
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// Validate 验证用户数据
func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username cannot be empty")
	}

	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

// isValidEmail 验证邮箱格式
func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// NewUser 创建新用户
func NewUser(id int, username, email string) (*User, error) {
	user := &User{
		ID:       id,
		Username: username,
		Email:    email,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

// Business Rules
func (u *User) IsAdmin() bool {
	// 示例业务规则：ID为1的用户为管理员
	return u.ID == 1
}

func (u *User) CanAccessResource(resourceID int) bool {
	// 示例业务规则：管理员可以访问所有资源，普通用户只能访问自己的资源
	if u.IsAdmin() {
		return true
	}
	return resourceID == u.ID
}
