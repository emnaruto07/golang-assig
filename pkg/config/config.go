package config

import (
	"os"
)

var defs = map[string]string{
	"port":           "8000",
	"host":           "localhost",
	"connectionHost": "db.mysecretapp.dev",
	"connectionDB":   "secretapp",
	"connectionUser": "root",
}

var cfg = make(map[string]string)

func init() {
	cfg = make(map[string]string)
	for k, v := range defs {
		env := os.Getenv(k)
		if len(env) == 0 {
			env = v
		}
		cfg[k] = env
	}
}

func Get(key string) string {
	return cfg[key]
}

func Set(key, value string) {
	cfg[key] = value
}
