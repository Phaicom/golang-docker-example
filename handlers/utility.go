package handlers

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Printf("Error while parsing template: %v", err)
	}
	err = t.Execute(w, templateData)
	if err != nil {
		log.Printf("Error while execute template: %v", err)
	}
}

func RandomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(65 + rand.Intn(25))
	}
	return string(bytes)
}
