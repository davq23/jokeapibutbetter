package libs

import (
	"encoding/xml"
	"os"
)

type StandardReponse struct {
	XMLName xml.Name    `json:"-" xml:"response" yaml:"-"`
	Link    string      `json:"link,omitempty" xml:"link,omitempty" yaml:"link,omitempty"`
	Status  int64       `json:"status" xml:"status,attr" yaml:"status"`
	Data    interface{} `json:"data,omitempty" xml:"data>data" yaml:"data,omitempty"`
	Message string      `json:"message,omitempty" xml:"message,omitempty" yaml:"message,omitempty"`
	Token   string      `json:"token,omitempty" xml:"token,omitempty" yaml:"token,omitempty"`
}

type StandardReponseList struct {
	StandardReponse
	NextLink string `json:"next-link,omitempty" xml:"next-link,omitempty" yaml:"next-link,omitempty"`
	LastLink string `json:"last-link,omitempty" xml:"last-link,omitempty" yaml:"last-link,omitempty"`
}

type AuthRequest struct {
	XMLName         xml.Name `json:"-" xml:"auth" yaml:"-"`
	UsernameOrEmail string   `json:"user" xml:"user" yaml:"user" validate:"required"`
	Password        string   `json:"password" xml:"password" yaml:"password" validate:"required"`
}

type AuthResponse struct {
	XMLName   xml.Name `json:"-" xml:"auth" yaml:"-"`
	UserID    string   `json:"user_id" xml:"user_id,attr" yaml:"user_id"`
	Username  string   `json:"username" xml:"username,attr" yaml:"username"`
	Email     string   `json:"email" xml:"email,attr" yaml:"email"`
	Token     string   `json:"token" xml:",cdata" yaml:"token"`
	Roles     []string `json:"roles" xml:"roles>role" yaml:"roles"`
	ExpiresAt int      `json:"expires_at" xml:"expires_at,attr" yaml:"expires_at"`
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
	RefreshSecret     string   `json:"refresh_secret" xml:"refresh_secret" yaml:"refresh_secret"`
	Timezone          string   `json:"timezone" xml:"timezone" yaml:"timezone"`
	SSLMode           string   `json:"ssl_mode" xml:"ssl_mode" yaml:"ssl_mode"`
	CORSDomains       string   `json:"cors_domains" xml:"cors_domains" yaml:"cors_domains"`
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
