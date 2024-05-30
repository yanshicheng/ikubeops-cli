package global

type AppConfig struct {
	HttpPort          int    `mapstructure:"http_port" json:"http_port" http_port:"http_port" env:"APP_HTTP_PORT"`
	HttpAddr          string `mapstructure:"http_addr" json:"http_addr" yaml:"http_addr" env:"APP_HTTP_ADDR"`
	GrpcPort          int    `mapstructure:"grpc_port" json:"grpc_port" yaml:"grpc_port" env:"APP_GRPC_PORT"`
	GrpcAddr          string `mapstructure:"grpc_addr" json:"grpc_addr" yaml:"grpc_addr" env:"APP_GRPC_ADDR"`
	Language          string `mapstructure:"language" json:"language" yaml:"language" env:"APP_LANGUAGE"`
	MaxHeaderSize     int    `mapstructure:"max_header_size" json:"max_header_size" yaml:"max_header_size" env:"APP_MAX_HEADER_SIZE"`
	ReadTimeout       int    `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout" env:"APP_READ_TIMEOUT"`
	ReadHeaderTimeout int    `mapstructure:"read_header_timeout" json:"read_header_timeout" yaml:"read_header_timeout" env:"APP_READ_HEADER_TIMEOUT"`
	WriteTimeout      int    `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout" env:"APP_WRITE_TIMEOUT"`
	Tls               bool   `mapstructure:"tls" json:"tls" yaml:"tls" env:"APP_TLS"`
	CertFile          string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file" env:"APP_CERT_FILE"`
	KeyFile           string `mapstructure:"key_file" json:"key_file" yaml:"key_file" env:"APP_KEY_FILE"`
	ShutdownTimeout   int    `mapstructure:"shutdown_timeout" json:"shutdown_timeout" yaml:"shutdown_timeout" env:"APP_SHUTDOWN_TIMEOUT"`
}

type LoggerConfig struct {
	Output     string `json:"output" yaml:"output" mapstructure:"output" env:"LOG_OUTPUT"`
	Format     string `json:"format" yaml:"format" mapstructure:"format"  env:"LOG_FORMAT"`
	Level      string `json:"level" yaml:"level" mapstructure:"level"  env:"LOG_LEVEL"`
	Dev        bool   `json:"dev" yaml:"dev" mapstructure:"dev"  env:"LOG_DEV"`
	FilePath   string `json:"file_path" yaml:"file_path" mapstructure:"file_path"  env:"LOG_FILE_PATH"`
	MaxSize    int    `json:"max_size" yaml:"max_size" mapstructure:"max_size"  env:"LOG_MAX_SIZE"`
	MaxAge     int    `json:"max_age" yaml:"max_age" mapstructure:"max_age"  env:"LOG_MAX_AGE"`
	MaxBackups int    `json:"max_backups" yaml:"max_backups" mapstructure:"max_backups"  env:"LOG_MAX_BACKUPS"`
}

type CLiConfig struct {
	Demo string `mapstructure:"demo" json:"demo" yaml:"demo" env:"ClI_DEMO"`
}

type Config struct {
	App    AppConfig    `mapstructure:"app" json:"app" yaml:"app" env:"IKUBEOPS"`
	Logger LoggerConfig `mapstructure:"logger" json:"logger" yaml:"logger" env:"IKUBEOPS"`
	Cli    CLiConfig    `mapstructure:"cli" json:"cli" yaml:"cli" env:"IKUBEOPS"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		HttpPort:          8080,
		HttpAddr:          "0.0.0.0",
		GrpcPort:          8081,
		GrpcAddr:          "0.0.0.0",
		Language:          "zh",
		MaxHeaderSize:     1,
		ReadTimeout:       60,
		ReadHeaderTimeout: 60,
		WriteTimeout:      60,
		Tls:               false,
		KeyFile:           "",
		CertFile:          "",
		ShutdownTimeout:   60,
	}
}

func NewLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Output:     "stdout",
		Format:     "json",
		Level:      "debug",
		Dev:        true,
		FilePath:   "./logs",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 100,
	}
}
func NewCliConfig() *CLiConfig {
	return &CLiConfig{
		Demo: "demo",
	}
}

func NewDefaultConfig() *Config {
	return &Config{
		App:    *NewAppConfig(),
		Logger: *NewLoggerConfig(),
		Cli:    *NewCliConfig(),
	}
}
