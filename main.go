package main

import (
	"github.com/klovercloud-ci-cd/light-house-query/api"
	"github.com/klovercloud-ci-cd/light-house-query/config"
	_ "github.com/klovercloud-ci-cd/light-house-query/docs"
)

// @title Klovercloud-ci-light-house-query API
// @description Klovercloud-light-house-query API
func main() {
	e := config.New()
	api.Routes(e)
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
