package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

func Marshal() {
	p := Person{
		Name:  "tom",
		Age:   20,
		Email: "abc@gmail.com",
	}
	b, _ := json.Marshal(p) // 将struct转为json 字节切片
	fmt.Printf("b:%v\n", string(b))
}

func Unmarshal() {
	b1 := []byte(`{"Name":"tom","Age":20,"Email":"tom@gmail.com"}`)
	var p Person
	json.Unmarshal(b1, &p) // json转为struct
	fmt.Printf("p:%v\n", p)

}

func test_encode() {
	f, _ := os.Open("a.json")
	defer f.Close()
	dec := json.NewDecoder(f)
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("v: %v\n", v)
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}

}

func main() {
	test_encode()
}
