package main


func toJSON(pb proto.Message) string {

	out, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
	return ""
	}
	return string(out)
}

func fromJSON(in string, pb proto.Message){
	if err != protojson.Unmarshal([]bytes(in), pb){
		log.Fatalln("Can't unmarshall fromJSON", err)
	}

}