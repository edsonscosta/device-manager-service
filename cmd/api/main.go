package main

import (
	"database/sql"
	"device-manager-service/internal/device"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

////go:embed file://db/migrations/*.sql
//var embedMigrations embed.FS

func main() {
	logger := log.New(os.Stderr, "device-manager-service-api", 0)

	logger.Println("[API] Init API")
	connStr := "host=localhost port=5432 user=device password=Device@123 dbname=device-manager-service sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Up(db, "./db/migrations"); err != nil {
		panic(err)
	}

	deviceHandler := device.NewDeviceHandler(logger, db)

	router := gin.Default()
	api := router.Group("/v1")
	{
		api.POST("/devices", deviceHandler.Create)
		api.PUT("/devices/:id", deviceHandler.Update)
		api.PATCH("/devices/:id", deviceHandler.Patch)
		api.DELETE("/devices/:id", deviceHandler.Delete)
		api.GET("/devices/:id", deviceHandler.Get)
		api.GET("/devices", deviceHandler.GetAll)
	}

	err = router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
