package entities

import (
	"time"
)

type Address struct {
	ID            int64     `gorm:"id"`
	CustomerID    int64     `gorm:"customer_id"`
	AddressTypeID int32     `gorm:"addrs_type_id"`
	AddressLabel  string    `gorm:"addrs_label"`
	AddressLine1  string    `gorm:"addrs_line_1"`
	AddressLine2  string    `gorm:"addrs_line_2"`
	AddressLine3  string    `gorm:"addrs_line_3"`
	City          string    `gorm:"city"`
	StateID       int32     `gorm:"state_id"`
	PIN           int32     `gorm:"pin"`
	Active        int32     `gorm:"active"`
	CreatedAt     time.Time `gorm:"created_at"`
	ModifiedAt    time.Time `gorm:"modified_at"`
	CreatedBy     int64     `gorm:"created_by"`
	ModifiedBy    int64     `gorm:"modified_by"`
}

func (Address) TableName() string {
	return "csp_address"
}
