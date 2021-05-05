package models

import (
	"fmt"
	"log"

	// postgres driver

	"CAP_AdminPortal/models/entities"
	pb "CAP_AdminPortal/service/protofiles"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Result type
type Result struct {
	ID               int32
	ProdCode         string
	ProdName         string
	ProdDesc         string
	CatgName         string
	CatgDesc         string
	CatgImage        string
	BrandName        string
	SkuID            int32 `json:"sku_id"`
	SkuImageLink     string
	SkuAvailableQty  int32
	ColourName       string
	ColorDescription string
	SizeName         string
	SizeDescription  string
	Mrp              string
	BuyPrice         string
	SellPrice        string
	Discount         string
	OfferPrice       string
}

// finalSubCatg type
type finalSubCatg struct {
	categoryData    category
	subCategoryData []SubCatgData
}

// Category type
type category struct {
	CatgName  string `json:"catg_name"`
	CatgDesc  string `json:"catg_desc"`
	CatgImage string `json:"catg_image"`
}

// SubCatgData type
type SubCatgData struct {
	ID           int64  `json:"id"`
	SubCatgName  string `json:"sub_catg_name"`
	SubCatgDesc  string `json:"sub_catg_desc"`
	SubCatgImage string `json:"sub_catg_image"`
}

// func GetAllCategoriesData fetches all categories and related sub-categories from the database
func (pgdb *PostgresDB) GetAllCategoriesData() []*pb.FinalSubCatg {
	fmt.Println("GetAllCategoriesData database function called...")

	res := []Result{}
	subCatRes := make([]*pb.SubCategory, 0)
	// finalResult := make([]finalSubCatg, 0)
	finalResult := make([]*pb.FinalSubCatg, 0)

	// Getting products by categories
	pgdb.Handler.Raw("SELECT id,catg_name,catg_desc,catg_image FROM csp_category ORDER BY catg_name ASC;").Scan(&res)

	// get all sub-catg for category
	for _, cat := range res {
		log.Printf("category_row: %v", cat)
		pgdb.Handler.Raw("SELECT id, sub_catg_name,sub_catg_desc,sub_catg_image FROM csp_sub_category WHERE catg_id=? ORDER BY sub_catg_name ASC;", cat.ID).Scan(&subCatRes)
		categoryDat := &pb.Category{Id: int64(cat.ID), CatgName: cat.CatgName, CatgDesc: cat.CatgDesc, CatgImage: cat.CatgImage}

		finalResult = append(finalResult, &pb.FinalSubCatg{CategoryData: categoryDat, SubCategoryData: subCatRes})
	}

	return finalResult
}

func (pgdb *PostgresDB) CreateCategories(CatgName string, CatgDesc string, CatgImage string, SubCatgName string, SubCatgDesc string, SubCatgImage string) bool {
	category := entities.Category{VertID: 1, CatgName: CatgName, CatgDesc: CatgDesc, CatgImage: CatgDesc, Active: 1}
	result := pgdb.Handler.Debug().Create(&category) // pass pointer of data to Create

	log.Printf("db  category rows affected: %v, ID: %v", result.RowsAffected, category.ID)
	if result.RowsAffected > 0 {
		// subCategory := entities.SubCategory{CatgID: category.ID, SubCatgName: SubCatgName, SubCatgDesc: SubCatgDesc, SubCatgImage: SubCatgImage, Active: 1}
		subCatResult := pgdb.Handler.Debug().Exec("INSERT INTO csp_sub_category(catg_id,sub_catg_name,sub_catg_desc,sub_catg_image,catg_image,active) VALUES(?,?,?,?,?,B'1')", category.ID, SubCatgName, SubCatgDesc, SubCatgImage, -1) // pass pointer of data to Create
		log.Printf("db subCategory rows affected: %v", subCatResult.RowsAffected)

		return subCatResult.RowsAffected > 0
	}
	return false
}

func (pgdb *PostgresDB) EditCategory(CatgName string, CatgDesc string, catgID int64) bool {
	log.Printf("CatgName: %v, CatgDesc: %v, catgID: %v ", CatgName, CatgDesc, catgID)
	category := entities.Category{}
	result := pgdb.Handler.Debug().Exec("UPDATE csp_category SET catg_name=?, catg_desc=? WHERE id=?", CatgName, CatgDesc, catgID)
	// result := pgdb.Handler.Debug().Model(&category).Where("id = ?", catgID).Update(CatgName, CatgDesc)

	log.Printf("db  category rows affected: %v, ID: %v", result.RowsAffected, category.ID)
	return result.RowsAffected > 0
}

func (pgdb *PostgresDB) RemoveCategory(catgID int64) bool {
	result := pgdb.Handler.Debug().Exec("DELETE FROM csp_category WHERE id=?", catgID)
	// result := pgdb.Handler.Debug().Delete(&category, catgID)

	log.Printf("db category rows affected: %v, ID: %v", result.RowsAffected, catgID)
	return result.RowsAffected > 0
}

func (pgdb *PostgresDB) EditSubCategory(SubCatgId int64, SubCatgName string, SubCatgDesc string) bool {
	result := pgdb.Handler.Debug().Exec("UPDATE csp_sub_category SET sub_catg_name=?, sub_catg_desc=? WHERE id=?", SubCatgName, SubCatgDesc, SubCatgId)

	log.Printf("db category rows affected: %v, ID: %v", result.RowsAffected, SubCatgId)
	return result.RowsAffected > 0
}

func (pgdb *PostgresDB) RemoveSubCategory(SubCatgId int64) bool {
	result := pgdb.Handler.Debug().Exec("DELETE FROM csp_sub_category WHERE id=?", SubCatgId)
	// result := pgdb.Handler.Debug().Delete(&category, catgID)

	log.Printf("db category rows affected: %v, ID: %v", result.RowsAffected, SubCatgId)
	return result.RowsAffected > 0
}
