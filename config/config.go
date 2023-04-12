package config

import (
	"context"
	"fmt"
	"github.com/linzhengen/mii-go/pkg/logger"
	"github.com/sethvargo/go-envconfig"
	"time"
)

type EnvConfig struct {
	AppEnv  string `env:"APP_ENV,required"`
	WebHost string `env:"WEB_HOST,default="`
	WebPort int    `env:"WEB_PORT,default=8080"`
	Log     Log
	MySQL   MySQL
	CORS    CORS
}

type Log struct {
	Level  int    `env:"LOG_LEVEL,default=5"` //  1: fatal 2: error, 3: warn, 4: info, 5: debug, 6: trace
	Format string `env:"LOG_FORMAT,default=json"`
}

type MySQL struct {
	User         string        `env:"MYSQL_USER,required"`
	Pass         string        `env:"MYSQL_PASS,required"`
	Port         int           `env:"MYSQL_PORT,required"`
	Host         string        `env:"MYSQL_HOST,required"`
	DBName       string        `env:"MYSQL_DB_NAME,required"`
	MaxLifetime  time.Duration `env:"MYSQL_MAX_LIFE_TIME,default=7200s"`
	MaxOpenConns int           `env:"MYSQL_MAX_OPEN_CONNS,default=10"`
	MaxIdleConns int           `env:"MYSQL_MAX_IDLE_CONNS,default=10"`
}

func (m MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		m.User, m.Pass, m.Host, m.Port, m.DBName)
}

type CORS struct {
	AllowOrigins     []string      `env:"CORS_ALLOW_ORIGINS,default=*"`
	AllowMethods     []string      `env:"CORS_ALLOW_METHODS,default=GET,POST,PUT,DELETE,PATCH"`
	AllowHeaders     []string      `env:"CORS_ALLOW_HEADERS"`
	AllowCredentials bool          `env:"CORS_ALLOW_CREDENTIALS,default=true"`
	MaxAge           time.Duration `env:"CORS_MAX_AGE,default=7200s"`
}

func New(ctx context.Context) EnvConfig {
	var c EnvConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		logger.Severe(err)
	}
	return c
}
