package config

import (
	"calend/internal/modules/app"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"time"
)

// Config ...
type Config struct {
	Secret string `envconfig:"SECRET"`

	DBDriver     string `envconfig:"DB_DRIVER" default:"postgres"`
	DBConnString string `envconfig:"DB_CONN_STRING"`

	ElasticHost string `envconfig:"ELASTIC_HOST" default:"elastic"`
	ElasticPort string `envconfig:"ELASTIC_PORT" default:"9200"`

	SearchService    string `envconfig:"SEARCH_SERVICE" default:"postgres" validate:"oneof=postgres elastic"`
	SQLSlowThreshold int    `envconfig:"SQL_SLOW_THRESHOLD" default:"600"`
	TraceSQLCommands bool   `envconfig:"TRACE_SQL_COMMANDS" default:"false"`
	AutoMigrate      bool   `envconfig:"AUTO_MIGRATE" default:"false"`
	LogLevel         string `envconfig:"LOG_LEVEL" default:"info" validate:"oneof=debug info warn error dpanic panic fatal"`

	HTTPServerHost         string        `envconfig:"HTTP_SERVER_HOST" default:"0.0.0.0"`
	HTTPServerPort         string        `envconfig:"HTTP_SERVER_PORT" default:"8080"`
	HTTPServerReadTimeOut  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"10m"`
	HTTPServerWriteTimeOut time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"13m"`
	HTTPServerPrefork      bool          `envconfig:"HTTP_SERVER_PREFORK" default:"false"`
	HTTPServerDevelopMode  bool          `envconfig:"HTTP_SERVER_DEVELOP_MODE" default:"false"`
}

func NewConfig(app app.App, logger *zap.Logger, logLevel zap.AtomicLevel) (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}

	logger.Info("получена конфигурация", zap.Any("config", config))

	// Принудительно инициализируем уровень логирования из конфигурации
	err = logLevel.UnmarshalText([]byte(config.LogLevel))
	if err != nil {
		return Config{}, err
	}

	return config, err
}
