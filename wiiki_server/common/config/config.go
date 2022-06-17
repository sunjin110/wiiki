package config

import (
	"wiiki_server/common/utils/fileutil"
	"wiiki_server/common/utils/tomlutil"
)

type WiikiConfig struct {
	Env      string      `toml:"env"`
	Port     string      `toml:"port"`
	Postgres []*Postgres `toml:"postgresses"`
}

type Postgres struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	DBName   string `toml:"dbname"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func New(path string) (*WiikiConfig, error) {
	b, err := fileutil.GetBytes(path)
	if err != nil {
		return nil, err
	}
	conf := &WiikiConfig{}
	err = tomlutil.Unmarshal(b, conf)
	if err != nil {
		return nil, err
	}
	return conf, nil
}
