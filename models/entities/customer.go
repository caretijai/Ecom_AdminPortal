package entities

import (
	"time"
)

type Customer struct {
	ID        int64     `gorm:"column:id"`
	Mob       string    `gorm:"column:cust_mobile"`
	CustKey   string    `gorm:"column:cust_key"`
	Name      string    `gorm:"column:cust_name"`
	Email     string    `gorm:"column:cust_email"`
	CreatedBy int64     `gorm:"column:created_by"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Active    int64     `gorm:"column:active"`
}

//TableName fetches the database table name.
func (custMgmtRec *Customer) TableName() string {
	return "cust_mgmt_customer"
}
