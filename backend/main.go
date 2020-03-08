package main

import (
	"fmt"
	"github.com/go-pg/pg/orm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	pg "github.com/go-pg/pg"
	generated "hazuki/generated"
	user "hazuki/service"
)

func createDBtables(db *pg.DB) {
	//create table from struct generated from proto
	opts := &orm.CreateTableOptions{
		//Temp: true,
		//IfNotExists: true,
	}
	fmt.Println("Creating tables")
	err := db.CreateTable(&user.User{}, opts)
	if err != nil {
		log.Println("Error while creating table", err)
	}
	err = db.CreateTable(&user.PendingUser{}, opts)
	if err != nil {
		log.Println("Error while creating table", err)
	}
}

func connectPostgress() *pg.DB{
	options := pg.Options{
		User: "Nico",
		Addr: "localhost:5432",
		Database: "matcha",
	}
	db := pg.Connect(&options)

	return db
}


func main() {
	if os.Getenv("PASS") == "" {
		log.Fatal("Env 'PASS' not set: password needed for sending emails")
	}
	//Tmp flag chose -> move to flag
	port := ":4242"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	fmt.Println("Starting matcha server")
	lis, err := net.Listen("tcp", port) //Double check this part
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//Postgress
	db := connectPostgress()
	createDBtables(db)
	defer db.Close()

	// grpc server / register service, no Interceptors
	grpcServer := grpc.NewServer()
	generated.RegisterAccountServer(grpcServer, &user.UserService{db})
	reflection.Register(grpcServer)

	//serve server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
