package main

import (
	"database/sql"
	"deployment-railway/controllers"
	"deployment-railway/database"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	envFile := "config/.env.railway"
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL tidak ditemukan di environment variable")
	}

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Gagal open database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Tidak bisa konek ke database: %v", err)
	}

	fmt.Println("Berhasil konek ke Railway PostgreSQL")

	database.DBMigrate(db)

	router := gin.Default()
	router.POST("/bioskop", controllers.CreateBioskop(db))
	router.GET("/bioskop", controllers.GetBioskop(db))
	router.GET("/bioskop/:id", controllers.GetBioskopByID(db))
	router.PUT("/bioskop/:id", controllers.UpdateBioskop(db))
	router.DELETE("/bioskop/:id", controllers.DeleteBioskop(db))

	router.Run(":8080")
}
