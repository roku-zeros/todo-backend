package config_utils

type MainConfig struct {
	Port    string                 `yaml:"port"`
	Mode    string                 `yaml:"mode"`
	Secrets map[string]interface{} `yaml:"secrets,omitempty"`
}

type AppConfig struct {
	Port string
}

type MongoConfig struct {
	Url string
}
