package startup

import (
	"database/sql"
	"fmt"
	"github.com/MikeMwita/messaging_app/internal/ports"
	"github.com/MikeMwita/messaging_app/internal/repository"
	"github.com/gin-gonic/gin"
	"os"
)

//Initialize sets up the application.

func Initialize() (*gin.Engine, ports.Repository, error) {
	password := os.Getenv("DB_PASSWORD")
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")

	if dbConnectionString == "" {
		dbConnectionString = fmt.Sprintf("postgres://mike:%s@dpg-clk5aoeg1b2c739gus30-a.oregon-postgres.render.com/mike_dy9f?sslmode=require", password)
	}

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create a repository instance
	repo := repository.NewMessageRepository(db)
	err = repo.CreateTable()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create table: %v", err)
	}

	// Seed the database with messages from CSV

	csvFilePath := "GeneralistRails_Project_MessageData.csv"
	err = repo.SeedMessagesFromCSV(csvFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to seed the database: %v", err)
	}

	router := gin.Default()
	return router, repo, nil
}
