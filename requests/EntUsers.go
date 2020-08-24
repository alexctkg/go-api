package requests

// EntUsersStore struct to store a new user (super or external)
type EntSuperUsersStore struct {
	Email           string  `json:"email" validate:"required,max=255"`
	Cnpj            *string `json:"cnpj"`
	RazaoSocial     *string `json:"razao_social"`
	Password        string  `json:"password" validate:"required,min=6,max=255"`
	ConfirmPassword *string `json:"confirm_password" validate:"eqfield=Password,max=255"`
	Type            int     `json:"-"`
}

// EntUsersStore struct to store a new user (super or external)
type EntExternalUsersStore struct {
	Email           string  `json:"email" validate:"required,max=255"`
	Cnpj            *string `json:"cnpj" validate:"required"`
	RazaoSocial     *string `json:"razao_social" validate:"required,max=60"`
	Password        string  `json:"password" validate:"required,min=6,max=255"`
	ConfirmPassword *string `json:"confirm_password" validate:"eqfield=Password,max=255"`
	Type            int     `json:"-"`
}

// EntUsersLogin struct to login
type EntUsersLogin struct {
	Email    string `json:"email" validate:"required,max=255"`
	Password string `json:"password" validate:"required,max=255"`
}
