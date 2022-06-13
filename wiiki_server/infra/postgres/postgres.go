package postgres

import (
	"fmt"
	"wiiki_server/common/config"
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
