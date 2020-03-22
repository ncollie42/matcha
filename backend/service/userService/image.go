package userService

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	generated "hazuki/generated"
	"strconv"
)

func (db *UserService) UpdateImage(_ context.Context, request *generated.ImageData) (*generated.Reply, error) {

	//Replace 42 with the ID from the header of the user
	//limit index 1...5?
	_ = os.Mkdir("./UserImages/user" + "42", 0755)
	ioutil.WriteFile("./UserImages/user"+"42" + "/test"+ strconv.Itoa(int(request.Index)) +".png", request.Image, 0644)

	return &generated.Reply{Message: "Image uploaded"}, nil
}

func (db *UserService) GetImages(request *generated.ImageRequest, stream generated.Account_GetImagesServer) error {
	_, err := checkSession(stream.Context(), db)
	if isError(err) {
		return  err
	}
	count := request.GetCount()
	log.Println("Trying to get images:", count)
	bb, err := ioutil.ReadFile("./UserImages/userX/pic")
	if err != nil {
		log.Println("Error with image,", err)
	}
	//for count > 0 {
	//getImages(stream, userID, count) -> goes to user's file, opens file -> send -> count-- -> loop
	stream.Send(&generated.ImageData{Image: bb})
	//log.Println("Just sent 1 image")
	//count--
	//}
	bb, err = ioutil.ReadFile("./UserImages/userX/pic2")
	if err != nil {
		log.Println("Error with image,", err)
	}
	stream.Send(&generated.ImageData{Image: bb})
	bb, err = ioutil.ReadFile("./UserImages/userX/pic3")
	if err != nil {
		log.Println("Error with image,", err)
	}
	stream.Send(&generated.ImageData{Image: bb})
	log.Println("sent both!")
	return nil
}
