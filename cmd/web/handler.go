package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("Welcome to habit Tracker"))
}

func habitsHandler(w http.ResponseWriter, req *http.Request) {
	// serve the template
	mytemplate := "ui/main.gohtml"

	template, err := template.ParseFiles(mytemplate)
	if err != nil {
		log.Println(err)
	}

	if req.Method == "POST" {

		title := req.FormValue("title")
		action := req.FormValue("action")
		duration := req.FormValue("time")

		habit := Habit{
			Title:    title,
			Action:   action,
			Duration: duration,
		}

		fmt.Println(habit)

	}

	template.Execute(w, "")
}
