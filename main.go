package main

import (
	"DDD/pb"
	"DDD/server"
	"log"
	"time"
)

func main() {
	stdSvr := server.NewStudentServer()

	addReq := &pb.AddStudentRequest{
		Name:      "chj",
		Email:     "1021780227@qq.com",
		Phone:     "13632014571",
		BirthDate: time.Now(),
		Address: pb.AddressItem{
			Province: "gd",
			City:     "sz",
			County:   "zg",
			Street:   "xx",
		},
	}
	_, err := stdSvr.AddStudent(addReq)
	if err != nil {
		log.Println(err)
		return
	}

}
