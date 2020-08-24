package models

import (
	"strconv"
	"tdez/requests"
	"time"
)

// ResProduct ...
type ResProduct struct {
	Code            int        `gorm:"column:pro_code; primary_key:true"`
	CostumerCpfCnpj int        `gorm:"column:pro_costumermid_cpf_cnpj"`
	CostumerEmail   string     `gorm:"column:pro_costumermid_email"`
	CostumerType    int        `gorm:"column:pro_costumermid_type"` //0 cpf 1 cnpj
	Status          int        `gorm:"column:pro_status; default:0"`
	StatusReason    *string    `gorm:"column:pro_status_reason"`
	DateUpdt        *time.Time `gorm:"column:pro_date_updt"`
	DateIns         *time.Time `gorm:"column:pro_date_ins; default:now()"`
	DateDel         *time.Time `gorm:"column:pro_date_del"`
	CodeExtUse      int        `gorm:"column:use_code_ext"`
	CodeIntUse      *int       `gorm:"column:use_code_int"`
}

// TableName schema and table references
func (r *ResProduct) TableName() string {
	return "resources.res_product"
}

//ResProductFill fill model by request
func (r *ResProduct) ResProductFill(req requests.ResProduct) error {

	cpfCnpj, err := strconv.Atoi(req.CostumerCpfCnpj)
	if err != nil {
		return err
	}
	r.CostumerCpfCnpj = cpfCnpj
	r.CostumerEmail = req.CostumerEmail
	r.CostumerType = req.CostumerType

	return nil
}
