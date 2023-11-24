package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySqlOption struct {
	Address     string
	Username    string
	Password    string
	Port        string
	Database    string
	IsPopulated bool
	IsMigrate   bool
}

type mySqlConnection struct {
	address     string
	username    string
	password    string
	port        string
	database    string
	isPopulated bool
	isMigrate   bool
	DB          *gorm.DB
}

func NewMySqlConnection(option MySqlOption) *mySqlConnection {
	return &mySqlConnection{
		address:     option.Address,
		username:    option.Username,
		password:    option.Password,
		port:        option.Port,
		database:    option.Database,
		isPopulated: option.IsPopulated,
		isMigrate:   option.IsMigrate,
	}
}

func (conn *mySqlConnection) ConnectToDB() error {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dsn := conn.username + ":" + conn.password + "@tcp(localhost:3306)/" + conn.database + "?charset=utf8&parseTime=True&loc=Local"
	conn.DB, err = gorm.Open(mysql.New(
		mysql.Config{
			DSN: dsn,
		}),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		log.Fatalln("? : Could Established Connection to Databases", err)
	}

	if conn.isMigrate {
		err = conn.DB.AutoMigrate()
		fmt.Println("Error DB Migration : ", err)
		fmt.Println("Table Migration is done")
	}

	if conn.isPopulated {
		populateData(conn.DB)
	}

	return err
}

func populateData(db *gorm.DB) {
	fmt.Println("Data has been Populated!!!!")
}
