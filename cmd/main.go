package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	users "github.com/qystaubai/crm/users"
	usersRoutes "github.com/qystaubai/crm/users/controller/gin"
	usersStore "github.com/qystaubai/crm/users/repo/pq"
)

func main() {

	// create http router with gin
	router := gin.Default()

	// api layer
	api := router.Group("/api")

	// users
	us := usersStore.New()
	ua := users.NewUsersCoreImpl(us)
	ur := usersRoutes.New(ua)

	ur.AddRoutes(api)

	appServePort := 9908

	fmt.Printf("Listening and serving HTTP on %v", appServePort)

	if err := router.Run(fmt.Sprintf(":%v", appServePort)); err != nil {
		log.Fatal(err)
	}
}
