package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// 現在時刻を取得
func getNow() string {
	var loc, _ = time.LoadLocation("Asia/Tokyo")
	var now = time.Now().In(loc).Format("2006-01-02 15:04:05")
	return now
}

// Handler
func hello(c *echo.Context) error {
	var content struct {
		Response  string `json:"response"`
		Timestamp string `json:"timestamp"`
	}
	content.Response = "Hello, World!"
	content.Timestamp = getNow()
	return c.JSON(http.StatusOK, &content)
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/", hello)

	// Start server
	e.Run(":1323")
}
