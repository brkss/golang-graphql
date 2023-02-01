package middleware

import (
	"fmt"
	"net/http"
)



func AuthMiddleware() func(http.Handler) http.Handler{
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			fmt.Println(">>> Middleware Called");
			next.ServeHTTP(w, r)
		})
	}
}
