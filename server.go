package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Структура для получения данных от клиента
type Message struct {
	Text string `json:"text"`
}

// Структура для ответа сервера
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func StartServer(address string, port string) error {
	// Обработчик для главной страницы (отдаем HTML)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Обработчик для CSS
	http.HandleFunc("/main.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "main.css")
	})

	// Обработчик для получения сообщений (POST)
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "POST" {
			var msg Message

			err := json.NewDecoder(r.Body).Decode(&msg)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			fmt.Printf("📨 Получено сообщение от пользователя: %s\n", msg.Text)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Response{
				Status:  "success",
				Message: "Сообщение успешно доставлено на сервер",
			})
		}
	})

	// Обработчик для получения данных (GET)
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "GET" {
			fmt.Printf("📊 Получен GET запрос на информацию\n")

			// Генерируем примерные данные
			info := map[string]any{
				"server_time":    time.Now().Format("2006-01-02 15:04:05"),
				"server_uptime":  "Запущен только что",
				"requests_count": 42,
				"features": []string{
					"Обработка текстовых сообщений",
					"REST API",
					"Real-time обновления",
					"Красивый интерфейс",
				},
				"status": "🟢 Онлайн",
			}

			json.NewEncoder(w).Encode(Response{
				Status: "success",
				Data:   info,
			})
		}
	})

	log.Printf("🚀 Сервер запущен на http://%s:%s", address, port)
	return http.ListenAndServe(address+":"+port, nil)
}
