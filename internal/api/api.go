package api

import (
	"api/internal/config"
	db "api/internal/database"
	e "api/internal/error"
	"errors"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type API struct {
	cfg *config.Config
	db  *db.Database
}

func New() (*API, error) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("initializing API")

	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		return nil, err
	}

	return &API{cfg, db}, nil
}

func (api *API) Close() {
	slog.Debug("closing API")
	api.db.Close()
}

func (api *API) Run() error {
	mux := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ErrorHandler:          errorHandler,
	})

	mux.Use(logger.New())
	mux.Use(recover.New())

	api.SetRoutes(mux)

	slog.Info("listening", "address", api.cfg.LISTEN_ADDRESS)
	return mux.Listen(api.cfg.LISTEN_ADDRESS)
}

func (api *API) SeedWithFakeData() error {
	return api.db.SeedWithFakeData()
}

func errorHandler(c *fiber.Ctx, err error) error {
	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		return c.SendStatus(fiberErr.Code)
	}

	var apiErr e.Error
	if errors.As(err, &apiErr) {
		slog.Info("error", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": apiErr})
	}

	slog.Error("internal server error", "err", err)
	return c.SendStatus(fiber.StatusInternalServerError)
}
