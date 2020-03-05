package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"go-teachers-api/driver"
	"go-teachers-api/models"
)

func main() {
	db := driver.ConnectionDB()
	defer db.Close()

	models.InitTable(db)

	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Starting web server at port: 9096")
	err := http.ListenAndServe(": 9096", router)
	if err != nil {
		logrus.Fatal(err)
	}
}
