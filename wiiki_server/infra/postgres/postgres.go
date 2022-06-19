package postgres

import (
	"context"
	"fmt"
	"wiiki_server/common/config"
	"wiiki_server/common/wiikictx"
	"wiiki_server/common/wiikierr"

	"xorm.io/xorm"

	_ "github.com/lib/pq"
)

func New(conf *config.Postgres) (*xorm.Engine, error) {

	dataSourceName := generateDataSourceName(conf)
	engine, err := xorm.NewEngine("postgres", dataSourceName)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedCreateXormEngine, "dataSourceName=%s", dataSourceName)
	}

	err = engine.Ping()
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedPingXormEngine, "dataSourceName=%s", dataSourceName)
	}

	return engine, nil
}

func generateDataSourceName(conf *config.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", conf.Host, conf.Port, conf.User, conf.DBName, conf.Password)
}

func WithReadWriteDB(ctx context.Context, postgresEngine *xorm.Engine) (context.Context, func(err error), error) {

	transactionSession := postgresEngine.NewSession()
	err := transactionSession.Begin()
	if err != nil {
		return nil, nil, wiikierr.Bind(err, wiikierr.FailedBeginTransaction, "")
	}

	readOnly := postgresEngine.NewSession()

	tx := &wiikictx.Transaction{
		TransactionDB: transactionSession,
		ReadOnlyDB:    readOnly,
		IsTransaction: true,
	}

	close := func(err error) {
		if err != nil {
			rollbackErr := transactionSession.Rollback()
			wiikictx.AddError(ctx, rollbackErr)
		} else {
			commitErr := transactionSession.Commit()
			wiikictx.AddError(ctx, commitErr)
		}
	}

	return wiikictx.WithTransaction(ctx, tx), close, nil

}

func WithReadDB(ctx context.Context, postgresEngine *xorm.Engine) context.Context {
	tx := &wiikictx.Transaction{
		TransactionDB: nil,
		ReadOnlyDB:    postgresEngine.NewSession(),
		IsTransaction: false,
	}
	return wiikictx.WithTransaction(ctx, tx)
}
