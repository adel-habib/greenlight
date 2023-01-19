package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// js := `{"status": "available", "environment": %q, "version": %q}`
	data := envelope{
		"status":      "available",
		"system_info": map[string]string{
		"environment": app.config.env,
		"version":     version,},
	}
	err := app.writeJSON(w,http.StatusOK,data,nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w,"the server encountered an error and could not process your request!",http.StatusInternalServerError)
	}
}
