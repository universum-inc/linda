package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/universum-inc/linda/internal/account"
	"gopkg.in/natefinch/lumberjack.v2"
)

const portDefault = 1455
const defaultLogLevel = "info"
const configDefault = "config.yaml"

func main() {

	viper.SetDefault("config.path", configDefault)
	_ = viper.BindEnv("config.path", "CONFIG_PATH")
	configPath := viper.GetString("config.path")
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetDefault("logging.level", defaultLogLevel)
	_ = viper.BindEnv("logging.level", "log-level")
	loggingLevel := viper.GetString("logging.level")
	fileLogger := &lumberjack.Logger{
		Filename:   viper.GetString("logging.file"),
		MaxSize:    viper.GetInt("logging.max-file-size"),
		MaxBackups: viper.GetInt("logging.max-backups"),
		MaxAge:     viper.GetInt("logging.max-age"),
		Compress:   viper.GetBool("logging.compress-rotated-log"),
	}
	level, err := zerolog.ParseLevel(loggingLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
	zerolog.TimeFieldFormat = viper.GetString("logging.time-format")

	standardOutput := viper.GetString("logging.standard-output")
	if standardOutput == "stderr" {
		log.Logger = log.Output(zerolog.MultiLevelWriter(os.Stderr, fileLogger))
	} else if standardOutput == "stdout" {
		log.Logger = log.Output(zerolog.MultiLevelWriter(os.Stdout, fileLogger))
	} else {
		log.Logger = log.Output(zerolog.MultiLevelWriter(fileLogger))
	}

	_ = viper.BindEnv("database.dataSourceName", "DATABASE_DSN")
	dataSourceName := viper.GetString("database.dataSourceName")

	repo, err := account.NewRepository(dataSourceName)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create repository")
	}

	service := account.NewService(repo)

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/accounts", func(c fiber.Ctx) error {

		limitParam := c.Query("limit", "10")
		limit, errParse := strconv.ParseInt(limitParam, 10, 64)
		if errParse != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid limit parameter",
			})
		}

		accounts, errGetAccounts := service.GetAccounts(limit)
		if errGetAccounts != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": errGetAccounts.Error(),
			})
		}
		return c.JSON(accounts)
	})

	viper.SetDefault("server.http-port", portDefault)
	_ = viper.BindEnv("server.http-port", "http_port")

	port := viper.GetInt("server.http-port")

	err = app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed")
	}
}
