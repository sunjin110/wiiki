package testtool

import (
	"context"
	"log"
	"path"
	"path/filepath"
	"runtime"
	"time"
	"wiiki_server/common/config"
	"wiiki_server/common/utils/fileutil"
	"wiiki_server/common/utils/jsonutil"
	"wiiki_server/common/utils/tomlutil"
	"wiiki_server/common/wiikictx"
	"wiiki_server/infra/postgres"
)

func Config() *config.WiikiConfig {
	testConfigPath := filepath.Join(TestRootDir(), "config", "local.toml")

	b, err := fileutil.GetBytes(testConfigPath)
	Chk(err)
	conf := &config.WiikiConfig{}
	err = tomlutil.Unmarshal(b, conf)
	Chk(err)
	return conf
}

func Context() (context.Context, func(commit bool)) {
	conf := Config()
	engine, err := postgres.New(conf.Postgres[0])
	Chk(err)

	transactionSession := engine.NewSession()
	err = transactionSession.Begin()
	Chk(err)

	tx := &wiikictx.Transaction{
		TransactionDB: transactionSession,
		ReadOnlyDB:    engine.NewSession(),
		IsTransaction: true,
	}

	common := &wiikictx.Common{
		TxTime:      time.Now(),
		AccessToken: "test_access_token",
	}

	ctx := wiikictx.WithTransaction(context.Background(), tx)
	ctx = wiikictx.WithCommon(ctx, common)

	close := func(commit bool) {
		if commit {
			err := tx.TransactionDB.Commit()
			Chk(err)
		} else {
			err := tx.TransactionDB.Rollback()
			Chk(err)
		}

		// start new session
		transactionSession := engine.NewSession()
		err := transactionSession.Begin()
		Chk(err)
		tx.TransactionDB = transactionSession
	}

	return ctx, close
}

func TestRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Join(filepath.Dir(d), "..")
}

func Chk(err error) {
	if err != nil {
		panic(err)
	}
}

func Log(logPrefix string, obj interface{}) {
	log.Println(logPrefix, jsonutil.MustMarshal(obj))
}
