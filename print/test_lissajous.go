package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jajir/hellogo/lev3"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("Ahoj\n")
	var lis lev3.Lissajous = lev3.NewLissajous(1000, 1000, 3, 5)
	picture := lis.GetPicture()
	save("pok.txt", picture)

	m := Message{"Alice", "Hello", 1294706395881547000, 3.14159}
	b, _ := json.Marshal(m)
	fmt.Printf("%s\n", b)

	dat, err := ioutil.ReadFile("pok.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	
	var pic lev3.Picture
	err = json.Unmarshal(dat, &pic)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	
	fmt.Printf(" %v\n", pic)
}

var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

func save(path string, lis interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(lis)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

type Message struct {
	Name string
	Body string
	Time int64
	Pok  float64 `json:"flt,string"`
}
