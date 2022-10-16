package main

import (
	"fmt"
	"mygram-api/controller"
	"mygram-api/database"
	"mygram-api/router"
	"mygram-api/service"
)

func main() {
	db, err := database.Start()
	if err != nil {
		fmt.Println("Error start db", err)
		return
	}

	s := service.New(db)
	c := controller.New(s)

	err = router.StartServer(c)
	if err != nil {
		fmt.Println("Error start server", err)
	}
}
