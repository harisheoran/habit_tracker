package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "harisheoran:harisheoran@tcp(127.0.0.1:3306)/habitdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DATABASE ERROR", err)
	}

	db.AutoMigrate(Habit{})

	// router
	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)

	router.HandleFunc("/habits", func(w http.ResponseWriter, req *http.Request) {

		mytemplate := "ui/main.gohtml"

		template, err := template.ParseFiles(mytemplate)
		if err != nil {
			log.Println(err)
		}

		if req.Method == "POST" {

			title := req.FormValue("title")
			action := req.FormValue("action")
			time := req.FormValue("time")

			habit := Habit{
				Title:    title,
				Action:   action,
				Duration: time,
			}

			result := db.Create(&habit)

			if result.Error != nil {
				log.Fatal("ERROR saving data", result.Error)
			} else {
				log.Println("Saved to DB")
			}

		}

		template.Execute(w, "")

	})

	// start the web server
	port := ":3000"
	fmt.Println("Starting the server on port", port)
	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal("Unable to start the server")
	}
}
