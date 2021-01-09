package main

import (
	"dmha/hello-grpc/pb"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"time"
)

func main() {

	loc, _ := time.LoadLocation("America/Bogota")

	// ================= Protobuf Process ===================

	pbFile, _ := ioutil.ReadFile("out/person.pb")
	pbData := pb.Person{}

	pbNow := time.Now().In(loc)
	_ = proto.Unmarshal(pbFile, &pbData)
	pdDiff :=time.Now().In(loc).Sub(pbNow)

	fmt.Printf("Protobuf deserialization process takes %s\n", pdDiff.String())

	// ================= JSON Process ===================

	jsonFile, _ := ioutil.ReadFile("out/person.json")
	jsonData := pb.Person{}

	jsonNow := time.Now().In(loc)
	_ = json.Unmarshal(jsonFile, &jsonData)
	jsonDiff :=time.Now().In(loc).Sub(jsonNow)

	fmt.Printf("JSON deserialization process takes %s\n", jsonDiff.String())

	// ================= XML Process ===================

	xmlFile, _ := ioutil.ReadFile("out/person.xml")
	xmlData := pb.Person{}

	xmlNow := time.Now().In(loc)
	_ = xml.Unmarshal(xmlFile, &xmlData)
	xmlDiff :=time.Now().In(loc).Sub(xmlNow)

	fmt.Printf("XML deserialization process takes %s\n", xmlDiff.String())

}
