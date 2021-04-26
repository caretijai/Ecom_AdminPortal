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
type Server struct {
	DB models.Database

	// Embed the unimplemented server
	pb.UnimplementedUserDetailServer
}

// InitServer function - Start gRPC server
func (s *Server) InitServer() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("RPC_CAP_USERSRV_PORT"))
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

	pb.RegisterUserDetailServer(serv, &Server{
		DB: &pgdb,
	})

	// Register reflection service on gRPC server.
	reflection.Register(serv)
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// GetAllCategories implements ProductCatalogueService.GetAllCategories
func (s *Server) GetAllUserData(ctx context.Context, in *pb.UserDataRequest) (*pb.UserDataResponse, error) {
	log.Printf("Got request for listing user details....")

	// Get All the users
	userData, err := s.DB.ReadAllCSPUser()
	if err != nil {
		return nil, err
	}
	log.Printf("User Data: %v", userData)

	return &pb.UserDataResponse{UserData: userData}, nil
}
