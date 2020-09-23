package main

type Config struct {
	Port           int
	Host           string
	ConnectionHost string
	ConnectionDB   string
	ConnectionUser string
}

func Values(text string) Config {
	config := Config{}
	config.Port = 8000
	config.Host = "localhost"
	config.ConnectionHost = "db.mysecretapp.dev"
	config.ConnectionDB = "secretapp"
	config.ConnectionUser = "root"
	return config
}

func main() {
}
