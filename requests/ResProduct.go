package requests

import "time"

// ResPorduct ...
type ResPorduct struct {
	Code            int        `json:"code" validate:"required"`
	CostumermidCnpj int        `json:"costumermid_cnpj" validate:"required"`
	CostumerEmail   string     `json:"costumer_email" validate:"required,max=255"`
	CostumerCpfCnj  int        `json:"costumer_cpf_cnj" validate:"required"`
	Status          int        `json:"status" validate:"required"`
	StatusReason    *string    `json:"status_reason"`
	DateUpdt        *time.Time `json:"date_updt"`
	DateIns         *time.Time `json:"date_ins" validate:"required"`
	DateDel         *time.Time `json:"date_del"`
	CodeExtUse      int        `json:"code_ext_use" validate:"required"`
	CodeIntUse      *int       `json:"code_int_use"`
}
