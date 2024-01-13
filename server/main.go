package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.bug.st/serial"
)

var (
	port       = flag.Int("port", 8080, "The port to run the server on")
	devicePath = flag.String("device-path", "/dev/outlet", "The path to the outlet device")

	deviceMode = serial.Mode{
		BaudRate: 9600,
	}
)

func main() {
	flag.Parse()

	if *devicePath == "" {
		log.Fatal("You must provide a device with --device-path")
	}

	if _, err := os.Stat(*devicePath); err != nil {
		log.Fatalf("The device %s is not accessible or does not exist", *devicePath)
	}

	dev, err := openOutletDevice()
	if err != nil {
		log.Fatalf("error opening the outlet device: %s", err)
	}

	signalC := make(chan os.Signal, 1)
	signal.Notify(signalC, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalC
		dev.Close()
	}()

	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		dev.Write([]byte("close\n"))
	})

	http.HandleFunc("/open", func(w http.ResponseWriter, r *http.Request) {
		dev.Write([]byte("open\n"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func openOutletDevice() (serial.Port, error) {
	return serial.Open(*devicePath, &deviceMode)
}
