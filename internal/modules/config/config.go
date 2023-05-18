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

	DBDriver         string `envconfig:"DB_DRIVER" default:"postgres"`
	DBConnString     string `envconfig:"DB_CONN_STRING"`
	SQLSlowThreshold int    `envconfig:"SQL_SLOW_THRESHOLD" default:"600"`
	TraceSQLCommands bool   `envconfig:"TRACE_SQL_COMMANDS" default:"false"`
	AutoMigrate      bool   `envconfig:"AUTO_MIGRATE" default:"false"`

	ElasticHost         string `envconfig:"ELASTIC_HOST" default:"elastic"`
	ElasticPort         string `envconfig:"ELASTIC_PORT" default:"9200"`
	TraceElasticQueries bool   `envconfig:"TRACE_ELASTIC_QUERIES" default:"false"`

	SearchService string `envconfig:"SEARCH_SERVICE" default:"postgres" validate:"oneof=postgres elastic"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"info" validate:"oneof=debug info warn error dpanic panic fatal"`

	HTTPServerHost         string        `envconfig:"HTTP_SERVER_HOST" default:"0.0.0.0"`
	HTTPServerPort         string        `envconfig:"HTTP_SERVER_PORT" default:"8080"`
	HTTPServerReadTimeOut  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"10m"`
	HTTPServerWriteTimeOut time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"13m"`
	HTTPServerPrefork      bool          `envconfig:"HTTP_SERVER_PREFORK" default:"false"`

	TimeEval     bool   `envconfig:"TIME_EVAL" default:"false"`
	TimeEvalFile string `envconfig:"TIME_EVAL_FILE" default:"time.txt"`
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
