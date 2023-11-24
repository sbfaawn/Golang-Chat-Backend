package main

import (
	"fmt"
	"golang-chat-backend/api/server"
)

func main() {
	fmt.Println("Chat Message Server is Running")

	r := server.NewServer()
	r.InitalizeServer()
	r.Start("8080")
	fmt.Println("Server is running on port 8080")
}
