package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/zykunov/RESTCurrencyChange/internal/app/handlers"
)

func main() {

	e := echo.New()

	//подгрузка конфига
	if err := godotenv.Load("./configs/.env"); err != nil {
		e.Logger.Info("No .env file found")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	e.Logger.SetLevel(log.INFO)

	e.POST("/change", func(c echo.Context) error {
		return handlers.Change(c)
	})

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Старт сервера
	go func() {
		if err := e.Start(host + ":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Ждем прерывающего сигнала
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Graceful shutdown
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
