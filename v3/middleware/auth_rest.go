package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/clicworth-com/gomicroutils/v3/grpcclient"
	"github.com/gorilla/mux"
)

func RequestAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ai := grpcclient.AuthInfo{
			Authorised:  false,
			Name:        "",
			EmailId:     "",
			PhoneNumber: "",
			Role:        "",
		}
		byteArray, _ := json.Marshal(ai)

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			ctx := context.WithValue(r.Context(), "claims", string(byteArray))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		jwtToken := authHeader[1]
		claims, err := grpcclient.GetAuthClient().Verify(jwtToken, mux.CurrentRoute(r).GetName())
		if err != nil {
			ctx := context.WithValue(r.Context(), "claims", string(byteArray))
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}

		ai.Authorised = claims.Authorised
		ai.Name = claims.Name
		ai.EmailId = claims.EmailId
		ai.PhoneNumber = claims.PhoneNumber
		ai.Role = claims.Role

		byteArray, _ = json.Marshal(ai)

		ctx := context.WithValue(r.Context(), "claims", string(byteArray))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
