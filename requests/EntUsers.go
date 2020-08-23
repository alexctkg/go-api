package requests

// EntUsers ...
type EntUsers struct {
	Code            int     `json:"code"`
	Email           string  `json:"email" validate:"required,max=255"`
	Cnpj            *string `json:"cnpj"`
	RazaoSocial     *string `json:"razao_social" validate:"omitempty,max=60"`
	Password        string  `json:"password" validate:"required,max=255"`
	ConfirmPassword *string `json:"confirm_password" validate:"eqfield=Password,max=255"`
	Token           *string `json:"token"`
	Type            int     `json:"type"`
}
