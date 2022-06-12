package config

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
