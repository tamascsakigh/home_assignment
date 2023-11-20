package main

import (
	_ "github.com/lib/pq"
	"github.com/tamascsakigh/home_assignment/api/middleware"
	"github.com/tamascsakigh/home_assignment/api/router"
	"github.com/tamascsakigh/home_assignment/config"
	"github.com/tamascsakigh/home_assignment/database/repository"
	"github.com/tamascsakigh/home_assignment/handler"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	configuration, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	database, err := gorm.Open(postgres.Open(configuration.DBConnectionString()))
	if err != nil {
		panic(err)
	}
	db, err := database.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err = router.NewRouter(middleware.NewMiddleware(), handler.NewHandler(repository.NewRepository(database))).InitApi().Run(":8080"); err != nil {
		panic(err)
	}
}
