// cmd/api/main.go

package main

import (
	"database/sql"
	"fmt"
	"github.com/MikeMwita/messaging_app.git/internal/delivery/htttp"
	"github.com/MikeMwita/messaging_app.git/internal/repository"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	repo := repository.NewRepository(db)

	// Seed the database with messages from CSV
	csvFilePath := "/home/mike/Downloads/GeneralistRails_Project_MessageData.csv"
	err = repo.SeedMessagesFromCSV(csvFilePath)
	if err != nil {
		log.Fatal("Failed to seed the database:", err)
	}

	router := gin.Default()
	htttp.SetRoutes(router, repo)

	// Start the server
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	serverAddr := fmt.Sprintf(":%d", portStr)
	fmt.Printf("Server is running on http://localhost:%d\n", portStr)
	log.Fatal(router.Run(serverAddr))
}
