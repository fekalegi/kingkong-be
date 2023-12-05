package initiator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sharing-vision-2021/config"
	"sharing-vision-2021/domain/post"
)

type InitiationManager interface {
	initGin()
	initDB()

	GetDB() *gorm.DB
	GetGin() *gin.Engine
}

type initiator struct {
	gin *gin.Engine
	db  *gorm.DB
}

func (i *initiator) GetDB() *gorm.DB {
	return i.db
}

func (i *initiator) GetGin() *gin.Engine {
	return i.gin
}

func NewInit() InitiationManager {
	initiation := new(initiator)
	initiation.initDB()
	initiation.initGin()
	return initiation
}

func (i *initiator) initGin() {
	i.gin = gin.Default()
}

func (i *initiator) initDB() {
	conf := config.NewDBConfig()
	mysqlConf := conf.GetMySQLConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/mysql?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Create the database if it doesn't exist
	db.Exec("CREATE DATABASE IF NOT EXISTS " + mysqlConf.Database)

	// Close the temporary connection
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close temporary connection")
	}
	sqlDB.Close()

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.Database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database : ", err)
	}
	if err = db.AutoMigrate(&post.Post{}); err != nil {
		log.Panic("Failed to migrate database : ", err)
	}

	i.db = db
}
