package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DelusionalOptimist/goserver/router"
)

var (
	portFlag = flag.Int("port", 8080, "Port that the server will listen on")
	helpFlag = flag.Bool("help", false, "Show this help message")
)

var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	flag.Parse()
	if *helpFlag {
		Usage()
		os.Exit(0)
	}

	l := log.New(os.Stdout, "goserver: ", log.LstdFlags)

	router := router.NewRouter(*portFlag, l)
	err := router.Run()
	if err != nil {
		l.Fatalln(err)
	}
}
