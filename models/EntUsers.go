package models

import (
	"strconv"
	"tdez/requests"
	"time"
)

// EntUsers ...
type EntUsers struct {
	Code            int        `gorm:"column:use_code; primary_key:true"`
	Email           string     `gorm:"column:use_email"`
	Cnpj            *int       `gorm:"column:use_cnpj"`
	RazaoSocial     *string    `gorm:"column:use_razao_social"`
	Password        string     `gorm:"column:use_password"`
	ConfirmPassword *string    `gorm:"column:use_confirm_password"`
	Token           *string    `gorm:"column:use_token"`
	Type            int        `gorm:"column:use_type"`
	DateIns         *time.Time `gorm:"column:use_date_ins; default:now()"`
	DateDel         *time.Time `gorm:"column:use_date_del"`
}

// TableName Seta o nome da tabela
func (e *EntUsers) TableName() string {
	return "entity.ent_users"
}

//EntUsersFill preenche o model a partir de um request
func (e *EntUsers) EntUsersFill(req requests.EntUsers) error {
	e.Code = req.Code
	e.Email = req.Email

	if req.Cnpj != nil {
		cnpj, err := strconv.Atoi(*req.Cnpj)
		if err != nil {
			return err
		}
		e.Cnpj = &cnpj
	}

	e.RazaoSocial = req.RazaoSocial
	e.Password = req.Password
	e.ConfirmPassword = req.ConfirmPassword
	e.Token = req.Token
	e.Type = req.Type

	return nil
}
