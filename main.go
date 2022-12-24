package main

import (
	"github.com/dhiwira25/task-5-vix-btpns-wira/database"
	"github.com/dhiwira25/task-5-vix-btpns-wira/models"
	"github.com/dhiwira25/task-5-vix-btpns-wira/router"
)

func main() {
	db := database.SetupDB()
    db.AutoMigrate(&models.User{})

    r := router.SetupRoutes(db)
    r.Run()
}