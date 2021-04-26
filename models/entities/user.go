package entities

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id"`
	VertID    int64     `gorm:"column:vert_id"`
	CustID    int64     `gorm:"column:cust_id"`
	Name      string    `gorm:"column:user_name"`
	UserName  string    `gorm:"column:user_login"`
	Password  string    `gorm:"column:hash_passwd"`
	Token     string    `gorm:"column:token"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Active    int64     `gorm:"column:active"`
}

func (User) TableName() string {
	return "csp_user"
}
