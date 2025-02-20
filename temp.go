package main

import (
	"net/http"
	"text/template"
)

func HandleHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/chatroom.html")
	if err != nil {
		http.Error(w, "Error serving html", http.StatusInternalServerError)
	}
	t.Execute(w, nil)
}
