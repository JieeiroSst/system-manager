package main

import (
	"fmt"
	"keikibook/controller"
	"keikibook/router"
	"log"
	"net/http"
	"time"
)

func main() {
	r:=router.InitRouter()
	//views file html
	r.Handle("/",http.FileServer(http.Dir("./views/")))
	//static css images
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", indexHandler)
	server:=http.Server{
		Addr:              "localhost:8080",
		Handler:           http.HandlerFunc(controller.Lb),
		TLSConfig:         nil,
		ReadTimeout:       time.Second*10,
		ReadHeaderTimeout: time.Second*10,
		WriteTimeout:      time.Second*10,
		IdleTimeout:       time.Second*10,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}
	err := server.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}