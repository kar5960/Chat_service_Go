package main

import (
	"log"
	"net/http"
	"os"

	"github.com/kar5960/chat_service/database"
	"github.com/kar5960/chat_service/handlers"
)

func main() {
	database.ConnectDB()
	database.MigrateDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/send_message", handlers.SendMessage)
	http.HandleFunc("/get_messages", handlers.GetMessages)
    http.HandleFunc("/delete_message", handlers.DeleteMessage)
	log.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
