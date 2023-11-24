package main

import (
	"context"
	"grpc-sample/common/config"
	"grpc-sample/common/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var localStorage *pb.GarageListByUser

func init() {
	localStorage = new(pb.GarageListByUser)
	localStorage.List = make(map[string]*pb.GarageList)
}

type GaragesServer struct {
	pb.GaragesServer
}

func (s *GaragesServer) Add(ctx context.Context, param *pb.GarageAndUserId) (*emptypb.Empty, error) {
	userId := param.UserId
	garage := param.Garage

	if _, ok := localStorage.List[userId]; !ok {
		localStorage.List[userId] = new(pb.GarageList)
		localStorage.List[userId].List = make([]*pb.Garage, 0)
	}
	localStorage.List[userId].List = append(localStorage.List[userId].List, garage)

	log.Println("Adding garage", garage.String(), "for user", userId)

	return new(emptypb.Empty), nil
}

func (s *GaragesServer) List(ctx context.Context, param *pb.GarageUserId) (*pb.GarageList, error) {
	userId := param.UserId

	return localStorage.List[userId], nil
}

func main() {
	srv := grpc.NewServer()

	pb.RegisterGaragesServer(srv, &GaragesServer{})

	log.Println("Starting RPC server at :", config.SERVICE_GARAGE_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_GARAGE_PORT, err)
	}

	log.Fatal(srv.Serve(listener))
}
