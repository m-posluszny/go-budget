package config

type DbConf struct {
	Host     string `mapstructure:"DB_HOST"`
	Port     int    `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

type AuthConf struct {
	Secret  string `mapstructure:"AUTH_SECRET"`
	Expires int    `mapstructure:"AUTH_EXPIRES"`
}

type ServerConf struct {
	Host  string `mapstructure:"SERVER_HOST"`
	Debug bool   `mapstructure:"SERVER_DEBUG"`
	Mode  string `mapstructure:"SERVER_MODE"`
}

type RedisConf struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Size     int    `mapstructure:"REDIS_SIZE"`
	Password string `mapstructure:"REDIS_PASSWORD"`
}

type Config struct {
	Server ServerConf `mapstructure:",squash"`
	Db     DbConf     `mapstructure:",squash"`
	Redis  RedisConf  `mapstructure:",squash"`
	Auth   AuthConf   `mapstructure:",squash"`
}
