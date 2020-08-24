package models

import (
	"strconv"
	"tdez/requests"
	"time"
)

// EntUsers ...
type EntUsers struct {
	Code        int        `gorm:"column:use_code; primary_key:true"`
	Email       string     `gorm:"column:use_email"`
	Cnpj        *int       `gorm:"column:use_cnpj"`
	RazaoSocial *string    `gorm:"column:use_razao_social"`
	Password    string     `gorm:"column:use_password"`
	Token       *string    `gorm:"column:use_token"`
	Type        int        `gorm:"column:use_type"`
	DateIns     *time.Time `gorm:"column:use_date_ins; default:now()"`
	DateDel     *time.Time `gorm:"column:use_date_del"`
}

// TableName schema and table references
func (e *EntUsers) TableName() string {
	return "entity.ent_users"
}

//SuperUserFill fill model by resource
func (e *EntUsers) SuperUserFill(req requests.EntSuperUserStore) error {

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
	e.Type = req.Type

	return nil
}

//ExternaUserFill fill model by resource
func (e *EntUsers) ExternaUserFill(req requests.EntExternalUserStore) error {

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
	e.Type = req.Type

	return nil
}
