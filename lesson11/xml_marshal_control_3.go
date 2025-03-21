package main

import (
	"encoding/xml"
	"fmt"
)

type User struct {
	XMLName xml.Name `xml:"user"`
	ID      uint32   `xml:"id,attr"`
	Name    string   `xml:"user-info>name>first"`
	Surname string   `xml:"user-info>name>second"`
}

func main() {
	user := User{
		ID:      1,
		Name:    "John",
		Surname: "Doe"}

	userAsXML, _ := xml.MarshalIndent(user, "", "    ")
	fmt.Println(string(userAsXML))
}
