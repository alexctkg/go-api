package utils

import "github.com/go-playground/validator/v10"

// Valid an request model
func Valid(obj interface{}) []string {
	validate := validator.New()
	err := validate.Struct(obj)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return []string{err.Error()}
		}
		tam := len(err.(validator.ValidationErrors))
		errs := make([]string, tam)
		for i, err := range err.(validator.ValidationErrors) {
			errs[i] = validMessage(err.Field(), err.Tag(), err.Param())
		}
		return errs
	}
	return nil
}

//validMessage
func validMessage(field string, tag string, param string) string {
	switch tag {
	case "min":
		return "O campo " + field + " deve conter no mínimo " + param + " caracter(es)"
	case "max":
		return "O campo " + field + " deve conter no máximo " + param + " caracter(es)"
	case "required":
		return "O campo " + field + " é obrigatório"
	case "oneof":
		return "O campo " + field + " deve conter um dos seguintes valores:" + param
	case "eqfield":
		return "O campo " + field + " deve ser igual ao campo " + param
	default:
		return "O campo " + field + " falhou na validação " + tag + " " + param
	}
}
