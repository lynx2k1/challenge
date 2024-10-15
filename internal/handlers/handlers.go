package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

// ContactForm define a estrutura do formulário de contato.
type ContactForm struct {
    Name              string `json:"name" form:"name" binding:"required"`
    Email             string `json:"email" form:"email" binding:"required,email"`
    Comment           string `json:"comment" form:"comment" binding:"required"`
    CaptchaResponse   string `json:"g-recaptcha-response" form:"g-recaptcha-response" binding:"required"`
}

// HandleContactForm trata a solicitação do formulário de contato.
func HandleContactForm(c echo.Context) error {
    var form ContactForm

    // Vincula os dados da solicitação ao struct
    if err := c.Bind(&form); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
    }

    // Aqui você pode adicionar a lógica para verificar o reCAPTCHA,
    // enviar o e-mail, salvar em banco de dados, etc.

    // Exemplo: retorne um sucesso
    return c.NoContent(http.StatusCreated) // Retorna 201 em caso de sucesso
}
