package requests

import "time"

// ResProduct ...
type ResProduct struct {
	Code            *int       `json:"code"`
	CostumermidCnpj int        `json:"costumermid_cnpj" validate:"required"`
	CostumerEmail   string     `json:"costumer_email" validate:"required,max=255"`
	CostumerCpfCnj  string     `json:"costumer_cpf_cnj" validate:"required"`
	CostumerType    int        `json:"costumer_type" validate:"required,oneof=0 1"` //0 cpf 1 cnpj
	Status          int        `json:"status"`
	StatusReason    *string    `json:"reason"`
	DateUpdt        *time.Time `json:"date_updt"`
	DateIns         *time.Time `json:"date_ins" validate:"required"`
	DateDel         *time.Time `json:"date_del"`
	CodeExtUse      int        `json:"code_ext_use"`
	CodeIntUse      *int       `json:"code_int_use"`
}
