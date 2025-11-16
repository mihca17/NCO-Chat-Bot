package services

import (
	"NCO-Chat-Bot/database/repository"
	"NCO-Chat-Bot/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// PostService - сервис для операций создания
type PostService struct {
	repo *repository.SQLiteRepository
}

func NewPostService(repo *repository.SQLiteRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

// CreateNCO - бизнес-логика создания новой НКО
func (s *PostService) SaveNCO(req models.NCO) *models.Response {
	fmt.Printf("🎯 PostService: создание НКО - %s\n", req.Name)

	// Создаем модель NCO из запроса
	nco := models.NCO{
		Name:        req.Name,
		X:           req.X,
		Y:           req.Y,
		Category:    req.Category,
		Description: req.Description,
		Contacts:    req.Contacts,
		City:        req.City,
		Region:      req.Region,
	}

	// Сохраняем в репозитории
	err := s.repo.SaveNCO(&nco)
	if err != nil {
		return &models.Response{
			Status: "error",
			Error:  fmt.Sprintf("Ошибка при создании НКО: %v", err),
		}
	}

	fmt.Printf("✅ PostService: НКО создана")

	return &models.Response{
		Status:  "success",
		Message: "НКО успешно создана",
		//Data:    "",
	}
}

// WriteJSON - вспомогательная функция для отправки JSON ответа
func (s *PostService) WriteJSON(w http.ResponseWriter, statusCode int, response *models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
