package config

type Config struct {
	Port string
	Otakus string
}

func GetConfig() Config {
	return Config{
		Port: ":3000",
		Otakus: "https://otakuanimesscc.com/",
	}
}
