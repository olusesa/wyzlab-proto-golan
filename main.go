package main

import "fmt"
import pb "github.com/olusesa/wyzlab-proto-golan/proto"

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id: 45,
		Name: "Olusesi",
		IsSimple: true,
		SampleLists: []int32{1,2,3,4,5,6,7,8,9},
	}
}

func main() {
	fmt.Println(doSimple());
}