package libs

import "encoding/xml"

type StandardReponse struct {
	XMLName xml.Name    `json:"-" xml:"response"`
	Status  int64       `json:"status" xml:"status,attr"`
	Data    interface{} `json:"data,omitempty" xml:"data>data"`
	Message string      `json:"message,omitempty" xml:"message,omitempty"`
}

type AuthRequest struct {
	UsernameOrEmail string `json:"user"`
	Password        string `json:"password"`
}

type AuthResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
