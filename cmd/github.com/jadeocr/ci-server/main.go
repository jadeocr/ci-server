package main

import (
  "fmt"
  "log"
  "os/exec"
	"net/http"
	"github.com/labstack/echo/v4"
)

func dockerBuild() {
  cmd := exec.Command("docker-compose", "build")
  stdout, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(stdout)
}

func dockerUp() {
  cmd := exec.Command("docker-compose", "up", "-d")
  stdout, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(stdout)
}

func githubController(c echo.Context) error {
  dockerBuild()
  dockerUp()
  return c.String(http.StatusOK, "ok")
}

func helloController(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", helloController)
  e.POST("/github", githubController)

	e.Logger.Fatal(e.Start(":4000"))
}
