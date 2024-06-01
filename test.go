package main

import (
	"log"
	"net/http"

	"github.com/icfoxy/GoTools"
)

type DBdata[T1 any, T2 any] struct {
	DBName string
	Key    T1
	Value  T2
}

func TestAlive(w http.ResponseWriter, r *http.Request) {
	err := GoTools.RespondByJSON(w, 200, "WuYi is alive")
	if err != nil {
		log.Println("Test Alive Failed")
	}
}
