package main

import (
    "encoding/json"
    "fmt"
    "net/http"
	"bytes"
    "github.com/go-chi/chi/v5"
)

// Task ...
type Task struct {
	ID           string   `json:"id"`
	Description  string   `json:"description"`
	Note         string   `json:"note"`
	Applications []string `json:"applications"`
}

var tasks = map[string]Task{
	"1": {
		ID:          "1",
		Description: "Сделать финальное задание темы REST API",
		Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
		},
	},
	"2": {
		ID:          "2",
		Description: "Протестировать финальное задание с помощью Postmen",
		Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
		Applications: []string{
			"VS Code",
			"Terminal",
			"git",
			"Postman",
		},
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
    // сериализуем данные из слайса tasks
    resp, err := json.Marshal(tasks)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // в заголовок записываем тип контента, у нас это данные в формате JSON
    w.Header().Set("Content-Type", "application/json")
    // так как все успешно, то статус OK
    w.WriteHeader(http.StatusOK)
    // записываем сериализованные в JSON данные в тело ответа
    w.Write(resp)
}

func main() {
	r := chi.NewRouter()

	r.Get("/tasks", getTasks)

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
