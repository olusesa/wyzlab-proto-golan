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

// func doEnum() *pb.Enumeration {
// 	return &pb.Enumeration{
// 		EyeColor: &pb.EyeColor_EYE_COLOR_GREEN,
// 	}
// }

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"wyzId1": {Id : 56},
			"wyzId2": {Id : 57},
			"wyzId3": {Id : 58},
		},
	}
}

func doOneOf(message interface{}) {
	Switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func main() {
	//fmt.Println(doSimple());
	//fmt.Println(doComplex());
	//fmt.Println(doEnum());
	//fmt.Println(doMap());
	fmt.Println("This should be an Id")
	doOneOf(&PB.Result_Id{Id: 56})
	fmt.Println("This should be a Message")
	doOneOf(&PB.Result_Message{Message: "Allah is the greatest!!!"})
}