package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("env load failed")
		return
	}
	port := os.Getenv("Port")
	mux := http.NewServeMux()
	mux.HandleFunc("/testAlive", TestAlive)
	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: mux,
	}
	log.Println("server starts at:", port)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("server start failed", err)
	}
}
