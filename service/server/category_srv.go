package server

import (
	"context"
	"log"
	"net"
	"os"

	"CAP_AdminPortal/models"
	pb "CAP_AdminPortal/service/protofiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server type is used to create User Detail Service
type ProductServer struct {
	DB models.Database

	// Embed the unimplemented server
	pb.UnimplementedProductCatalogueServiceServer
}

// ProductServer function - Start gRPC server
func (s *ProductServer) ProductServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("RPC_CAP_CATEGORYSRV_PORT"))
	log.Printf("RPC_Server Running: %v", &lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Init DB connection
	pgdb, err := models.GetDBConnection()
	if err != nil {
		log.Printf("failed to connect db: %v", err)
	}
	defer pgdb.Handler.Close()

	serv := grpc.NewServer()

	pb.RegisterProductCatalogueServiceServer(serv, &ProductServer{
		DB: &pgdb,
	})

	// Register reflection service on gRPC server.
	reflection.Register(serv)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// GetAllCategories implements ProductCatalogueService.GetAllCategories
func (s *ProductServer) GetAllCategories(ctx context.Context, in *pb.GetAllCategoriesRequest) (*pb.GetAllCategoriesResponse, error) {
	log.Printf("Got request for listing user details....")

	// Get all category data
	categoryData := s.DB.GetAllCategoriesData()
	log.Printf("new_cat_data: %v", categoryData)

	return &pb.GetAllCategoriesResponse{CategoryData: categoryData}, nil
}

func (s *ProductServer) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	log.Printf("Got request for listing create category....")

	// Get all category data
	categoryData := s.DB.CreateCategories(in.CatgDesc, in.CatgDesc, in.CatgImage, in.SubCatgName, in.SubCatgDesc, in.SubCatgImage)
	log.Printf("new_cat_data: %v", categoryData)

	return &pb.CreateCategoryResponse{Status: categoryData}, nil
}

func (s *ProductServer) EditCategory(ctx context.Context, in *pb.EditCategoryRequest) (*pb.EditCategoryResponse, error) {
	log.Printf("Got request for edit category....")

	// update category data
	categoryData := s.DB.EditCategory(in.CatgName, in.CatgDesc, in.CatgId)
	log.Printf("new_cat_data: %v", categoryData)

	return &pb.EditCategoryResponse{Status: categoryData}, nil
}

func (s *ProductServer) RemoveCategory(ctx context.Context, in *pb.RemoveCategoryRequest) (*pb.RemoveCategoryResponse, error) {
	log.Printf("Got request for removing category....")

	// update category data
	categoryData := s.DB.RemoveCategory(in.CatgId)
	log.Printf("new_cat_data: %v", categoryData)

	return &pb.RemoveCategoryResponse{Status: categoryData}, nil
}

func (s *ProductServer) EditSubCategory(ctx context.Context, in *pb.EditSubCategoryRequest) (*pb.EditSubCategoryResponse, error) {
	log.Printf("Got request for editting sub-category....")

	// update sub-category data
	subCategoryData := s.DB.EditSubCategory(in.SubCatgId, in.SubCatgName, in.SubCatgDesc)
	log.Printf("new_cat_data: %v", subCategoryData)

	return &pb.EditSubCategoryResponse{Status: subCategoryData}, nil
}

func (s *ProductServer) RemoveSubCategory(ctx context.Context, in *pb.RemoveSubCategoryRequest) (*pb.RemoveSubCategoryResponse, error) {
	log.Printf("Got request for removing sub-category....")

	// remove sub-category data
	subCategoryData := s.DB.RemoveSubCategory(in.SubCatgId)
	log.Printf("new_sub_cat_data: %v", subCategoryData)

	return &pb.RemoveSubCategoryResponse{Status: subCategoryData}, nil
}
