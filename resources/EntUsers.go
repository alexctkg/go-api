package resources

import (
	"tdez/models"
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

//EntUsersResource preenche um ressource apartir de um model
func (e *EntUsers) EntUsersResource(mod models.EntUsers) {
	e.Code = mod.Code
	e.Email = mod.Email
	e.Cnpj = mod.Cnpj
	e.RazaoSocial = mod.RazaoSocial
	e.Password = mod.Password
	e.ConfirmPassword = mod.ConfirmPassword
	e.Token = mod.Token
	e.Type = mod.Type
}
