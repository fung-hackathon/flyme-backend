package main

import (
	"github.com/fung-hackathon/flyme-backend/app/interface/server"
)

func main() {
	s := server.NewServer()

	s.Start()
}
