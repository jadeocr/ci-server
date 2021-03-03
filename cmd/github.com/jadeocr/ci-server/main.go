package main

import (
  "fmt"
  "log"
  "os/exec"
	"net/http"
	"github.com/labstack/echo/v4"
)

func gitPull() {
  cmd := exec.Command("git", "pull")
  stdout, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(stdout)
}

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

func dockerClean() {
  cmd := exec.Command("docker", "image", "prune", "-f")
  stdout, err := cmd.Output()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(stdout)
}

func githubController(c echo.Context) error {
  gitPull()
  dockerBuild()
  dockerUp()
  fmt.Println("/github called")
  return c.String(http.StatusOK, "ok")
}

func helloController(c echo.Context) error {
  fmt.Println("/ called")
  return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()
	e.GET("/", helloController)
  e.POST("/github", githubController)

	e.Logger.Fatal(e.Start(":4000"))
}
