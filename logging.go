package lib

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Logger() (*log.Logger, *os.File) {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(f, "", 0)

	return logger, f
}

func Decorator(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		logger, logFile := Logger()
		logger.Println(t.Format(time.RFC3339), r.RemoteAddr, r.Method, r.RequestURI)
		logFile.Close()
		f(w, r)
	}
}
