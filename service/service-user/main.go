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

var localStorage *pb.UserList

func init() {
	localStorage = new(pb.UserList)
	localStorage.List = make([]*pb.User, 0)
}

// UsersServer ini akan menjadi implementasi dari generated interface pb.UsersServer (struct model)
type usersServer struct {
	pb.UnimplementedUsersServer
}

func (sv *usersServer) Register(ctx context.Context, param *pb.User) (*emptypb.Empty, error) {
	localStorage.List = append(localStorage.List, param)

	log.Println("Registering user :", param.String())

	return new(emptypb.Empty), nil
}

func (sv *usersServer) List(ctx context.Context, void *emptypb.Empty) (*pb.UserList, error) {
	return localStorage, nil
}

func main() {

	// Pembuatan objek grpc server dilakukan lewat grpc.NewServer()
	srv := grpc.NewServer()

	// RegisterUsersServer() otomatis digenerate oleh protoc, karena service Users didefinisikan.
	// Contoh lainnya misal service SomeServiceTest disiapkan, maka fungsi RegisterSomeServiceTestServer() dibuat.
	pb.RegisterUsersServer(srv, &usersServer{})

	log.Println("Starting RPC server at", config.SERVICE_USER_PORT)

	// Selanjutnya, siapkan objek listener yang listen ke port config.SERVICE_USER_PORT,
	// lalu gunakan listener tersebut sebagai argument method .Serve() milik objek rpc server.

	listerner, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen to %s: %v", config.SERVICE_USER_PORT, err)
	}

	log.Fatal(srv.Serve(listerner))
}
