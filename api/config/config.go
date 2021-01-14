package config

import cfg "github.com/Yalantis/go-config"

var config Config

func Get() *Config {
	return &config
}

func Load(fileName string) error {
	return cfg.Init(&config, fileName)
}

type (
	Config struct {
		//AppName            string `json:"app_name" envconfig:"API_APP_NAME" default:"api"`
		//LogPreset          string `json:"log_preset" envconfig:"API_LOG_PRESET" default:"development"`
		ListenURL string `json:"listen_url" envconfig:"API_LISTEN_URL" default:":5000"`
		//PaginationMaxLimit int64  `json:"pagination_max_limit" envconfig:"API_PAGINATION_MAX_LIMIT" default:"1000"`

		Postgres Postgres `json:"postgres"`
		Auth     Postgres `json:"auth"`
	}

	Auth struct {
		JWTSecret string `json:"jwt_secret"          envconfig:"JWT_SECRET"              default:"jwt secret yalantis lol 123 !@##33"`
	}

	Postgres struct {
		Host         string       `json:"host"          envconfig:"POSTGRES_HOST"              default:"localhost"`
		Port         string       `json:"port"          envconfig:"API_POSTGRES_PORT"          default:"5432"`
		Database     string       `json:"database"      envconfig:"API_POSTGRES_DATABASE"      default:"tasks"`
		User         string       `json:"user"          envconfig:"API_POSTGRES_USER"          default:"tasks_managment_user"`
		Password     string       `json:"password"      envconfig:"API_POSTGRES_PASSWORD"      default:"tasks_managment_user!@#$342"`
		PoolSize     int          `json:"pool_size"     envconfig:"API_POSTGRES_POOL_SIZE"     default:"10"`
		MaxRetries   int          `json:"max_retries"   envconfig:"API_POSTGRES_MAX_RETRIES"   default:"5"`
		ReadTimeout  cfg.Duration `json:"read_timeout"  envconfig:"API_POSTGRES_READ_TIMEOUT"  default:"10s"`
		WriteTimeout cfg.Duration `json:"write_timeout" envconfig:"API_POSTGRES_WRITE_TIMEOUT" default:"10s"`
	}
)
