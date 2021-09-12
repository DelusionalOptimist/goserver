package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DelusionalOptimist/goserver/router"
)

// commandline flags
var (
	portFlag = flag.Int("port", 8080, "Port that the server will listen on")
	helpFlag = flag.Bool("help", false, "Show this help message")
)

// help message
var Usage = func() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	// parsing commandline flags
	flag.Parse()
	if *helpFlag {
		Usage()
		os.Exit(0)
	}

	// creating a new logger
	l := log.New(os.Stdout, "goserver: ", log.LstdFlags)

	// creating a new router for the specified port
	router := router.NewRouter(*portFlag, l)

	// serving
	err := router.Run()
	if err != nil {
		l.Fatalln(err)
	}
}
