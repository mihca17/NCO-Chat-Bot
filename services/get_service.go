package services

import (
	"NCO-Chat-Bot/database/repository"
	"NCO-Chat-Bot/logger"
	"NCO-Chat-Bot/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type GetService struct {
	repo   *repository.SQLiteRepository
	logger *logger.Logger
}

func NewGetService(repo *repository.SQLiteRepository, logger *logger.Logger) *GetService {
	return &GetService{
		repo:   repo,
		logger: logger,
	}
}

// ================== BUSINESS LOGIC ==================

// Получение НКО по ID
func (g *GetService) GetNCOByID(id int64) *models.Response {
	g.logger.Info("Контроллер: получен запрос на НКО с ID: " + strconv.FormatInt(id, 10))

	nco, err := g.repo.GetByID(id)
	if err != nil {
		return &models.Response{
			Status: "error",
			Error:  err.Error(),
		}
	}

	g.logger.Info("Контроллер: найдена НКО - " + nco.Name)

	return &models.Response{
		Status: "success",
		Data:   nco,
	}
}

// Получение НКО по городу
//func (g *GetService) getNCOsByCity(city string) *models.Response {
//	fmt.Printf("🎯 Контроллер: получен запрос на НКО в городе: %s\n", city)
//
//	ncos, err := g.repo.FindByCity(city)
//	if err != nil {
//		return &models.Response{
//			Status: "error",
//			Error:  err.Error(),
//		}
//	}
//
//	fmt.Printf("✅ Контроллер: найдено %d НКО в городе %s\n", len(ncos), city)
//
//	return &models.Response{
//		Status: "success",
//		Data:   ncos,
//	}
//}

// Получение всех НКО
func (g *GetService) GetAllNCOs() *models.Response {
	g.logger.Info("Контроллер: получен запрос на все НКО")

	ncos, err := g.repo.GetAll()
	if err != nil {
		return &models.Response{
			Status: "error",
			Error:  err.Error(),
		}
	}

	g.logger.Info("Контроллер: найдено " + strconv.FormatInt(int64(len(ncos)), 10) + " НКО")

	return &models.Response{
		Status: "success",
		Data:   ncos,
	}
}

// ================== HELPER METHODS ==================

// Вспомогательная функция для отправки JSON ответа
func (g *GetService) WriteJSON(w http.ResponseWriter, statusCode int, response *models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
