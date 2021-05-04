package rest

import (
	"context"
	"net/http"
)

const RequestHeaderContextKey = "request_header"

type RequestHeaders struct {
	ClientId string `json:"client_id,omitempty"`
	DeviceId string `json:"device_id,omitempty"`
	Version  string `json:"version,omitempty"`
}

func RequestHeader() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var ctx = r.Context()
			ctx = context.WithValue(ctx, RequestHeaderContextKey, RequestHeaders{
				ClientId: r.Header.Get("X-Client-Id"),
				DeviceId: r.Header.Get("X-Device-Id"),
				Version:  r.Header.Get("X-Client-Version"),
			})
			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}
