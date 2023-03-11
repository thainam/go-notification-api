package logs

import (
	"log"
	"os"
)

func DbLogError(dbError error) {
	file, err := os.OpenFile("./db-errors.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	errorLog := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Println(dbError)
}
