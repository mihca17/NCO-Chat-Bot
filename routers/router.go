package routers

import (
	"NCO-Chat-Bot/controllers"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type Server struct {
	address string
	port    string
	router  *http.ServeMux
	gc      *controllers.GetController
	pc      *controllers.PostController
}

func NewServer(address, port string, gc *controllers.GetController, pc *controllers.PostController) *Server {
	return &Server{
		address: address,
		port:    port,
		router:  http.NewServeMux(),
		gc:      gc,
		pc:      pc,
	}
}

func (s *Server) setupRoutes() {
	// Инициализируем обработчики

	// GET маршруты
	s.router.HandleFunc("GET /api/nco", s.gc.GetNCOByID)
	s.router.HandleFunc("GET /api/nco/all", s.gc.GetAllNCOs)

	// POST маршруты
	s.router.HandleFunc("POST /api/nco", s.pc.CreateNCO)

	// Статические файлы
	s.router.HandleFunc("/static/", s.staticHandler)
	s.router.HandleFunc("/", s.homeHandler)
}

func (s *Server) staticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	if filePath == "" {
		http.NotFound(w, r)
		return
	}

	if strings.Contains(filePath, "..") {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	ext := strings.ToLower(filepath.Ext(filePath))
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "application/javascript")
	case ".png":
		w.Header().Set("Content-Type", "image/png")
	case ".jpg", ".jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
	case ".html":
		w.Header().Set("Content-Type", "text/html")
	default:
		w.Header().Set("Content-Type", "text/plain")
	}

	http.ServeFile(w, r, "static/"+filePath)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "static/index.html")
	} else {
		http.NotFound(w, r)
	}
}

func (s *Server) Start() error {
	s.setupRoutes()

	log.Printf("🚀 Сервер запущен на http://%s:%s", s.address, s.port)
	log.Printf("📁 Обслуживаются статические файлы из директории static/")

	return http.ListenAndServe(s.address+":"+s.port, s.router)
}
