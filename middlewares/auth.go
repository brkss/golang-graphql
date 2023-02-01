package middleware

import (
	"context"
	"fmt"

	"github.com/brkss/golang-graphql/token"

	"net/http"
	"strings"
)


const (
	AuthorizationTypeBearer = "bearer"
	AuthorizationKeyHeader = "Authorization"
	AuthorizationPayloadKey = "payload"
)

func errorResponse(msg string) string{
	return fmt.Sprintf(`{"error": "%s"}`, msg)
}

func AuthMiddleware(maker token.Maker) func(http.Handler) http.Handler{
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			//fmt.Println(">>> Middleware Called");
			auth := r.Header.Get(AuthorizationKeyHeader)
		
			// no authorization provided pass for unprotected queries and mutations ! 
			if len(auth) == 0 {
				next.ServeHTTP(w, r)
				return;
			}
			
			fields := strings.Split(auth, " ")
			if len(fields) < 2 {
				http.Error(w, errorResponse("Invalid Token"), http.StatusForbidden)
				return;
			}
			// check type !
			if strings.ToLower(fields[0]) != AuthorizationTypeBearer {
				http.Error(w, errorResponse("Invalid Token Type !"), http.StatusForbidden)
				return;
			}

			
			token := fields[1]
			payload, err := maker.VerifyToken(token)	
			if err != nil {
				fmt.Printf("err: %+v", err)
				http.Error(w,  errorResponse("Invalid Token !"), http.StatusForbidden)
				return;
			}
			// inject payload through the request context !	
			ctx := context.WithValue(r.Context(), AuthorizationPayloadKey, payload) 
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func GetPayload(ctx context.Context) *token.Payload {
	payload, _ := ctx.Value(AuthorizationPayloadKey).(*token.Payload)
	return payload
}
