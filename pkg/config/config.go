package config

var cfg = make(map[string]string)

func Get(key string) string {
	return cfg[key]
}

func Set(key, value string) {
	cfg[key] = value
}
