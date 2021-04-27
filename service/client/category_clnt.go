package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"

	pb "CAP_AdminPortal/service/protofiles"
)

func GetCategoryData(w http.ResponseWriter, r *http.Request) {
	log.Println("GetAllCategoriesSrv started...")

	// processing request before gRPC server call
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("RPC_HOST")+":"+os.Getenv("RPC_CAP_CATEGORYSRV_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductCatalogueServiceClient(conn)

	// Contact the server and print out its response.
	res, err := c.GetAllCategories(context.Background(), &pb.GetAllCategoriesRequest{Id: 1})
	if err != nil {
		log.Fatalf("Could not get all categories: %v", err)
	}
	// log.Printf("service responded: %s", res)

	// return service data to web-client, marshalling protobuf response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(res)
	w.Write(respJSON)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateCategorySrv started...")

	// processing request before gRPC server call
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("RPC_HOST")+":"+os.Getenv("RPC_CAP_CATEGORYSRV_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductCatalogueServiceClient(conn)

	// Contact the server and print out its response.
	res, err := c.CreateCategory(context.Background(), &pb.CreateCategoryRequest{CatgName: "Test1", CatgDesc: "Test category description", CatgImage: "", SubCatgName: "Sub Category name", SubCatgDesc: "Sub Category description", SubCatgImage: "Sub Category image"})
	if err != nil {
		log.Fatalf("Could not get all categories: %v", err)
	}
	// log.Printf("service responded: %s", res)

	// return service data to web-client, marshalling protobuf response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJSON, _ := json.Marshal(res)
	w.Write(respJSON)
}
