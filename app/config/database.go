package config

import "github.com/joho/godotenv"

func LoadDBConfig() (map[string]string, error) {
	return godotenv.Read("db.env")
}
