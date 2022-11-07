package handlers

import "net/http"

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Get(w http.ResponseWriter, r *http.Request) {

}
