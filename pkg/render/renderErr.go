package render

import (
	"fmt"
	"net/http"
	"text/template"
)

type ErrorData struct {
	Status  int
	Message string
}

func RenderError(w http.ResponseWriter, status int, message string) error {
	// Загружаем шаблон ошибки
	t, err := template.New("error.html").ParseFiles("web/templates/error.html")
	if err != nil {
		fmt.Println("Error parsing error template:", err)
		return err
	}

	w.WriteHeader(status)

	data := ErrorData{
		Status:  status,
		Message: message,
	}

	// Рендерим шаблон с переданными данными
	if err := t.Execute(w, data); err != nil {
		fmt.Println("Error executing error template:", err)
		return err
	}
	return nil
}
