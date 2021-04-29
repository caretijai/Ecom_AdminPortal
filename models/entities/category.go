package entities

import (
	"time"
)

type Category struct {
	ID         int64     `gorm:"id"`
	VertID     int64     `gorm:"vert_id"`
	CatgName   string    `gorm:"catg_name"`
	CatgDesc   string    `gorm:"catg_desc"`
	CatgImage  string    `gorm:"catg_image"`
	Active     int32     `gorm:"active"`
	CreatedAt  time.Time `gorm:"created_at"`
	ModifiedAt time.Time `gorm:"modified_at"`
	CreatedBy  int64     `gorm:"created_by"`
	ModifiedBy int64     `gorm:"modified_by"`
}

func (Category) TableName() string {
	return "csp_category"
}
