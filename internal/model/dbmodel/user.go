package dbmodel

import "time"

const (
	TCUserID        = "user_id"
	TCNickname      = "nickname"
	TCUserEmail     = "email"
	TCUserPassword  = "password"
	TCLastLoginIP   = "last_login_ip"
	TCUserCreatedAt = "created_time"
	TCUserUpdatedAt = "updated_time"
)

// https://gorm.io/zh_CN/docs/models.html
type User struct {
	UserID      int64 `gorm:"PRIMARY_KEY"`
	Nickname    string
	Email       string `gorm:"UNIQUE"`
	Password    string
	LastLoginIP string
	CreatedAt   time.Time `gorm:"column:created_time"`
	UpdatedAt   time.Time `gorm:"column:updated_time"`
}

func (*User) TableName() string {
	return "user"
}
