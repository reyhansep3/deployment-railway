package main

import (
	"database/sql"
	"deployment-railway/controllers"
	"deployment-railway/database"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env.railway")
	if err != nil {
		panic("Error loading .env file")
	}
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("successfully connected to database")

	database.DBMigrate(db)

	router := gin.Default()
	router.POST("/bioskop", controllers.CreateBioskop(db))
	router.GET("/bioskop", controllers.GetBioskop(db))
	router.GET("/bioskop/:id", controllers.GetBioskopByID(db))
	router.PUT("/bioskop/:id", controllers.UpdateBioskop(db))
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop(db))

	router.Run(":8080")

}
