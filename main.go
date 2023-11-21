package main

import (
	"fmt"
	"golang-chat-backend/router"
)

func main() {
	fmt.Println("Chat Message Server is Running")

	r := router.NewRouter()
	r.InitalizeRouter()
	r.Start("8080")
	fmt.Println("Server is running on port 8080")
}
