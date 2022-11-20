package data

import (
	"encoding/xml"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const IDRegexp = "[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}"

type Data struct {
	ID         string     `json:"id,omitempty" xml:"id,attr,omitempty" yaml:"id"`
	AddedAt    *time.Time `json:"added_at,omitempty" xml:"added_at,attr,omitempty" yaml:"added_at"`
	ModifiedAt *time.Time `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

func (d *Data) GenerateID() {
	d.ID = uuid.NewString()
}

type UserDataKey struct{}

type User struct {
	XMLName xml.Name `json:"-" xml:"user" yaml:"-"`
	Data
	Email    string   `json:"email" xml:"email" yaml:"email" validate:"required,email"`
	Username string   `json:"username" xml:"username" yaml:"username" validate:"required"`
	Hash     string   `json:"hash,omitempty" xml:"hash,omitempty" yaml:"hash,omitempty" validate:"required"`
	Roles    []string `json:"roles,omitempty" xml:"roles>role,omitempty" yaml:"roles,omitempty"`
	isHashed bool
}

const USER_ROLE_ADMIN = "ADMIN"
const USER_ROLE_USER = "USER"

func (u *User) HashPassword() (err error) {
	if !u.isHashed {
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(u.Hash), bcrypt.DefaultCost)

		if err == nil {
			u.Hash = string(hashBytes)
			u.isHashed = true
		}
	}
	return err
}

type Joke struct {
	XMLName xml.Name `json:"-" xml:"joke" yaml:"-"`
	Data
	Description string   `json:"description" xml:"description" yaml:"description" validate:"max=255"`
	Text        string   `json:"text" xml:"text" yaml:"text" validate:"required,max=255"`
	Tags        []string `json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
	AuthorID    string   `json:"author_id" xml:"author_id,attr" yaml:"author_id" validate:"required,uuid"`
	User        *User    `json:"user,omitempty" yaml:"user,omitempty" xml:"user,omitempty"`
	Language    string   `json:"lang" xml:"lang,attr" yaml:"lang" validate:"required,bcp47_language_tag"`
	Stars       *float64 `json:"stars,omitempty"  xml:"stars,attr,omitempty" yaml:"stars,omitempty"`
}

type Rating struct {
	XMLName xml.Name `json:"-" xml:"rating" yaml:"-"`
	Data
	Stars   float64 `json:"stars" xml:"start,attr" yaml:"stars" validate:"required,lte=5,gte=0"`
	UserID  string  `json:"user_id" xml:"user_id,attr" yaml:"user_id" validate:"required,uuid"`
	JokeID  string  `json:"joke_id" xml:"joke_id,attr" yaml:"joke_id" validate:"required,uuid"`
	Comment string  `json:"comment,omitempty" xml:",cdata,omitempty" yaml:"comment,omitempty" validate:"max=255"`
}
