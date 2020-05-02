package main

import (
	"fmt"
	"github.com/jajir/hellogo/lev3"
)

func main() {
	fmt.Printf("Ahoj\n")
	var lis lev3.Lissajous = lev3.NewLissajous(1000, 1000, 3, 5)
	lis.Prepare()
	lis.PrintInfo()
}
