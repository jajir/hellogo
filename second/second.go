// Copyright ©2016 The ev3go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Allows to read SVG file and convert curves into motor moves.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Welcome to server 'second'\n")
		fmt.Fprintf(w, "Server is responsible form controlling mottor for plotter.\n")
		fmt.Fprintf(w, "Here is quick list of available enpoints\n")
		fmt.Fprintf(w, "/stop\n")
		fmt.Fprintf(w, "    GET - immediatelly stop current operation\n")
		fmt.Fprintf(w, "       200 - OK, no body\n")
		fmt.Fprintf(w, "/interrupt\n")
		fmt.Fprintf(w, "    GET - immediatelly stop current operation and exit program\n")
		fmt.Fprintf(w, "       200 - OK, no body\n")
		fmt.Fprintf(w, "/isReady\n")
		fmt.Fprintf(w, "    GET - get info about readiness to print\n")
		fmt.Fprintf(w, "       200 - OK, response is in body\n")
		fmt.Fprintf(w, "/calibrate\n")
		fmt.Fprintf(w, "    GET - calibre motors for axe X and for axe Y\n")
		fmt.Fprintf(w, "       200 - OK, no body\n")
		fmt.Fprintf(w, "/paint\n")
		fmt.Fprintf(w, "    POST - accept SVG XML file. response:\n")
		fmt.Fprintf(w, "       200 - OK, file will be printed.\n")
		fmt.Fprintf(w, "       503 - Can't print, I'm doing another job\n")
		//503 Service Unavailable
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
func stop(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusCreated)
		os.Exit(0)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func interrupt(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusCreated)
		os.Exit(0)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func isReady(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusCreated)
		os.Exit(0)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func calibrate(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusCreated)
		os.Exit(0)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func paint(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/paint" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "Address = %s\n", address)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/stop", stop)
	http.HandleFunc("/interrupt", interrupt)
	http.HandleFunc("/isReady", isReady)
	http.HandleFunc("/calibrate", calibrate)
	http.HandleFunc("/paint", paint)

	fmt.Printf("Starting server 'second'.\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
