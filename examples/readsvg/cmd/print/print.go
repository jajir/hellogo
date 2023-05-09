package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/jajir/goprint/internal/reading"
)

/*
	func Hello(name string) string {
		message := fmt.Sprintf("Hi, %v. Welcome!", name)
		return message
	}
*/
func main() {

	flag.Parse()
	for _, s := range flag.Args() {
		readFile(s)
	}
	fmt.Println("Done.")
}

type Group struct {
	Id    string `xml:"id,attr"`
	Paths []Path `xml:"path"`
}

type Path struct {
	Id    string `xml:"id,attr"`
	D     string `xml:"d,attr"`
	Style string `xml:"style,attr"`
}

type SVG struct {
	Xmlns  string  `xml:"xmlns,attr"`
	Groups []Group `xml:"g"`
}

func readFile(filePath string) {
	fmt.Println("parsing file: " + filePath)

	// Open our xmlFile
	xmlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	svg := &SVG{}
	err = xml.Unmarshal([]byte(xmlFile), &svg)
	if err != nil {
		fmt.Println(err)
	}

	for _, group := range svg.Groups {
		for _, path := range group.Paths {
			_, err := reading.ParseD(path.D)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
