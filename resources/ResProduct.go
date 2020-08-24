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
	CostumerCpfCnpj string     `json:"costumermid_cpf_cnpj"`
	CostumerEmail   string     `json:"costumer_email"`
	CostumerCpfCnj  string     `json:"costumer_cpf_cnj"`
	CostumerType    int        `json:"costumer_type"`
	Status          string     `json:"status"`
	StatusReason    *string    `json:"status_reason"`
	DateUpdt        *time.Time `json:"date_updt"`
}

//ResProductResource fill resource by model
func (r *ResProduct) ResProductResource(mod models.ResProduct) {
	r.Code = &mod.Code

	cpfCnpj := mod.CostumerCpfCnpj
	if mod.CostumerType == 0 {
		r.CostumerCpfCnpj = fmt.Sprintf("%011s", strconv.Itoa(int(cpfCnpj)))
	} else if mod.CostumerType == 1 {
		r.CostumerCpfCnpj = fmt.Sprintf("%014s", strconv.Itoa(int(cpfCnpj)))

	}
	r.CostumerType = mod.CostumerType
	if mod.Status == 0 {
		r.Status = "Pendente"
	} else if mod.Status == 1 {
		r.Status = "Aprovado"
	} else if mod.Status == 2 {
		r.Status = "Rejeitado"
	}

	r.StatusReason = mod.StatusReason
	r.DateUpdt = mod.DateUpdt

}
