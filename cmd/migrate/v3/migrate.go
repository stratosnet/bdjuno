package v5

import (
	"fmt"

	v3db "github.com/forbole/bdjuno/v4/database/migrate/v3"
	parse "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/forbole/juno/v5/database"
	"github.com/forbole/juno/v5/database/postgresql"
	loggingconfig "github.com/forbole/juno/v5/logging/config"
	nodeconfig "github.com/forbole/juno/v5/node/config"
	parserconfig "github.com/forbole/juno/v5/parser/config"
	"github.com/forbole/juno/v5/types/config"
)

// RunMigration runs the migrations to v5
func RunMigration(parseConfig *parse.Config) error {
	cfg, err := GetConfig()
	if err != nil {
		return fmt.Errorf("error while reading config: %s", err)
	}

	// Migrate the database
	err = migrateDb(cfg, parseConfig)
	if err != nil {
		return fmt.Errorf("error while migrating database: %s", err)
	}

	return nil
}

func migrateDb(cfg config.Config, parseConfig *parse.Config) error {
	// Build the codec
	encodingConfig := parseConfig.GetEncodingConfigBuilder()()

	// set runtime config
	config.Cfg = config.NewConfig(
		nodeconfig.Config{},
		config.ChainConfig{},
		cfg.Database,
		parserconfig.Config{},
		loggingconfig.Config{},
	)

	// Get the db
	databaseCtx := database.NewContext(cfg.Database, &encodingConfig, parseConfig.GetLogger())
	db, err := postgresql.Builder(databaseCtx)
	if err != nil {
		return fmt.Errorf("error while building the db: %s", err)
	}

	// Build the migrator and perform the migrations
	migrator := v3db.NewMigrator(db.(*postgresql.Database))
	return migrator.Migrate()
}
