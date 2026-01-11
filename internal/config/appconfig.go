package config

var (
	Version = "0.0.1"
)

type ApplicationConfig struct {
	envValues *EnvConfig
}

func NewApplicationConfig() *ApplicationConfig {
	envValues := NewEnvironmentConfig()

	return &ApplicationConfig{
		envValues: envValues,
	}
}

func (cfg *ApplicationConfig) ServiceName() string {
	return cfg.envValues.ServiceName
}

func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}
