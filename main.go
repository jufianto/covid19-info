package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	ACCOUNT_ID     string = ""
	TOKEN          string = ""
	SERVICE_NUMBER string = ""
)

func main() {

	fmt.Println("Start The Application")

	ACCOUNT_ID = os.Getenv("ACCOUNTID")
	TOKEN = os.Getenv("TOKEN")
	SERVICE_NUMBER = os.Getenv("SERVICE_NUMBER")

	if ACCOUNT_ID == "" || TOKEN == "" || SERVICE_NUMBER == "" {
		log.Fatal("ENV Not set, Please set your ENV", ACCOUNT_ID, TOKEN, SERVICE_NUMBER)
	}

	log.Println(ACCOUNT_ID, TOKEN, SERVICE_NUMBER)

	http.HandleFunc("/", handlerTwilioPost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	log.Println("Listen and Serve on Port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
