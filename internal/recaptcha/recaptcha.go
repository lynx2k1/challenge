package recaptcha

import (
    "encoding/json"
    "net/http"
    "net/url"
)

const secretKey = "your-secret-key"

// Validate reCAPTCHA
func Validate(recaptchaResponse string) (bool, error) {
    resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify",
        url.Values{"secret": {secretKey}, "response": {recaptchaResponse}})

    if err != nil {
        return false, err
    }
    defer resp.Body.Close()

    var response struct {
        Success bool `json:"success"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return false, err
    }

    return response.Success, nil
}
