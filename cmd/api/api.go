package api

import (
	"calibri/internal/config"
	"calibri/internal/service/client"
	"calibri/pkg/logger"
	"calibri/pkg/logger/middleware"
	"fmt"
	"log"
	"net/http"
	"github.com/jmoiron/sqlx"
)

type ApiServer struct {
	addr string
	db   *sqlx.DB
}

func NewApiServer(addr string, db *sqlx.DB) *ApiServer {
	return &ApiServer{
		addr: addr,
		db:   db,
	}
}

func (s *ApiServer) Run() error {
	mux := http.NewServeMux()

	storeClient := client.NewStore(s.db)
	clientHandler := client.NewHandler(storeClient)
	clientHandler.RegisterRoutes(mux)

	loggerJournal, err := logger.NewLogger(s.db)
	if err != nil {
		return fmt.Errorf("failed to create logger: %w", err)
	}
	LoggerMux := middleware.LoggerMiddleware(mux, loggerJournal)
	cors := config.Cors(LoggerMux)

	if err := loggerJournal.LoggerBasic(logger.INFO_LOG, "Server Calibri started on port 8070"); err != nil {
		log.Println("Failed to log to database:", err)
	}
	fmt.Println("Server started on port 8070")

	if err := http.ListenAndServe(":8070", cors); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return fmt.Errorf("server error")
}
