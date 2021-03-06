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
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpected type: %v", x)
	}
}

func doFile(p proto.Message){
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readDataFile(path, message)
	fmt.Println(message)
}

func doToJSON(p proto.Message) string{
	jsonstring := toJSON(p)
	fmt.Println(jsonstring)
	return jsonstring
}

func doFromJSON(jsonstring, t reflect.Type) proto.Message{
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonstring, message)
	return message
}

func main() {
	//fmt.Println(doSimple());
	//fmt.Println(doComplex());
	//fmt.Println(doEnum());
	//fmt.Println(doMap());
	// fmt.Println("This should be an Id")
	// doOneOf(&pb.Result_Id{Id: 56})
	// fmt.Println("This should be a Message")
	// doOneOf(&pb.Result_Message{Message: "Allah is the greatest!!!"})
	//doFile(doSimple)
	//doToJSON(doSimple)
	jsonString := doToJSON(doSimple)
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple()))
	fmt.Println(jsonString)
	fmt.Println(message)
}