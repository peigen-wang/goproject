package middleware

import (
	"log"
	"net/http"
)

//日志记录中间件
func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next(rw, r)
	}
}
