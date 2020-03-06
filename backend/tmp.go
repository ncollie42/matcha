package main

import (
	generated "hazuki/generated"
)
var personDB = []*generated.User {
	{
		ID:  1,
		FirstName: "bill",
		LastName: "boby",
		Password:  "aaasdf",
		Email: "this@gmail.com",
	},
	{
		ID:  2,
		FirstName: "sam",
		LastName: "boby",
		Password:  "aaasdf",
		Email: "this@gmail.com",
	},
	{
		ID:  3,
		FirstName: "Hazu",
		LastName: "boby",
		Password:  "aaasdf",
		Email: "this@gmail.com",
	},
}
//Get users from DB
//var tmp []generated.CreateRequest
//err := db.Model(&tmp).Select()
//if err != nil {
//panic(err)
//}
//for _, Q := range tmp {
//fmt.Println(Q)
//}

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
