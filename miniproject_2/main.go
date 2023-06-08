package main

import (
	"github.com/gin-gonic/gin"
	"miniproject_2/internal/account"
	"miniproject_2/internal/customers"
	"miniproject_2/internal/repositories"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Initialize the logging middleware
	router.Use(logging.Middleware())

	// Create the repositories
	customersRepo := repositories.NewMySQLCustomersRepository()
	adminRepo := repositories.NewMySQLAdminRepository()

	// Create the services
	customersService := customers.NewCustomersService(customersRepo)
	accountService := account.NewAccountService(adminRepo)

	// Register the routes
	cmd.RegisterRoutes(router, customersService, accountService)

	// Start the server
	router.Run(":8080")
}
