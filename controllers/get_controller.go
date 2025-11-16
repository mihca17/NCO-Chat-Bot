package controllers

import (
	"NCO-Chat-Bot/logger"
	"NCO-Chat-Bot/models"
	"NCO-Chat-Bot/services"
	"net/http"
	"strconv"
)

// NCOController - контроллер
type GetController struct {
	g      *services.GetService
	logger *logger.Logger
}

func NewGetController(g *services.GetService, logger *logger.Logger) *GetController {
	return &GetController{
		g:      g,
		logger: logger,
	}
}

// GetNCOByID - обработчик GET запроса для получения НКО по ID
func (c *GetController) GetNCOByID(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id из query string
	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		c.g.WriteJSON(w, http.StatusBadRequest, &models.Response{
			Status: "error",
			Error:  "Параметр id обязателен",
		})
		return
	}

	// Преобразуем строку в int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.g.WriteJSON(w, http.StatusBadRequest, &models.Response{
			Status: "error",
			Error:  "Параметр id должен быть числом",
		})
		return
	}

	// Проверяем, что id положительный
	if id <= 0 {
		c.g.WriteJSON(w, http.StatusBadRequest, &models.Response{
			Status: "error",
			Error:  "Параметр id должен быть положительным числом",
		})
		return
	}

	c.logger.Info("GET Обработчик: получен GET запрос с id=" + strconv.FormatInt(id, 10))

	// Вызываем бизнес-логику
	response := c.g.GetNCOByID(id)

	// Отправляем ответ
	if response.Status == "error" {
		c.g.WriteJSON(w, http.StatusNotFound, response)
	} else {
		c.g.WriteJSON(w, http.StatusOK, response)
	}
}

// GetNCOsByCity - обработчик GET запроса для получения НКО по городу
//func (c *GetController) GetNCOsByCity(w http.ResponseWriter, r *http.Request) {
//	city := r.URL.Query().Get("city")
//
//	if city == "" {
//		c.g.WriteJSON(w, http.StatusBadRequest, &models.Response{
//			Status: "error",
//			Error:  "Параметр city обязателен",
//		})
//		return
//	}
//
//	fmt.Printf("🔄 Обработчик: получен GET запрос с city=%s\n", city)
//
//	// Вызываем бизнес-логику
//	response := c.g.getNCOsByCity(city)
//
//	if response.Status == "error" {
//		c.g.WriteJSON(w, http.StatusNotFound, response)
//	} else {
//		c.g.WriteJSON(w, http.StatusOK, response)
//	}
//}

// GetAllNCOs - обработчик GET запроса для получения всех НКО
func (c *GetController) GetAllNCOs(w http.ResponseWriter, r *http.Request) {
	c.logger.Info("GET Обработчик: получен GET запрос на все НКО")

	// Вызываем бизнес-логику
	response := c.g.GetAllNCOs()
	c.g.WriteJSON(w, http.StatusOK, response)
}
