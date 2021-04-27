package models

import (
	"fmt"
	"log"

	// postgres driver

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
	SubCatgName  string
	SubCatgDesc  string
	SubCatgImage string
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
		pgdb.Handler.Raw("SELECT sub_catg_name,sub_catg_desc,sub_catg_image FROM csp_sub_category WHERE catg_id=? ORDER BY sub_catg_name ASC;", cat.ID).Scan(&subCatRes)
		categoryDat := &pb.Category{CatgName: cat.CatgName, CatgDesc: cat.CatgDesc, CatgImage: cat.CatgImage}

		finalResult = append(finalResult, &pb.FinalSubCatg{CategoryData: categoryDat, SubCategoryData: subCatRes})
	}

	return finalResult
}

func (pgdb *PostgresDB) CreateCategories() bool {

	return true
}
