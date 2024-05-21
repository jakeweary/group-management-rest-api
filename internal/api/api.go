package api

import (
	"api/internal/config"
	db "api/internal/database"
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	cfg config.Config
	db  db.Database
}

func New() API {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Info("loading config")
	cfg := config.Load()

	slog.Info("connecting to db")
	db := db.Connect(cfg.DATABASE_URL)

	return API{cfg, db}
}

func (api *API) Close() {
	api.db.Close()
}

func (api *API) Run() {
	mux := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})

	mux.Use(logger.New())
	mux.Use(recover.New())

	api.SetRoutes(mux)

	slog.Info("listening", "address", api.cfg.LISTEN_ADDRESS)
	err := mux.Listen(api.cfg.LISTEN_ADDRESS)
	if err != nil {
		log.Fatal(err)
	}
}

func (api *API) SeedWithFakeData() {
	api.db.SeedWithFakeData()
}

func errorHandler(c *fiber.Ctx, err error) error {
	var e *fiber.Error
	if errors.As(err, &e) {
		return c.SendStatus(e.Code)
	}

	var apiErr Error
	if errors.As(err, &apiErr) {
		slog.Info("api error", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": apiErr})
	}

	var dbErr db.Error
	if errors.As(err, &dbErr) {
		slog.Info("db error", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": dbErr})
	}

	slog.Error("internal server error", "err", err)
	return c.SendStatus(fiber.StatusInternalServerError)
}
