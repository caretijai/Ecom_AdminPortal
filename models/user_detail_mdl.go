package models

import (
	"CAP_AdminPortal/models/entities"
	"fmt"
	"log"

	// postgres driver

	pb "CAP_AdminPortal/service/protofiles"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type UserDetail struct {
	Name         string `gorm:"column:user_name"`
	UserName     string `gorm:"column:user_login"`
	Mob          string `gorm:"column:cust_mobile"`
	CustName     string `gorm:"column:cust_name"`
	Email        string `gorm:"column:cust_email"`
	AddressType  string `gorm:"addrs_type"`
	AddressLabel string `gorm:"addrs_label"`
	AddressLine1 string `gorm:"addrs_line_1"`
	AddressLine2 string `gorm:"addrs_line_2"`
	AddressLine3 string `gorm:"addrs_line_3"`
	City         string `gorm:"city"`
	State        string `gorm:"state_name"`
	CountryName  string `gorm:"country_name"`
	PIN          int32  `gorm:"pin"`
	Region       string `gorm:"region"`
}

//PsqlInfo prepares the connection string for gorm
func (pgdb *PostgresDB) PsqlInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		pgdb.Host, pgdb.Port, pgdb.User, pgdb.Password, pgdb.DBname)
}

// ReadAllCSPUser fetches all the rows for a given user
func (pgdb *PostgresDB) ReadAllCSPUser() ([]*pb.UserProfile, error) {
	fmt.Println("ReadAllCSPUser database function called...")
	cspuserrec := []*pb.UserProfile{}

	errVal := pgdb.Handler.Raw("SELECT custProfile.* , customerAddress.* FROM (SELECT cusr.user_name, cusr.user_login, cmcust.cust_mobile, cmcust.cust_email, cmcust.id  FROM csp_user AS cusr INNER JOIN cust_mgmt_customer AS cmcust ON cmcust.id=cusr.cust_id WHERE cmcust.active=B'1') AS custProfile INNER JOIN (SELECT custState.*, csp_country.name AS country_name, csp_country.region FROM (SELECT custAddressLine.*, csp_state.state_name, csp_state.country_id FROM (SELECT  csp_address.customer_id ,csp_address_type.addrs_type ,csp_address.addrs_label, csp_address.addrs_line_1, csp_address.addrs_line_2, csp_address.addrs_line_3, csp_address.city, csp_address.pin, csp_address.state_id FROM csp_address inner join csp_address_type on csp_address_type.id = csp_address.addrs_type_id ORDER BY csp_address_type.addrs_type ASC) AS custAddressLine INNER JOIN csp_state ON csp_state.id=custAddressLine.state_id) AS custState INNER JOIN csp_country ON csp_country.id=custState.country_id) AS customerAddress ON customerAddress.customer_id=custProfile.id").Find(&cspuserrec).Error
	// errVal := pgdb.Handler.Where(
	// 	"id = 18").Find(&cspuserrec).Error
	if errVal != nil {
		log.Printf("cspuserrec err: %v", errVal)
		if pgdb.IsRecordNotFoundError(errVal) {
			return cspuserrec, nil
		}
		return cspuserrec, fmt.Errorf("database error %v", errVal)
	}

	return cspuserrec, nil
}

//ReadCSPUser fetches the rows based on the ID
func (pgdb *PostgresDB) ReadCSPUser(id int64) (*entities.User, error) {
	fmt.Println("ReadCSPUser database function called...")
	cspuserrec := entities.User{}

	errVal := pgdb.Handler.Where(
		"id = ?",
		id).First(&cspuserrec).Error
	if errVal != nil {
		if pgdb.IsRecordNotFoundError(errVal) {
			return nil, nil
		}
		return nil, fmt.Errorf("database error %v", errVal)
	}

	return &cspuserrec, nil
}

//ReadCustomerMgmt fetches the rows based on the ID
func (pgdb *PostgresDB) ReadCustomerMgmt(id int64) (*entities.Customer, error) {
	fmt.Println("ReadCustomerMgmt database function called...")
	custMgmtRec := entities.Customer{}
	err := pgdb.Handler.Joins("inner join csp_user on csp_user.cust_id = cust_mgmt_customer.id").Where(""+
		"csp_user.id = ?", id).First(&custMgmtRec).Error
	if err != nil {
		if pgdb.IsRecordNotFoundError(err) {
			return &custMgmtRec, nil
		}
		return &custMgmtRec, fmt.Errorf("database error: %v", err)
	}
	return &custMgmtRec, nil
}

//ReadCustomerAddress fetches the rows based on the ID
func (pgdb *PostgresDB) ReadCustomerAddress(cust_id int64) ([]*pb.Address, error) {
	fmt.Println("ReadCSPUser database function called...")
	custMgmtAddrsRec := []*pb.Address{}

	err := pgdb.Handler.Debug().Raw("SELECT custState.*, csp_country.name, csp_country.region FROM (SELECT custAddressLine.*, csp_state.state_name, csp_state.country_id FROM (SELECT csp_address_type.addrs_type, csp_address.addrs_label, csp_address.addrs_line_1, csp_address.addrs_line_2, csp_address.addrs_line_3, csp_address.city, csp_address.pin, csp_address.state_id FROM csp_address inner join csp_address_type on csp_address_type.id = csp_address.addrs_type_id WHERE (csp_address.customer_id = ?) ORDER BY csp_address_type.addrs_type ASC) AS custAddressLine INNER JOIN csp_state ON csp_state.id=custAddressLine.state_id) AS custState INNER JOIN csp_country ON csp_country.id=custState.country_id", cust_id).Scan(&custMgmtAddrsRec).Error

	if err != nil {
		if pgdb.IsRecordNotFoundError(err) {
			return custMgmtAddrsRec, nil
		}
		return custMgmtAddrsRec, fmt.Errorf("database error: %v", err)
	}
	return custMgmtAddrsRec, nil
}
