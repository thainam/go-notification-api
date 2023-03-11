package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"go-notification-api/cmd"
	"go-notification-api/internal/infrastructure/database"
	"go-notification-api/internal/usecase"
	"go-notification-api/pkg/sqlite"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
	Message string      `json:"message"`
}

func main() {
	DBConnection := sqlite.OpenConnection()
	defer sqlite.CloseConnection(DBConnection)

	go cmd.InitConsumer(DBConnection)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := Response{
			Message: "I'm totally fine! Let's GO!",
		}
		WriteJsonResponse(w, response)
	})

	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var response Response
		repo := database.NewNotificationRepository(DBConnection)
		useCase := usecase.ListNotifications{NotificationRepository: repo}
		notifications, err := useCase.Execute()
		if err != nil {
			WriteJsonResponseError(w, err, "Failed to retrieve notifications")
		}
		response = Response{
			Message: "Received Notifications",
			Data:    notifications,
		}
		WriteJsonResponse(w, response)
	})

	fmt.Println("Starting Server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func WriteJsonResponse(w io.Writer, response Response) {
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal()
	}
	_, err = w.Write(jsonResponse)

	if err != nil {
		log.Fatal(err)
	}
}

func WriteJsonResponseError(w io.Writer, err error, friendlyMessage string) {
	var response Response
	response.Errors = []string{friendlyMessage}
	response.Message = "Oh no, an error occurred :("

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal()
	}
	_, err = w.Write(jsonResponse)

	if err != nil {
		log.Fatal(err)
	}
}
