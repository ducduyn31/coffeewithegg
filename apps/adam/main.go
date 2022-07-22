package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"time"

	_ "github.com/99designs/gqlgen"

	_projectHttpDelivery "coffeewithegg/apps/adam/project/delivery/http"
)

func init() {
	viper.SetConfigFile(`config.yaml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func initializeDB() {}

func closeDB() {

}

func main() {
	initializeDB()

	defer func() {
		closeDB()
	}()

	e := echo.New()

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	_projectHttpDelivery.NewProjectHandler(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
