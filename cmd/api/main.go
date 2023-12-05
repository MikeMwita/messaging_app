package main

import (
	"database/sql"
	"fmt"
	"github.com/MikeMwita/messaging_app/internal/delivery/htttp"
	"github.com/MikeMwita/messaging_app/internal/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	password := os.Getenv("DB_PASSWORD")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	if dbConnectionString == "" {
		dbConnectionString = fmt.Sprintf("postgres://mike:%s@dpg-clk5aoeg1b2c739gus30-a.oregon-postgres.render.com/mike_dy9f?sslmode=require", password)
	}

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	// Create a repository instance
	repo := repository.NewRepository(db)
	err = repo.CreateTable()
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}

	// Seed the database with messages from CSV
	csvFilePath := "GeneralistRails_Project_MessageData.csv"
	err = repo.SeedMessagesFromCSV(csvFilePath)
	if err != nil {
		log.Fatal("Failed to seed the database:", err)
	}

	router := gin.Default()

	htttp.SetRoutes(router, repo)

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}
	// Starting  the server
	serverAddr := fmt.Sprintf(":%s", portStr)
	fmt.Printf("Server is running on http://localhost:%s\n", portStr)
	log.Fatal(router.Run(serverAddr))
}
