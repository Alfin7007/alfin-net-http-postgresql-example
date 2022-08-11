package config

import "os"

func Secret_JWT() string {
	return os.Getenv("SECRET_JWT")
}
