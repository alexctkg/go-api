package resources

import (
	"crypto/rand"
	"fmt"
	"os"
	"tdez/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// EntUsers ...
type EntUsers struct {
	Code            int     `json:"code"`
	Email           string  `json:"email"`
	Cnpj            int     `json:"cnpj"`
	RazaoSocial     *string `json:"razao_social"`
	Password        string  `json:"password"`
	ConfirmPassword *string `json:"confirm_password"`
	Type            int     `json:"-"`
}

//EntUsersResource fill resource by model
func (e *EntUsers) EntUsersResource(mod models.EntUsers) {
	e.Code = mod.Code
	e.Email = mod.Email
	e.RazaoSocial = mod.RazaoSocial
	e.Password = mod.Password
	e.ConfirmPassword = mod.ConfirmPassword
	e.Type = mod.Type
}

//CreateTokenJWT
func CreateTokenJWT(user models.EntUsers) (string, error) {

	var remember bool
	b := make([]byte, 16)
	rand.Read(b)
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	var ttl time.Duration

	if user.Type == 0 {
		remember = true
		ttl = 168 * time.Hour
	} else {
		remember = false
		ttl = 24 * time.Hour
	}

	claims := jwt.MapClaims{
		"iat":      time.Now().Unix(),
		"nbf":      time.Now().Unix(),
		"exp":      time.Now().Add(ttl).Unix(),
		"jti":      uuid,
		"remember": remember,
		"access":   user.Type, //0 and 1
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
