package entities

import (
	"time"
)

type SubCategory struct {
	ID           int64     `gorm:"id"`
	CatgID       int64     `gorm:"catg_id"`
	SubCatgName  string    `gorm:"sub_catg_name"`
	SubCatgDesc  string    `gorm:"sub_catg_desc"`
	SubCatgImage string    `gorm:"sub_catg_image"`
	CatgImage    int32     `gorm:"catg_image"`
	Active       int32     `gorm:"active"`
	CreatedAt    time.Time `gorm:"created_at"`
	ModifiedAt   time.Time `gorm:"modified_at"`
	CreatedBy    int64     `gorm:"created_by"`
	ModifiedBy   int64     `gorm:"modified_by"`
}

func (SubCategory) TableName() string {
	return "csp_sub_category"
}
