package config

type Config struct {
	DBConnection   string
	ServerAddress  string
	AllowedOrigins []string
}

var AppConfig Config

func LoadConfig() {
	AppConfig = Config{
		DBConnection:   "anon:tokyo@tcp(localhost:3306)/mygo?charset=utf8mb4&parseTime=True&loc=Local",
		ServerAddress:  "0.0.0.0:8080",
		AllowedOrigins: []string{"http://localhost:3000", "https://c8763yee.github.io"},
	}
}
