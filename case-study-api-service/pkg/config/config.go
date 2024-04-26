package config

type Configurations struct {
	RestServer   RestServer
	Redis        Redis
	LoggerConfig LoggerConfig
}

type RestServer struct {
	Addr        string
	PprofEnable int
	Username    string `json:"-"`
	Password    string `json:"-"`
}

type LoggerConfig struct {
	AppName         string
	LogLevel        int
	LogEncoding     string
	EnvironmentType string
}

type Redis struct {
	Address   string
	KeyPrefix string
	Username  string
	Password  string `json:"-"`
}
