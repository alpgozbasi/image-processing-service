package handler

import (
	"github.com/alpgozbasi/image-processing-service/internal/config"
	"github.com/alpgozbasi/image-processing-service/internal/processor"
	"github.com/alpgozbasi/image-processing-service/pkg/logger"
	"io"
	"net/http"
	"strconv"
)

func ConvertImageHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// limit file size
		r.Body = http.MaxBytesReader(w, r.Body, cfg.MaxUploadSize)

		file, _, err := r.FormFile("file")
		if err != nil {
			logger.ErrorLogger.Println("form file error:", err)
			http.Error(w, "failed to read file", http.StatusBadRequest)
			return
		}

		// read file into memory
		data, err := io.ReadAll(file)
		if err != nil {
			logger.ErrorLogger.Println("file read error:", err)
			http.Error(w, "failed to read file", http.StatusInternalServerError)
			return
		}

		// quality param (optional) -> default 75
		qualityParam := r.FormValue("quality")
		if qualityParam == "" {
			qualityParam = "75"
		}
		quality, err := strconv.Atoi(qualityParam)
		if err != nil {
			logger.ErrorLogger.Println("quality parse error:", err)
			quality = 75
		}

		// convert to webp
		output, err := processor.ConvertToWebP(data, quality)
		if err != nil {
			logger.ErrorLogger.Println("conversion error:", err)
			http.Error(w, "failed to process image", http.StatusInternalServerError)
			return
		}

		// set headers and write response
		w.Header().Set("Content-Type", "image/webp")
		w.WriteHeader(http.StatusOK)
		w.Write(output)
	}
}
