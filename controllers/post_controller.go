package controllers

import (
	"NCO-Chat-Bot/models"
	"NCO-Chat-Bot/services"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostController - контроллер для POST запросов (создание данных)
type PostController struct {
	ps *services.PostService
}

func NewPostController(ps *services.PostService) *PostController {
	return &PostController{
		ps: ps,
	}
}

// ================== HTTP HANDLERS ==================

// CreateNCO - обработчик POST запроса для создания новой НКО
func (c *PostController) CreateNCO(w http.ResponseWriter, r *http.Request) {
	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		c.ps.WriteJSON(w, http.StatusMethodNotAllowed, &models.Response{
			Status: "error",
			Error:  "Метод не разрешен. Используйте POST",
		})
		return
	}

	// Проверяем Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		c.ps.WriteJSON(w, http.StatusBadRequest, &models.Response{
			Status: "error",
			Error:  "Content-Type должен быть application/json",
		})
		return
	}

	// Декодируем JSON из тела запроса
	var ncoRequest models.NCO
	err := json.NewDecoder(r.Body).Decode(&ncoRequest)
	if err != nil {
		c.ps.WriteJSON(w, http.StatusBadRequest, &models.Response{
			Status: "error",
			Error:  fmt.Sprintf("Неверный JSON формат: %v", err),
		})
		return
	}

	// Валидируем данные
	if validationErr := c.validateNCORequest(ncoRequest); validationErr != nil {
		c.ps.WriteJSON(w, http.StatusBadRequest, validationErr)
		return
	}

	fmt.Printf("🔄 POST обработчик: получен запрос на создание НКО - %s\n", ncoRequest.Name)

	// Вызываем бизнес-логику
	response := c.ps.SaveNCO(ncoRequest)

	// Отправляем ответ
	if response.Status == "error" {
		c.ps.WriteJSON(w, http.StatusInternalServerError, response)
	} else {
		c.ps.WriteJSON(w, http.StatusCreated, response)
	}
}

// ================== VALIDATION ==================

// validateNCORequest - валидация данных НКО
func (c *PostController) validateNCORequest(req models.NCO) *models.Response {
	if req.Name == "" {
		return &models.Response{
			Status: "error",
			Error:  "Поле 'name' обязательно для заполнения",
		}
	}

	if req.X < -180 || req.X > 180 {
		return &models.Response{
			Status: "error",
			Error:  "Координата X должна быть в диапазоне от -180 до 180",
		}
	}

	if req.Y < -90 || req.Y > 90 {
		return &models.Response{
			Status: "error",
			Error:  "Координата Y должна быть в диапазоне от -90 до 90",
		}
	}

	if req.Category == "" {
		return &models.Response{
			Status: "error",
			Error:  "Поле 'category' обязательно для заполнения",
		}
	}

	return nil
}
