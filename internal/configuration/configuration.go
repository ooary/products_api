package configuration

type Configuration struct {
	PORT   string `mapstructure:"PORT"`
	DB_URL string `mapstructure:"DB_URL"`
}
