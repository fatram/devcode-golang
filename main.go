package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatram/devcode-golang/config"
	"github.com/fatram/devcode-golang/controller/http"
)

func beforeTerminate() {
	fmt.Println("Good bye!")
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS. We then handle this by calling
// our clean up procedure and exiting the program.
func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		beforeTerminate()
		os.Exit(0)
	}()
}

// @title Golang Test Documentation
// @version 1.0
// @description This is a API server for Simple Login Online

// @contact.name Fatur Rahman
// @contact.email frfatram@gmail.com

// @host localhost
// @BasePath /
func main() {
	setupCloseHandler()
	config.ReadConfig(".env")
	http.NewHttpController().Start("", config.Configuration().Port)
}
