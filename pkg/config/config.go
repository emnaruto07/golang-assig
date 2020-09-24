package config

var cfg = make(map[string]string)

func get(key string) string {
	return cfg[key]
}

func set(key, value string) {
	cfg[key] = value
}
