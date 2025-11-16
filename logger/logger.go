package logger

import (
	"io"
	"log"
	"os"
	"sync"
)

type Logger struct {
	successLogger *log.Logger
	errorLogger   *log.Logger
	infoLogger    *log.Logger
	logFile       *os.File
}

// Инициализация логгера с выводом в консоль и файл
func Init(logFilename string) (*Logger, error) {
	var initErr error
	var once sync.Once
	logger := &Logger{}

	once.Do(func() {
		file, err := os.OpenFile(logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			initErr = err
			return
		}
		multiWriter := io.MultiWriter(os.Stdout, file)

		successLogger := log.New(multiWriter, "✅ ", log.Ldate|log.Ltime)
		errorLogger := log.New(multiWriter, "❌ ", log.Ldate|log.Ltime)
		infoLogger := log.New(multiWriter, "ℹ️ ", log.Ldate|log.Ltime)

		logger = &Logger{
			successLogger: successLogger,
			errorLogger:   errorLogger,
			infoLogger:    infoLogger,
		}
	})
	return logger, initErr
}

// Функция закрытия соединения с файлом логов
func (l *Logger) Close() error {
	if l.logFile != nil {
		return l.logFile.Close()
	}
	return nil
}

// Логгирование успешных операций
func (l *Logger) Success(message string) {
	l.successLogger.Println(message)
}

// Логгирование ошибок
func (l *Logger) Error(message string, err error) {
	if err != nil {
		l.errorLogger.Printf("%s: %v", message, err)
	} else {
		l.errorLogger.Println(message)
	}
}

// Логгирование информационных сообщений
func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

// Логгирование фатальных ошибок и завершение выполнения программы
func (l *Logger) Fatal(message string, err error) {
	if err != nil {
		l.errorLogger.Fatalf("%s: %v", message, err)
	} else {
		l.errorLogger.Fatal(message)
	}
}
