package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"spy-cat/internal/configs"
	"spy-cat/internal/database"
	migrate "spy-cat/internal/database/migration"
	handlers "spy-cat/internal/handler"
	"spy-cat/internal/repository"
	"spy-cat/internal/router"
	"spy-cat/internal/service"
)

func RunApi() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	cfg, err := configs.Load()
	if err != nil {
		return fmt.Errorf("can not load config: %e", err)
	}

	db, err := database.NewPostgreSQLDB(cfg)
	if err != nil {
		return fmt.Errorf("could not connect to the database: %w", err)
	}
	defer db.Close() // close the connection to the database when the work is completed

	if err = migrate.RunMigrations(db); err != nil {
		return fmt.Errorf("could not migrate database: %w", err)
	}

	catRepo := repository.NewCatRepository(db)
	catService := service.NewCatService(catRepo)
	missionRepo := repository.NewMissionRepository(db)
	missionService := service.NewMissionService(missionRepo)

	catHandler := handlers.NewCatHandler(catService)
	missionHandler := handlers.NewMissionHandler(missionService)

	router.NewRouter(e, catHandler, missionHandler)
	log.Println(fmt.Sprintf("Server is running on port %d", cfg.Api.Port))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.Api.Port)))

	return nil

}
