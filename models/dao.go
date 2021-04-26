// DAO Layer
package models

import (
	"CAP_AdminPortal/models/entities"
	pb "CAP_AdminPortal/service/protofiles"
)

//Database is the interface that can be implemented by any database structure.
type Database interface {
	ReadCSPUser(id int64) (*entities.User, error)
	ReadAllCSPUser() ([]*pb.UserProfile, error)
	ReadCustomerMgmt(id int64) (*entities.Customer, error)
	ReadCustomerAddress(cust_id int64) ([]*pb.Address, error)
	GetAllCategoriesData() []*pb.FinalSubCatg
}
