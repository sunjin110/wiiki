package middleware

import "net/http"

type Transaction interface {
	Transaction() func(http.Handler) http.Handler
}

type ErrorHandling interface {
	ErrorHandling() func(http.Handler) http.Handler
}

type Auth interface {
	Auth() func(http.Handler) http.Handler
}
