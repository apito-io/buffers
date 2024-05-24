package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
)

func main() {
	dsn := "postgres://postgres:Apitorocks247$@35.223.170.154:5432/apito_system?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := runMigrations(db); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	fmt.Println("Migrations ran successfully")
}

func runMigrations(db *bun.DB) error {
	ctx := context.Background()

	models := []interface{}{
		(*protobuff.SystemUser)(nil),
		(*protobuff.Project)(nil),
		(*protobuff.ProjectSchema)(nil),
		(*protobuff.PluginDetails)(nil),
		(*protobuff.AddOnsDetails)(nil),
		(*protobuff.APIToken)(nil),
		//(*protobuff.Role)(nil),
		(*protobuff.DriverCredentials)(nil),
		(*protobuff.UsagesTracking)(nil),
		(*protobuff.TeamMembers)(nil),
		(*protobuff.SystemMessage)(nil),
		(*protobuff.Workspace)(nil),

		(*protobuff.ModelType)(nil),
		(*protobuff.CloudFunction)(nil),
	}

	for _, model := range models {
		_, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
