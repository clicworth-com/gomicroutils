package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/clicworth-com/gomicroutils/v3/grpcclient"
	"github.com/gorilla/mux"
)

func RequestAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		authorised := "false"
		m["authorised"] = authorised
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) == 2 {
			jwtToken := authHeader[1]
			claims, err := grpcclient.GetAuthClient().Verify(jwtToken, mux.CurrentRoute(r).GetName())
			if err != nil {
				ctx := context.WithValue(r.Context(), "claims", m)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}
			if claims.Authorised {
				authorised = "true"
			}
			m["authorised"] = authorised
			m["name"] = claims.Name
			m["emailid"] = claims.EmailId
			m["phone"] = claims.PhoneNumber
			m["role"] = claims.Role

			ctx := context.WithValue(r.Context(), "claims", m)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			ctx := context.WithValue(r.Context(), "claims", m)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
