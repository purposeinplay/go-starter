package logs

import (
	"fmt"
	"net/http"
	"time"

	chimiddleware "github.com/go-chi/chi/middleware"
	ctx "github.com/oakeshq/go-starter/context"

	"github.com/sirupsen/logrus"
)

func NewStructuredLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return chimiddleware.RequestLogger(&structuredLogger{logger})
}

type structuredLogger struct {
	Logger *logrus.Logger
}

func (l *structuredLogger) NewLogEntry(r *http.Request) chimiddleware.LogEntry {
	entry := &structuredLoggerEntry{Logger: logrus.NewEntry(l.Logger)}
	logFields := logrus.Fields{}

	logFields["component"] = "api"
	logFields["method"] = r.Method
	logFields["path"] = r.URL.Path

	logFields["ts"] = time.Now().UTC().Format(time.RFC1123)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	if reqID := ctx.GetRequestID(r.Context()); reqID != "" {
		logFields["request_id"] = reqID
	}

	logFields["http_scheme"] = scheme
	logFields["http_proto"] = r.Proto
	logFields["http_method"] = r.Method

	logFields["remote_addr"] = r.RemoteAddr
	logFields["user_agent"] = r.UserAgent()

	logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

	entry.Logger = entry.Logger.WithFields(logFields)
	entry.Logger.Infoln("Request started")
	return entry
}

type structuredLoggerEntry struct {
	Logger logrus.FieldLogger
}

func (l *structuredLoggerEntry) Write(status, bytes int, elapsed time.Duration) {
	l.Logger = l.Logger.WithFields(logrus.Fields{
		"status":      status,
		"duration_ms": float64(elapsed.Nanoseconds()) / 1000000.0,
	})

	l.Logger.Info("Request completed")
}

func (l *structuredLoggerEntry) Panic(v interface{}, stack []byte) {
	l.Logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	}).Panic("unhandled request panic")
}

func GetLogEntry(r *http.Request) logrus.FieldLogger {
	entry, _ := chimiddleware.GetLogEntry(r).(*structuredLoggerEntry)
	if entry == nil {
		return logrus.NewEntry(logrus.StandardLogger())
	}
	return entry.Logger
}

func logEntrySetField(r *http.Request, key string, value interface{}) logrus.FieldLogger {
	if entry, ok := r.Context().Value(chimiddleware.LogEntryCtxKey).(*structuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithField(key, value)
		return entry.Logger
	}
	return nil
}

func logEntrySetFields(r *http.Request, fields logrus.Fields) logrus.FieldLogger {
	if entry, ok := r.Context().Value(chimiddleware.LogEntryCtxKey).(*structuredLoggerEntry); ok {
		entry.Logger = entry.Logger.WithFields(fields)
		return entry.Logger
	}
	return nil
}
