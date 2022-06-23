package config


type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int `mapstructure:"port" json:"port"`
}

type ServerConfig struct {
	Name string  `mapstructure:"name" json:"name"`

	ConsulConfig ConsulConfig  `mapstructure:"consul" json:"consul"`

}