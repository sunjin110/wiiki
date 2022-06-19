package middleware

import (
	"net/http"
	"time"
	"wiiki_server/common/wiikictx"
	"wiiki_server/infra/http/middleware"
)

type authImpl struct {
}

func NewAuth() middleware.Auth {
	return &authImpl{}
}

func (impl *authImpl) Auth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO auth token

			r = r.WithContext(wiikictx.WithCommon(r.Context(), &wiikictx.Common{
				TxTime:      time.Now(),
				AccessToken: "todo",
			}))
			next.ServeHTTP(w, r)
		})
	}
}
