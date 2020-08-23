package models

import (
	"strconv"
	"tdez/requests"
	"time"
)

// ResProduct ...
type ResProduct struct {
	Code            int        `gorm:"column:pro_code; primary_key:true"`
	CostumermidCnpj int        `gorm:"column:pro_costumermid_cnpj"`
	CostumerEmail   string     `gorm:"column:pro_costumer_email"`
	CostumerCpfCnj  int        `gorm:"column:pro_costumer_cpf_cnj"`
	CostumerType    int        `gorm:"column:pro_costumer_type"`
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
	r.CostumermidCnpj = req.CostumermidCnpj
	r.CostumerEmail = req.CostumerEmail

	cpfCnpj, err := strconv.Atoi(req.CostumerCpfCnj)
	if err != nil {
		return err
	}
	r.CostumerCpfCnj = cpfCnpj
	r.Status = req.Status
	r.StatusReason = req.StatusReason
	r.DateUpdt = req.DateUpdt
	r.DateIns = req.DateIns
	r.DateDel = req.DateDel
	r.CodeExtUse = req.CodeExtUse
	r.CodeIntUse = req.CodeIntUse

	return nil
}
