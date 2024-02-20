package cmd

import (
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/missingstudio/studio/backend/internal/storage/postgres/migrations"
	"github.com/missingstudio/studio/backend/pkg/database"
)

func RunMigrations(logger *slog.Logger, config database.Config) error {
	m, err := getDatabaseMigrationInstance(config)
	if err != nil {
		return err
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	migrationVer, dirty, err := m.Version()
	logger.Info("db migrated", "version", migrationVer, "dirty", dirty)
	return err
}

func RunRollback(logger *slog.Logger, config database.Config) error {
	m, err := getDatabaseMigrationInstance(config)
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	migrationVer, dirty, err := m.Version()
	logger.Info("db rolled back", "version", migrationVer, "dirty", dirty)
	return err
}

func getDatabaseMigrationInstance(config database.Config) (*migrate.Migrate, error) {
	fs := migrations.MigrationFs
	resourcePath := migrations.ResourcePath
	src, err := iofs.New(fs, resourcePath)
	if err != nil {
		return &migrate.Migrate{}, fmt.Errorf("db migrator: %v", err)
	}
	return migrate.NewWithSourceInstance("iofs", src, config.URL)
}
