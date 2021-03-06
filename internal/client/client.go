package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GoingFast/test6/pkg/env"
	pb "github.com/GoingFast/test6/protobuf"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type conns struct {
	crud pb.CRUDServiceClient
}

func newConns() conns {
	conn, err := grpc.Dial(env.FallbackEnv("DIAL_ADDR", ":50051"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewCRUDServiceClient(conn)
	return conns{
		crud: client,
	}
}

type service struct {
	cs conns
}

func NewService() service {
	return service{newConns()}
}

func JSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&v)
}

func (svc service) Read() http.HandlerFunc {
	type res struct {
		Err  string        `json:"err,omitempty"`
		Data []*pb.Request `json:"data,omitempty"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := svc.cs.crud.Read(r.Context(), &empty.Empty{})
		if err != nil {
			JSON(w, 500, res{Err: err.Error()})
			return
		}
		if len(resp.Data) <= 0 {
			JSON(w, http.StatusNotFound, res{Err: "couldn't find any"})
			return
		}

		JSON(w, 200, res{Data: resp.Data})
	}
}

func (svc service) Hostname() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := svc.cs.crud.Hostname(r.Context(), &empty.Empty{})
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("Server (GRPC): %s\nClient (HTTP): %s", res.Hostname, os.Getenv("HOSTNAME"))))
	}
}
