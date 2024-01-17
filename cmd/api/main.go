package main

import (
	"fmt"
	"github.com/MikeMwita/messaging_app/startup"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	router, _, err := startup.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8080"
	}

	// Starting the server
	serverAddr := fmt.Sprintf(":%s", portStr)
	fmt.Printf("Server is running on http://localhost:%s\n", portStr)
	log.Fatal(router.Run(serverAddr))
}


