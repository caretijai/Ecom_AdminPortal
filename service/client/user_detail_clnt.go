package client

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "CAP_AdminPortal/service/protofiles"
)

//Response data
type Response struct {
	Status  string
	Message string
}

// function GetUserData
func GetUserData(w http.ResponseWriter, r *http.Request, token string) {
	// processing request before gRPC server call
	// Set up a connection to the server.
	conn, err := grpc.Dial(":"+os.Getenv("RPC_ADMIN_PORTAL_PORT"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserDetailClient(conn)

	res, err_1 := c.GetAllUserData(context.Background(), &pb.UserDataRequest{Token: token})
	header := http.StatusOK
	response := Response{}
	if err_1 != nil {
		if st, ok := status.FromError(err_1); ok {
			switch st.Code() {
			case codes.InvalidArgument:
				header = http.StatusBadRequest
			case codes.Unauthenticated:
				header = http.StatusUnauthorized
			case codes.NotFound:
				header = http.StatusNotFound
			default:
				header = http.StatusInternalServerError
			}
			response = Response{"failure", st.Message()}

			Res, err := json.Marshal(response)
			if err != nil {
				log.Printf("err value %v: ", err)
			}
			w.WriteHeader(header)
			w.Header().Set("Content-Type", "application/json")
			w.Write(Res)
		}
	} else {
		Res, err_2 := json.Marshal(res)
		if err_2 != nil {
			log.Printf("error parsing response: ", err_2)
		}
		w.WriteHeader(header)
		w.Header().Set("Content-Type", "application/json")
		w.Write(Res)
	}
}
