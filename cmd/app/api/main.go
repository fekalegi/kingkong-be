package main

import (
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"kingkong-be/delivery/http/customer"
	"kingkong-be/delivery/http/part"
	"kingkong-be/delivery/http/supplier"
	"kingkong-be/delivery/http/user"
	"kingkong-be/initiator"
	"os"
	"path/filepath"
	"strings"

	customerDomain "kingkong-be/domain/customer"
	partDomain "kingkong-be/domain/part"
	supplierDomain "kingkong-be/domain/supplier"
	userDomain "kingkong-be/domain/user"
)

func main() {
	LoadEnvVars()
	i := initiator.NewInit()

	r := i.GetGin()
	db := i.GetDB()
	api := r.Group("/api")
	// CORS middleware setup allowing all origins
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))

	customerRepo := customerDomain.NewCustomerRepository(db)
	newCustomerService := customerDomain.NewCustomerImplementation(customerRepo)
	customerController := customer.NewCustomerController(newCustomerService)
	customerController.Route(api)

	partRepo := partDomain.NewPartRepository(db)
	newPartService := partDomain.NewPartImplementation(partRepo)
	partController := part.NewPartController(newPartService)
	partController.Route(api)

	supplierRepo := supplierDomain.NewSupplierRepository(db)
	newSupplierService := supplierDomain.NewSupplierImplementation(supplierRepo)
	supplierController := supplier.NewSupplierController(newSupplierService)
	supplierController.Route(api)

	userRepo := userDomain.NewUserRepository(db)
	newUserService := userDomain.NewUserImplementation(userRepo)
	userController := user.NewUserController(newUserService)
	userController.Route(api)

	r.Run("localhost:7000")
}

func LoadEnvVars() {
	cwd, _ := os.Getwd()
	dirString := strings.Split(cwd, "kingkong-be")
	dir := strings.Join([]string{dirString[0], "kingkong-be"}, "")
	AppPath := dir

	godotenv.Load(filepath.Join(AppPath, "/.env"))
}
