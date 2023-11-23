package main

import (
	"fmt"
	"sync"
)

// global variables
var pth bool
var cth bool
var about bool
var version bool
var mode int
var width int
var xSize int
var ver string = "dog version: 1.02"
var messageR string
var wg sync.WaitGroup

func main() {

	//fmt.Println("start")
	//fmt.Println("args", checkArgsExist())
	//fmt.Println("pipe", pipeDataExist())

	pth = pipeDataExist()
	cth = checkArgsExist()
	/*
		if !pth && !cth {
			usage()
			//parser()
		}
	*/
	//fmt.Println("contine the prog.")
	//if checkArgsExist() {
	parser()

	if !pth && !cth {
		usage()
		//parser()
	}
	if version {
		fmt.Println(ver)
	}
	//fmt.Println("mode:", mode, "width", width) //debug info
	//}

	if pipeDataExist() {
		wg.Add(1)
		go pipeThread()
	}

	wg.Wait()
	if messageR != "" {
		fmt.Print(messageR)
	}
}
