package resources

import (
	"fmt"
	"strconv"
	"tdez/models"
	"time"
)

// ResProduct ...
type ResProduct struct {
	Code            *int       `json:"code"`
	CostumermidCnpj int        `json:"costumermid_cnpj"`
	CostumerEmail   string     `json:"costumer_email"`
	CostumerCpfCnj  string     `json:"costumer_cpf_cnj"`
	CostumerType    int        `json:"costumer_type"`
	Status          int        `json:"status"`
	StatusReason    *string    `json:"status_reason"`
	DateUpdt        *time.Time `json:"date_updt"`
	DateIns         *time.Time `json:"date_ins"`
	DateDel         *time.Time `json:"date_del"`
	CodeExtUse      int        `json:"code_ext_use"`
	CodeIntUse      *int       `json:"code_int_use"`
}

//ResProductResource fill resource by model
func (r *ResProduct) ResProductResource(mod models.ResProduct) {
	r.Code = &mod.Code
	r.CostumermidCnpj = mod.CostumermidCnpj

	cpfCnpj := mod.CostumerCpfCnj
	if mod.CostumerType == 0 {
		r.CostumerCpfCnj = fmt.Sprintf("%011s", strconv.Itoa(int(cpfCnpj)))
	} else if mod.CostumerType == 1 {
		r.CostumerCpfCnj = fmt.Sprintf("%015s", strconv.Itoa(int(cpfCnpj)))

	}
	r.CostumerType = mod.CostumerType
	r.Status = mod.Status
	r.StatusReason = mod.StatusReason
	r.DateUpdt = mod.DateUpdt
	r.DateIns = mod.DateIns
	r.DateDel = mod.DateDel
	r.CodeExtUse = mod.CodeExtUse
	r.CodeIntUse = mod.CodeIntUse
}
