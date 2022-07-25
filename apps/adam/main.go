package main

import (
	"coffeewithegg/apps/adam/app/service"
	"coffeewithegg/apps/adam/graph"
	"coffeewithegg/apps/adam/graph/generated"
	"coffeewithegg/apps/adam/migrations"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/fsnotify/fsnotify"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

func init() {

	viper.NewWithOptions(
		viper.EnvKeyReplacer(strings.NewReplacer("_", ".")),
	).AutomaticEnv()

	viper.SetConfigFile(`config.yml`)
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed: ", e.Op.String())
	})

	// Default Values
	viper.SetDefault(`database.host`, `localhost`)
	viper.SetDefault(`database.port`, `5432`)
	viper.SetDefault(`database.user`, `postgres`)
	viper.SetDefault(`database.pass`, `password`)
	viper.SetDefault(`database.name`, `adam`)
	viper.SetDefault(`database.location`, `Asia/Ho_Chi_Minh`)
}

func initializeDB() *gorm.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", dbUser, dbPass, dbHost, dbPort, dbName)

	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		},
	)

	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: sqlLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = container.Singleton(func() *gorm.DB {
		return db
	})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func startServer() {
	e := echo.New()

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{}}),
	)

	e.POST("/graphql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func migrate(db *gorm.DB) {
	err := migrations.InitMigrationTable(db)
	if err != nil {
		log.Fatal("Failed to init migration table")
		return
	}

	err = migrations.DoMigration(db)
	if err != nil {
		log.Fatal("Failed to do migration")
		return
	}

	return
}

func main() {
	db := initializeDB()
	migrate(db)

	err := service.InitServicesContainer()
	if err != nil {
		log.Fatal(err)
	}

	startServer()
}
