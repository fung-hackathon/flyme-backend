package main

import (
	server "flyme-backend/app/interfaces"
)

func main() {
	s := server.NewServer()

	s.StartServer()
}
