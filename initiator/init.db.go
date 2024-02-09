package initiator

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kingkong-be/config"
	customerDomain "kingkong-be/domain/customer"
	partDomain "kingkong-be/domain/part"
	"kingkong-be/domain/price_changes_log"
	supplierDomain "kingkong-be/domain/supplier"
	"kingkong-be/domain/transaction"
	userDomain "kingkong-be/domain/user"
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

	// Auto Migrate
	err = db.AutoMigrate(&userDomain.User{}, &customerDomain.Customer{}, &partDomain.Part{}, &supplierDomain.Supplier{}, &price_changes_log.PriceChangesLog{}, &transaction.Transaction{}, &transaction.TransactionPart{})
	if err != nil {
		log.Println("failed to migrate DB : ", err)
	}

	i.db = db
}
