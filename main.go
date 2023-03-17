package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)


type Result struct {
	Start string `json:"start:string"`
	End string `json:"end:string"`
}

func main() {

	defer func() {
		log.Println("this is defer function")
	}()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		start := time.Now()
		time.Sleep(5 * time.Second)
		end := time.Now()
		const layout = time.RFC3339
		return c.JSON(
			http.StatusOK,
			Result{Start: start.Format(layout), End: end.Format(layout)})
	})
	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()


	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	log.Println("last line")
}