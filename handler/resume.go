package handler

import (
	"log"
	"net/http"
	"os"
	"time"
)

func PDF(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Println("Some error")
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		w.Header().Set("Content-Type", "application/pdf")

		http.ServeContent(w, r, filePath, fileStat(file), file)
	}
}

func fileStat(f *os.File) time.Time {
	stat, _ := f.Stat()
	return stat.ModTime()
}
