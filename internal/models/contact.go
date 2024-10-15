package models

import "regexp"

type ContactForm struct {
    RecaptchaResponse string `json:"g-recaptcha-response" validate:"required"`
    Name              string `json:"name" validate:"required"`
    Email             string `json:"mail" validate:"required,email"`
    Comment           string `json:"comment" validate:"required"`
}

// Função para validar email usando regex
func (cf *ContactForm) IsValidEmail() bool {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
    return re.MatchString(cf.Email)
}
