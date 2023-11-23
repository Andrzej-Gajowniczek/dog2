package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func checkArgsExist() bool {
	if len(os.Args) > 1 {
		return true
	} else {
		return false
	}
}

func pipeDataExist() bool {
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		return true
	} else {
		return false
	}

}

func setConfByFlags() error {
	return nil
}

func checkForErrors(err error, s string) {
	if err != nil {
		log.Fatalln(s, err)

	}
}

func usage() {

	flag.PrintDefaults()
	os.Exit(2)

}

func parser() {
	flag.IntVar(&mode, "mode", 0, "choose from 0-5 mode of color processing algorhythm")
	flag.IntVar(&width, "width", 0, "desired width of an image at terminal console, 0 means to fit to the term")
	flag.BoolVar(&version, "version", false, "prints release version")
	flag.BoolVar(&version, "v", false, "prints release version")
	flag.BoolVar(&about, "about", false, "info message")
	flag.BoolVar(&about, "info", false, "same as -about flag")
	flag.Parse()
}

func pipeThread() {
	defer wg.Done()
	messageR = fmt.Sprintln("hejka! TU lENKA!")

}
