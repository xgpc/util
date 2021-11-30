package util

import (
	"fmt"
	"log"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	type user struct {
		Name string
		Age  uint8
	}
	var u user
	dataByte := []byte(`{"Age":2,"Name":"James"}`)
	err := JsonDecode(dataByte, &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	// Output:
	// {James 2}

}

func TestJsonEncode(t *testing.T) {
	var u = struct {
		Name string
		Age  uint8
	}{Name: "James", Age: 2}

	encode, err := JsonEncode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encode)
	fmt.Println(string(encode))
	// Output:
	// [123 34 78 97 109 101 34 58 34 74 97 109 101 115 34 44 34 65 103 101 34 58 50 125]
	// {"Name":"James","Age":2}
}
