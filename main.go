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
	// password encryptor
	pe := util.NewPasswordEncryptor()

	// setup sql database
	mysqlOption := config.MySqlOption{
		Address:     "localhost",
		Username:    "root",
		Password:    "root",
		Port:        "3307",
		Database:    "chat-database",
		IsPopulated: false,
		IsMigrate:   true,
	}

	mysqlConn := config.NewMySqlConnection(mysqlOption)
	if err := mysqlConn.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	mysqlConn.MigrateData()

	// setup redis cache
	redisOption := config.RedisOption{
		Address:  "localhost",
		Port:     "6379",
		DbNum:    0,
		Password: "pass123",
	}
	redisConn := config.NewRedisConnection(redisOption)
	if err := redisConn.ConnectToRedis(); err != nil {
		log.Fatal(err)
	}

	// setup service & storage layer
	messageStorage := storage.NewMessageStorage(mysqlConn.GetDB())
	messageService := service.NewMessageService(messageStorage)

	sessionStorage := storage.NewSessionStorage(mysqlConn.GetDB(), redisConn.GetClient())
	sessionService := service.NewSessionService(sessionStorage)

	accountStorage := storage.NewAccountStorage(mysqlConn.GetDB())
	accountService := service.NewAccountService(accountStorage, pe)

	handler := handler.NewHttpHandler(accountService, sessionService, messageService)
	server := server.NewServer(handler)

	return server
}
