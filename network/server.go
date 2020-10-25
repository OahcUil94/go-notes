package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
)

type web1Handler struct {}
func (web1Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web1"))
}

type web2Handler struct {}
func (web2Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("web2"))
}

func main() {
	c := make(chan os.Signal)
	go func () {
		http.ListenAndServe(":9091", web1Handler{})
	}()

	go func () {
		http.ListenAndServe(":9092", web2Handler{})
	}()

	signal.Notify(c, os.Interrupt)
	s := <-c
	log.Println(s)
}
