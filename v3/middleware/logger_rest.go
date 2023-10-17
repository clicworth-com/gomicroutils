package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/clicworth-com/gomicroutils/v3/errs"
	"github.com/clicworth-com/gomicroutils/v3/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type contextKey string

const (
	correlationIDCtxKey contextKey = "correlation_id"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	errMsg     string
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, ""}
}

func (lrw *loggingResponseWriter) Header() http.Header {
	return lrw.ResponseWriter.Header()
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(p []byte) (int, error) {
	if lrw.statusCode >= http.StatusBadRequest {
		aerr := errs.RestError{}
		if err := json.Unmarshal(p, &aerr); err == nil {
			lrw.errMsg = aerr.Message
		}
	}
	return lrw.ResponseWriter.Write(p)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.Get()

		correlationID := xid.New().String()

		ctx := context.WithValue(
			r.Context(),
			correlationIDCtxKey,
			correlationID,
		)

		r = r.WithContext(ctx)

		l = l.With(zap.String(string(correlationIDCtxKey), correlationID))

		w.Header().Add("X-Correlation-ID", correlationID)

		lrw := newLoggingResponseWriter(w)

		r = r.WithContext(logger.WithCtx(ctx, l))

		defer func(start time.Time) {
			l.Info(
				fmt.Sprintf(
					"%s request to %s completed",
					r.Method,
					r.RequestURI,
				),
				zap.String("method", r.Method),
				zap.String("url", r.RequestURI),
				zap.String("user_agent", r.UserAgent()),
				zap.Int("status_code", lrw.statusCode),
				zap.String("err_msg", lrw.errMsg),
				zap.Duration("elapsed_ms", time.Since(start)),
			)
		}(time.Now())

		next.ServeHTTP(lrw, r)
	})
}
