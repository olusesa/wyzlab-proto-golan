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

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy {Id: 45, Name: "Olusesi"},
		MultipleDummies: []*pb.Dummy{
			{Id: 46, Name: "Lateef"},
			{Id: 47, Name: "Rahman"},
			{Id: 48, Name: "Basit"},
			{Id: 49, Name: "Rasheed"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: &pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"wyzId1" : {id : 56},
			"wyzId2" : {id : 57},
			"wyzId3" : {id : 58}
		}
	}
}

func main() {
	//fmt.Println(doSimple());
	//fmt.Println(doComplex());
	//fmt.Println(doEnum());
	fmt.Println(doMap());
}