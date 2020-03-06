package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	//pg "github.com/go-pb/pg"
	generated "hazuki/generated"
)

func main() {
	fmt.Println("Starting matcha server")
	lis, err := net.Listen("tcp", ":4242") //Double check this part

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	generated.RegisterPersonServer(grpcServer, &personService{})
	reflection.Register(grpcServer)


	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
	fmt.Println("DID I MAKE IT HERE?")
}

type personService struct{
	generated.UnimplementedPersonServer
}
//Return error?
func userDataisValid(req *generated.CreateRequest) bool {
	if req.GetEmail() == "" {
		return false;
	}
	if req.GetFirsName() == "" {
		return false;
	}
	if req.GetLastName() == "" {
		return false;
	}
	if req.GetPassword() == "" {
		return false;
	}
	return true;
}

func (*personService) Create(ctx context.Context, req *generated.CreateRequest) (*generated.Reply, error) {
	if !userDataisValid(req) {
		return &generated.Reply{Message: "invalid data"}, nil
	}
	email := req.GetEmail()
	firstName := req.GetFirsName()
	lastName := req.GetLastName()
	pass := req.GetPassword()
	tmp := fmt.Sprintf("Created user: %s %s\nEmail: %s\nPassword: %s", firstName, lastName, email, pass)
	fmt.Println(tmp)
	return &generated.Reply{Message: tmp}, nil
}

func (*personService) GetUsers(_ *generated.CreateRequest,stream generated.Person_GetUsersServer) error {
	for _, person := range personDB {
		stream.Send(person)
	}
	return nil
}
//TODO: validate input (no empty field)
//TODO: Hash password
//TODO: 
//
//type greeterService struct{}
//
//func (g *greeterService) SayHello(ctx context.Context, in *generated.HelloRequest) (*generated.HelloReply, error) {
//	name := in.GetName()
//	fmt.Println(name)
//	return &generated.HelloReply{Message: "Hey back to you"}, nil
//}
//
//func (g *greeterService) SayHelloAlot(req *generated.HazukiRequest, stream generated.Greeter_SayHelloAlotServer) error {
//	ticker := time.NewTicker(1 * time.Second)
//	quit := make(chan *struct{})
//	counter := 0
//	fmt.Println(stream.Context())
//	fmt.Println(("Sending streamin"))
//	go func() {
//		for {
//			select {
//			case <-ticker.C:
//				fmt.Println(("sending now"))
//				stream.Send(&generated.HelloReply{Message: "Hey hazuki! Yay"})
//				counter++
//				if counter == 20 {
//					return
//				}
//			case <-quit:
//				ticker.Stop()
//				fmt.Println("we stoped")
//				return
//			}
//		}
//	}()
//	<-quit
//	return nil
//}
