package rest

import (
	"context"
	"net/http"
)

const DebugContextKey = "debug"

func Debug(debug bool) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, DebugContextKey, debug)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
