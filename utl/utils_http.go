package utl

import (
	"log"
	"net/http"
)

var DefaultErrorHandler func(error, http.ResponseWriter) = func(err error, w http.ResponseWriter) {
	log.Println(err.Error())
	http.Error(w, "Internal Server Error: "+err.Error(), 500)
}

func NewHandler(p HandlerParams) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := p.Template.ExecuteTemplate(w, p.TemplateName, p.Model)
		if err != nil && p.ErrorHandler != nil {
			p.ErrorHandler(err, w)
		} else if err != nil {
			DefaultErrorHandler(err, w)
		}
	}
}

type HandlerParams struct {
	TemplateName string
	Model        interface{}
	ErrorHandler func(error, http.ResponseWriter)
	Template     TemplateExecutor
}
