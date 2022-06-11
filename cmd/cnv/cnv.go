// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Allows to read SVG file and convert curves into motor moves.
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Path struct {
	D           string `xml:"d,attr"`
	Fill        string `xml:"fill,attr"`
	Stroke      string `xml:"stroke,attr"`
	StrikeWidth string `xml:"stroke-width,attr"`
}
type Rect struct {
	X           int    `xml:"x,attr"`
	Desc        int    `xml:"y,attr"`
	Width       int    `xml:"width,attr"`
	Height      int    `xml:"height,attr"`
	Fill        string `xml:"fill,attr"`
	Stroke      string `xml:"stroke,attr"`
	StrikeWidth string `xml:"stroke-width,attr"`
}

type Svg struct {
	Title string `xml:"title"`
	Desc  string `xml:"desc"`
	Rect  []Rect `xml:"rect"`
	Path  []Path `xml:"path"`
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	p := filepath.Clean(os.Args[1])
	fmt.Printf("Reading file: %s\n", p)
	exists, err := fileExists(p)
	if err != nil {
		fmt.Printf("There is a error: %s", err)
		log.Fatal(err)
	}
	if !exists {
		log.Fatal(fmt.Errorf("file '%s' doesn't exists", p))
	}

	//read XML file
	data, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Printf("There is a error: %s", err)
		log.Fatal(err)
	}

	//marshall data to struct
	note := &Svg{}
	err = xml.Unmarshal(data, &note)
	if err != nil {
		fmt.Printf("There is a error: %s", err)
		log.Fatal(err)
	}

	fmt.Printf("Title '%s'\n", note.Title)
	fmt.Println(note.Rect)
}
