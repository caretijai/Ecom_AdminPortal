package client

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"gopkg.in/validator.v2"

	pb "CAP_AdminPortal/service/protofiles"
)

// Payload type - Create category update
type Payload struct {
	CatgName     string `json:"catg_name" validate:"nonzero" validate:"required"`
	CatgDesc     string `json:"catg_desc" validate:"nonzero" validate:"required"`
	CatgImage    string `json:"catg_image" validate:"nonzero" validate:"required"`
	SubCatgName  string `json:"sub_catg_name" validate:"nonzero" validate:"required"`
	SubCatgDesc  string `json:"sub_catg_desc" validate:"nonzero" validate:"required"`
	SubCatgImage string `json:"sub_catg_image" validate:"nonzero" validate:"required"`
}

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
	reqPayload := &Payload{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	errVal := json.Unmarshal([]byte(reqBody), reqPayload)
	if errVal != nil {
		log.Printf("unmarshalling error: %v", errVal)
	}

	if validtrErrs := validator.Validate(reqPayload); validtrErrs != nil {
		log.Printf("payload invalid: %v", validtrErrs)
	}
	log.Printf("CatgName: %v", reqPayload.CatgName)

	// processing request before gRPC server call
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("RPC_HOST")+":"+os.Getenv("RPC_CAP_CATEGORYSRV_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductCatalogueServiceClient(conn)

	// Contact the server and print out its response.
	res, err := c.CreateCategory(context.Background(), &pb.CreateCategoryRequest{CatgName: reqPayload.CatgName, CatgDesc: reqPayload.CatgDesc, CatgImage: reqPayload.CatgImage, SubCatgName: reqPayload.SubCatgName, SubCatgDesc: reqPayload.SubCatgDesc, SubCatgImage: reqPayload.SubCatgImage})
	if err != nil {
		log.Fatalf("Could not get all categories: %v", err)
	}
	log.Printf("service responded: %s", res)

	// return service data to web-client, marshalling protobuf response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Category data created successfully."}`))
	// respJSON, _ := json.Marshal(res)
	// w.Write(respJSON)
}
