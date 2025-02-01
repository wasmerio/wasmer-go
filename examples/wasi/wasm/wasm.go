package main

import (
	"github.com/common-nighthawk/go-figure"
)

func main() {}

//export HelloWorld
func HelloWorld() {
	myFigure := figure.NewFigure("Hello World", "", true)
	myFigure.Print()
}
