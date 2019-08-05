package handler

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func LogRequests(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := statusWriter{ResponseWriter: w}
		handler.ServeHTTP(&sw, r)
		duration := time.Now().Sub(start)
		userAgent := r.Header.Get("User-Agent")
		log.WithFields(log.Fields{
			"Host":       r.Host,
			"RemoteAddr": r.RemoteAddr,
			"Method":     r.Method,
			"RequestURI": r.RequestURI,
			"Proto":      r.Proto,
			"Status":     sw.status,
			"Length":     sw.length,
			"UserAgent":  userAgent,
			"Duration":   duration,
		}).Info("RequestLogging")
	}
}
