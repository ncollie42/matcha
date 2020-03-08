package userService

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"time"
	"unsafe"
	generated "hazuki/generated"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var src = rand.NewSource(time.Now().UnixNano())
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)
/*
	generates a random string for pending users
*/
func randomString(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

func findUserBy(db *UserService,by, login string) (*User, error) {
	user := new(User)
	err := db.DB.Model(user).Where(by + "= ?",login).Select()
	if err != nil {
		return nil, err
	}
	return user, nil
}



func replyError(location, message string, err error) (*generated.Reply, error) {
	log.Println(location,":", message," : ", err)
	newErr := errors.New(message)
	return &generated.Reply{Message: message}, newErr
}


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func userDataisValid(req *generated.CreateRequest) bool {
	//makesure it's a valid email
	if req.GetEmail() == "" {
		return false;
	}
	if req.GetFirstName() == "" {
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