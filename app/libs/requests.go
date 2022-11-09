package libs

import (
	"encoding/xml"
	"os"
)

type StandardReponse struct {
	XMLName xml.Name    `json:"-" xml:"response"`
	Status  int64       `json:"status" xml:"status,attr"`
	Data    interface{} `json:"data,omitempty" xml:"data>data"`
	Message string      `json:"message,omitempty" xml:"message,omitempty"`
}

type AuthRequest struct {
	XMLName         xml.Name `json:"-" xml:"auth"`
	UsernameOrEmail string   `json:"user" xml:"user"`
	Password        string   `json:"password" xml:"password"`
}

type AuthResponse struct {
	XMLName  xml.Name `json:"-" xml:"auth"`
	UserID   string   `json:"user_id" xml:"user_id,attr"`
	Username string   `json:"username" xml:"username,attr"`
	Token    string   `json:"token" xml:",cdata"`
}

type ConfigResponse struct {
	XMLName           xml.Name `json:"-" xml:"config" yaml:"-"`
	JokeServicePort   string   `json:"joke_port" xml:"joke_port" yaml:"joke_port"`
	JokeServiceURL    string   `json:"joke_url" xml:"joke_url" yaml:"joke_url"`
	UserServicePort   string   `json:"user_port" xml:"user_port" yaml:"user_port"`
	UserServiceURL    string   `json:"user_url" xml:"user_url" yaml:"user_url"`
	RatingServicePort string   `json:"rating_port" xml:"rating_port" yaml:"rating_port"`
	RatingServiceURL  string   `json:"rating_url" xml:"rating_url" yaml:"rating_url"`
	DBHost            string   `json:"db_host" xml:"db_host" yaml:"db_host"`
	DBUser            string   `json:"db_user" xml:"db_user" yaml:"db_user"`
	DBPassword        string   `json:"db_password" xml:"db_password" yaml:"db_password"`
	DBName            string   `json:"db_name" xml:"db_name" yaml:"db_name"`
	APISecret         string   `json:"api_secret" xml:"api_secret" yaml:"api_secret"`
	Timezone          string   `json:"timezone" xml:"timezone" yaml:"timezone"`
	SSLMode           string   `json:"ssl_mode" xml:"ssl_mode" yaml:"ssl_mode"`
}

func (cr *ConfigResponse) FixValues() {
	if cr.JokeServiceURL == "" {
		cr.JokeServiceURL = "http://" + os.Getenv("DOCKER_IP")
	}
	if cr.UserServiceURL == "" {
		cr.UserServiceURL = "http://" + os.Getenv("DOCKER_IP")
	}
	if cr.RatingServiceURL == "" {
		cr.RatingServiceURL = "http://" + os.Getenv("DOCKER_IP")
	}
}
