package main

import (
	"coffeewithegg/apps/adam/app/service"
	"coffeewithegg/apps/adam/graph"
	"coffeewithegg/apps/adam/graph/generated"
	"coffeewithegg/apps/adam/migrations"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CWE")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(`config.yml`)
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read config file, using ENV instead")
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
	viper.SetDefault(`database.slow-threshold`, 1)
	viper.SetDefault(`database.show-sql`, false)

	viper.SetDefault(`cache.host`, `localhost`)
	viper.SetDefault(`cache.port`, `6379`)
	viper.SetDefault(`cache.pass`, "password")
	viper.SetDefault(`cache.db`, 0)
}

func initializeDB() *gorm.DB {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", dbUser, dbPass, dbHost, dbPort, dbName)

	slowThreshold := viper.GetInt64(`database.slow-threshold`)
	showSql := viper.GetBool(`database.show-sql`)
	logLevel := logger.Error
	if showSql {
		logLevel = logger.Info
	}

	sqlLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Duration(slowThreshold) * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logLevel,
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

func initializeCache() *redis.Client {
	host := viper.GetString(`cache.host`)
	port := viper.GetString(`cache.port`)
	password := viper.GetString(`cache.pass`)
	db := viper.GetInt(`cache.db`)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: password,
		DB:       db,
	})

	err := container.Singleton(func() *redis.Client {
		return rdb
	})
	if err != nil {
		log.Fatal(err)
	}

	return rdb
}

func startServer() {
	e := echo.New()

	allowOrigins := viper.GetStringSlice("server.cors")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowOrigins,
	}))

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
	initializeCache()
	migrate(db)

	err := service.InitServicesContainer()
	if err != nil {
		log.Fatal(err)
	}

	startServer()
}
