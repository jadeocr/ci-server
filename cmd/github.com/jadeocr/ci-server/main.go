package main

import (
  "fmt"
  "log"
  "os/exec"
	"net/http"
	"github.com/labstack/echo/v4"
)

func githubController(c echo.Context) error {
  cmd := exec.Command("/bin/sh", "-c", "git pull; docker-compose build; docker-compose up -d; docker image prune -f")
  err := cmd.Run()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("/github called")
  return c.String(http.StatusOK, "ok")
}

func helloController(c echo.Context) error { // for testing
  fmt.Println("/ called")
  return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", helloController)
  e.POST("/github", githubController)

	e.Logger.Fatal(e.Start(":4000"))
}
