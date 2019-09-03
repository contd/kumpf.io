package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var appInfo AppInfo

// AppInfo is the application information to show
type AppInfo struct {
	Name    string    `json:"name"`
	Version string    `json:"version"`
	Date    time.Time `json:"date"`
}

func main() {
	var PORT = os.Getenv("SERVER_PORT")
	if PORT == "" {
		PORT = ":7777"
	}

	// AppInfo
	appInfo = AppInfo{
		Name:    "kumpf.io",
		Version: "0.1.0",
		Date:    time.Now(),
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://kumpf.io", "http://localhost"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Static
	//e.File("/", "../app/build/index.html")

	// Routes
	e.GET("/", indexHandler)

	// Start server
	e.Logger.Fatal(e.Start(PORT))
}

// indexHandler
func indexHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, appInfo)
}
