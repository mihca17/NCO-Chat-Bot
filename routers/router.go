package routers

import (
	"NCO-Chat-Bot/controllers"
	"NCO-Chat-Bot/logger"
	"net/http"
	"path/filepath"
	"strings"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Если это OPTIONS запрос (preflight), сразу отвечаем
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

type Server struct {
	address string
	port    string
	router  *http.ServeMux
	gc      *controllers.GetController
	pc      *controllers.PostController
	logger  *logger.Logger
}

func NewServer(address, port string, gc *controllers.GetController, pc *controllers.PostController, logger *logger.Logger) *Server {
	return &Server{
		address: address,
		port:    port,
		router:  http.NewServeMux(),
		gc:      gc,
		pc:      pc,
		logger:  logger,
	}
}

func (s *Server) Start() error {
	s.setupRoutes()

	s.logger.Success("🚀 Сервер запущен на http://" + s.address + ":" + s.port)
	s.logger.Success("📁 Обслуживаются статические файлы из директории static/")

	return http.ListenAndServe(s.address+":"+s.port, corsMiddleware(s.router))
}

func (s *Server) setupRoutes() {
	// GET маршруты
	s.router.HandleFunc("GET /api/nco", s.gc.GetNCOByID)
	s.router.HandleFunc("GET /api/nco/all", s.gc.GetAllNCOs)

	// POST маршруты
	s.router.HandleFunc("POST /api/nco", s.pc.SaveNCO)

	// Статические файлы
	s.router.HandleFunc("/static/", s.staticHandler)
	s.router.HandleFunc("/", s.homeHandler)
}

func (s *Server) staticHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.URL.Path == "/" {
		http.ServeFile(w, r, "index.html")
	} else {
		http.NotFound(w, r)
	}
}
