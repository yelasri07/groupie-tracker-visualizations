package controllers

import (
	"html/template"
	"net/http"
)

func RenderTempalte(w http.ResponseWriter, url string, data any, status int) error {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles(url)
	if err != nil {
		return err
	}

	
	err = tmpl.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}
