package middleware

import (
	"net/http"
	"wiiki_server/common/wiikictx"
	"wiiki_server/common/wiikierr"
	"wiiki_server/infra/http/middleware"

	"github.com/99designs/gqlgen/graphql"
	"xorm.io/xorm"
)

type transactionImpl struct {
	engine *xorm.Engine
}

func NewTransactionMiddleware(engine *xorm.Engine) middleware.Transaction {
	return &transactionImpl{
		engine: engine,
	}
}

// This middleware requires error handling middleware to be used first
func (impl *transactionImpl) Transaction() func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			engine := impl.engine
			transactionSession := engine.NewSession()
			err := transactionSession.Begin()
			if err != nil {
				err = wiikierr.Bind(err, wiikierr.FailedBeginTransaction, "postgres db")
				wiikictx.AddError(r.Context(), err)
				return
			}

			readOnly := engine.NewSession()

			tx := &wiikictx.Transaction{
				TransactionDB: transactionSession,
				ReadOnlyDB:    readOnly,
				IsTransaction: true,
			}

			r = r.WithContext(wiikictx.WithTransaction(r.Context(), tx))
			next.ServeHTTP(w, r)

			// check error
			errList := graphql.GetErrors(r.Context())
			if len(errList) > 0 {
				err := transactionSession.Rollback()
				if err != nil {
					err = wiikierr.Bind(err, wiikierr.FailedRollbackTransaction, "postgres db")
					wiikictx.AddError(r.Context(), err)
					return
				}
				return
			}

			err = transactionSession.Commit()
			if err != nil {
				err = wiikierr.Bind(err, wiikierr.FailedCommitTransaction, "postgres db")
				wiikictx.AddError(r.Context(), err)
				return
			}
			return
		})
	}
}
