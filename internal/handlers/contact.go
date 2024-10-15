package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/lynx2k1/challenge/internal/mail"
    "github.com/lynx2k1/challenge/internal/recaptcha"
    "github.com/lynx2k1/challenge/internal/models"
)

func HandleContactForm(c echo.Context) error {
    var contact models.ContactForm
    if err := c.Bind(&contact); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
    }

    // Validação do ReCaptcha
    validCaptcha, err := recaptcha.ValidateCaptcha(contact.RecaptchaResponse)
    if !validCaptcha || err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{
            "type":     "about:blank",
            "title":    "UnauthorizedError",
            "detail":   "The captcha is incorrect!",
            "instance": "/contact",
        })
    }

    // Validação dos dados do contato
    if contact.Name == "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "type":     "about:blank",
            "title":    "BadRequestError",
            "detail":   "The name is empty",
            "instance": "/contact",
        })
    }

    if !contact.IsValidEmail() {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "type":     "about:blank",
            "title":    "BadRequestError",
            "detail":   "The email is invalid",
            "instance": "/contact",
        })
    }

    // Envio de email
    err = mail.SendMail(contact)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "type":     "about:blank",
            "title":    "InternalServerError",
            "detail":   "Could not send the email.",
            "instance": "/contact",
        })
    }

    return c.NoContent(http.StatusCreated)
}
