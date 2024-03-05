package main

import (
	"fmt"
	"golang-chat-backend/api/config"
	"golang-chat-backend/api/handler"
	"golang-chat-backend/api/server"
	"golang-chat-backend/service"
	"golang-chat-backend/storage"
	"golang-chat-backend/util"
	"log"
)

func main() {
	fmt.Println("Chat Message Server is Running")

	server := initialization()
	server.InitalizeServer()
	log.Fatal(server.Start("8080"))

	fmt.Println("Server is running on port 8080")
}

func initialization() *server.Server {
	pe := util.NewPasswordEncryptor()

	mysqlOption := config.MySqlOption{
		Address:     "localhost",
		Username:    "root",
		Password:    "root",
		Port:        "3306",
		Database:    "chat",
		IsPopulated: false,
		IsMigrate:   true,
	}

	conn := config.NewMySqlConnection(mysqlOption)
	if err := conn.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	conn.MigrateData()

	messageStorage := storage.NewMessageStorage(conn.GetDB())
	messageService := service.NewMessageService(messageStorage)

	sessionStorage := storage.NewSessionStorage(conn.GetDB())
	sessionService := service.NewSessionService(sessionStorage)

	accountStorage := storage.NewAccountStorage(conn.GetDB())
	accountService := service.NewAccountService(accountStorage, pe)

	handler := handler.NewHttpHandler(accountService, sessionService, messageService)
	server := server.NewServer(handler)

	return server
}
