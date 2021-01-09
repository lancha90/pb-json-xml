package main

import (
	"dmha/hello-grpc/pb"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("America/Bogota")

	person := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	// ================= JSON Process ===================

	jsonNow := time.Now().In(loc)
	jsonData, _ := json.Marshal(&person)
	jsonDiff :=time.Now().In(loc).Sub(jsonNow)

	if err := ioutil.WriteFile("out/person.json", jsonData, 0644); err != nil {
		log.Fatalln("Failed to write JSON Person:", err)
	}

	fmt.Printf("JSON serialization process takes %s\n", jsonDiff.String())

	// ================= XML Process ===================

	xmlNow := time.Now().In(loc)
	xmlData, _ := xml.Marshal(&person)
	xmlDiff :=time.Now().In(loc).Sub(xmlNow)

	if err := ioutil.WriteFile("out/person.xml", xmlData, 0644); err != nil {
		log.Fatalln("Failed to write XML Person:", err)
	}


	fmt.Printf("XML serialization process takes %s\n", xmlDiff.String())

	// ================= Protobuf Process ===================

	pbNow := time.Now().In(loc)
	pbData, _ := proto.Marshal(&person)
	pdDiff :=time.Now().In(loc).Sub(pbNow)

	if err := ioutil.WriteFile("out/person.pb", pbData, 0644); err != nil {
		log.Fatalln("Failed to write PB Person:", err)
	}

	fmt.Printf("Protobuf serialization process takes %s\n", pdDiff.String())

}
