package utl

import (
	"log"
	"net/http"
)

var DefaultErrorHandler func(error, http.ResponseWriter) = func(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error: "+err.Error(), 500)
}

var Templates TemplateExecutor

// Create a new handler with a custom error handler
func NewHandler(TemplateName string, Model interface{}, ErrorHandler func(error, http.ResponseWriter)) http.HandlerFunc {
	if Templates == nil {
		log.Fatal("You must set default Templates (utl.Templates) value.")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := Templates.ExecuteTemplate(w, TemplateName, Model)
		if err != nil && ErrorHandler != nil {
			ErrorHandler(err, w)
		} else if err != nil {
			DefaultErrorHandler(err, w)
		}
	}
}

// Create a default error handler
func Handle(TemplateName string, Model interface{}) http.HandlerFunc {
	if Templates == nil {
		log.Fatal("You must set default Templates (utl.Templates) value.")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := Templates.ExecuteTemplate(w, TemplateName, Model)
		if err != nil {
			DefaultErrorHandler(err, w)
		}
	}
}
