package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	myBook := &Book{
		Name: "Animal Farm",
		Isbn: 104,
	}

	data, err := proto.Marshal(myBook)
	if err != nil {
		log.Fatal("marshalling error: ", err)
	}

	fmt.Println(data)

	myNewBook := &Book{}
	err = proto.Unmarshal(data, myNewBook)
	if err != nil {
		log.Fatal("unmarshalling error: ", err)
	}
	fmt.Println(myNewBook.GetName())
	fmt.Println(myNewBook.GetIsbn())
}
