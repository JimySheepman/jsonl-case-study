package config

type Configurations struct {
	Redis        Redis
	AwsS3Config  AwsS3Config
	LoggerConfig LoggerConfig
}

type Redis struct {
	Address   string
	KeyPrefix string
	Username  string
	Password  string `json:"-"`
}

type AwsS3Config struct {
	Bucket    string
	Region    string
	SecretId  string `json:"-"`
	SecretKey string `json:"-"`
	Token     string `json:"-"`
}

type LoggerConfig struct {
	AppName         string
	LogLevel        int
	LogEncoding     string
	EnvironmentType string
}
