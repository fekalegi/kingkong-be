package initiator

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kingkong-be/config"
	"log"
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

	dsn := fmt.Sprintf("%s&parseTime=True", mysqlConf.DSN)
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Panic("failed to connect to database", err)
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	i.db = db
}
