package main

import (
	"fmt"
	"miniproject2/modules/admin"
	"miniproject2/modules/customer"
	"miniproject2/utils/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql()

	fmt.Println("database connected..!")

	versionRoute := router.Group("/v1")

	adminHandler := admin.NewRouter(dbCrud)
	adminHandler.Handle(versionRoute)

	customerHandler := customer.NewRouter(dbCrud)
	customerHandler.Handle(versionRoute)

	errRouter := router.Run(":8080")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
