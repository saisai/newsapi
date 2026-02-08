package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/saisai/newsapi/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	logger.Info("server starting on port 8080")

	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
