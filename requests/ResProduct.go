package requests

// ResProduct - external app to create issue
type ResProduct struct {
	Code            *int   `json:"code"`
	CostumerCpfCnpj string `json:"costumermid_cpf_cnpj" validate:"required"` //0 cpf 1 cnpj
	CostumerEmail   string `json:"costumermid_email" validate:"required,max=255"`
	CostumerType    int    `json:"costumermid_type" validate:"oneof=0 1"`
}

// ResProductResponse - super user to aprove or decline issue
type ResProductResponse struct {
	Code   *int    `json:"code"  validate:"required"`
	Reason *string `json:"reason" validate:"required"`
}
