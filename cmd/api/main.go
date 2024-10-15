package main

import (
    "log"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/joho/godotenv"
    "github.com/lynx2k1/challenge/internal/handlers"
)

func main() {
    godotenv.Load() // Carregar variáveis de ambiente do .env
    e := echo.New()

    e.POST("/contact", handlers.HandleContactForm) // Rota para o formulário

    port := os.Getenv("PORT")
    if port == "" {
        port = "80"
    }
    log.Fatal(e.Start(":" + port))
}
