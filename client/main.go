package main

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-sample/common/config"
	"grpc-sample/common/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	user1 := pb.User{
		Id:       "uid001",
		Name:     "Sarap1",
		Password: "kw8d hl12/3m,a",
		Gender:   pb.UserGender(pb.UserGender_value["MALE"]),
	}

	user2 := pb.User{
		Id:       "uid002",
		Name:     "Sarap2",
		Password: "kw8d hl12/3m,a",
		Gender:   pb.UserGender(pb.UserGender_value["FEMALE"]),
	}

	garage1 := pb.Garage{
		Id:   "gid001",
		Name: "Garap1",
		Coordinate: &pb.GarageCoordinate{
			Longitude: 45.213233,
			Latitude:  54.1312321,
		},
	}

	garage2 := pb.Garage{
		Id:   "gid002",
		Name: "Garap2",
		Coordinate: &pb.GarageCoordinate{
			Longitude: 55.213233,
			Latitude:  64.1312321,
		},
	}

	// =========== SERVICE USER =========

	user := serviceUser()

	fmt.Println("\n", "===========> user test")

	// register user1
	user.Register(context.Background(), &user1)

	// register user2
	user.Register(context.Background(), &user2)

	res1, err := user.List(context.Background(), new(emptypb.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}

	resUser1, _ := json.Marshal(res1.List)
	log.Println("Data user :\n", string(resUser1))

	// ============ SERVICE GARAGE ===========

	garage := serviceGarage()

	fmt.Println("\n", "===========> garage test A")

	garage.Add(context.Background(), &pb.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})

	garage.Add(context.Background(), &pb.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage2,
	})

	garageRes1, err := garage.List(context.Background(), &pb.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}

	resGarage1, _ := json.Marshal(garageRes1.List)
	log.Println(string(resGarage1))

	fmt.Println("\n", "===========> garage test B")

}

// RPC client Garage
func serviceGarage() pb.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewGaragesClient(conn)
}

// RPC client user
func serviceUser() pb.UsersClient {
	port := config.SERVICE_USER_PORT

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}

	return pb.NewUsersClient(conn)
}
