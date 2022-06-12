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

	"github.com/jajir/hellogo/internal/lev3"
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
	path := filepath.Clean(os.Args[1])
	fmt.Printf("Reading file: %s\n", path)
	exists, err := fileExists(path)
	if err != nil {
		fmt.Printf("There is a error: %s", err)
		log.Fatal(err)
	}
	if !exists {
		log.Fatal(fmt.Errorf("file '%s' doesn't exists", path))
	}

	//read XML file
	data, err := ioutil.ReadFile(path)
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

	//Paint rectangles
	for _, value := range note.Rect {
		fmt.Printf("Rect '%s'\n", value)
	}

	//Paint paths
	for _, value := range note.Path {
		fmt.Printf("Path '%s'\n", value)
	}

	plotter := lev3.NewPlotter()
	plotter.PrintInfo()

	fmt.Printf("Done")

	//A4 size 210 x 297 mm
}
