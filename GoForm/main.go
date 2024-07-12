package main

import (
	"fmt"
	template2 "html/template"
	"net/http"
)

type ContactDetails struct {
	Name    string
	Email   string
	Message string
}

func main() {
	template := template2.Must(template2.ParseFiles("GoForm/assets/form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			template.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Name:    r.FormValue("Name"),
			Email:   r.FormValue("email"),
			Message: r.FormValue("message"),
		}

		_ = details
		fmt.Println("Message is: ", details)

		template.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8085", nil)
}
