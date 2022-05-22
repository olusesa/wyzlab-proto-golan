import "google.golang.org/protobuf/proto"

func writeToFile(fname string, pb proto.Message){
	out, err := proto.Marshall(pb);

	if err != nil {
		log.Fatalln("Cant't serialize to bytes", err)
		return
	}
	if err = ioutil.WriteFile(fname, out, 0644)
	if err != nil {
	 log.Fatalln("Can't write to file", err)
	 return
	}
	fmt.Println("Data has been written successfully!!!")
}

func readDataFile(fname string, pb proto.Message){
	// in, err := b, err := ioutil.ReadFile(fname)
	// if err != nil {
	//  log.Print(err)
	// }
	// fmt.Println(string(b))
	in , err := ioutil.ReadFile(fname)
	if err != nill{
		log.Fatalln("Can;t read file", err)
		return
	}
	if err = proto.Unmarshal(in,pb); err != nil{
		log.Fatalln("Couldn't unmarshal", err)
		return
	}
}