package main

import (
	"fmt"
	"github.com/alpgozbasi/image-processing-service/internal/config"
	"github.com/alpgozbasi/image-processing-service/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()
	r := router.NewRouter(cfg)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("info: server running on port %s", cfg.Port)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("error: %v", err)
	}
}
