package orm

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Importing for configuration only
	"github.com/lmuench/skynet/internal/pkg/services/types"
	"github.com/qor/admin"
)

// InitPostgresDev automigrates gorm models and returns DB connection pointer
func InitPostgresDev() (*gorm.DB, *admin.Admin) {
	conf := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("SKYNET_DEV_DB_HOST"),
		os.Getenv("SKYNET_DEV_DB_PORT"),
		os.Getenv("SKYNET_DEV_DB_DBNAME"),
		os.Getenv("SKYNET_DEV_DB_USER"),
		os.Getenv("SKYNET_DEV_DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", conf)
	if err != nil {
		panic("failed to connect to database")
	}

	db.DropTableIfExists(
		&types.Unit{},
		&types.Prototype{},
	)
	db.AutoMigrate(
		&types.Unit{},
		&types.Prototype{},
	)

	adm := admin.New(&admin.AdminConfig{DB: db})
	adm.AddResource(&types.Unit{})
	adm.AddResource(&types.Prototype{})

	return db, adm
}

// InitPostgresTest drops tables, automigrates gorm models and returns DB connection pointer
func InitPostgresTest() *gorm.DB {
	conf := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		os.Getenv("SKYNET_TEST_DB_HOST"),
		os.Getenv("SKYNET_TEST_DB_PORT"),
		os.Getenv("SKYNET_TEST_DB_DBNAME"),
		os.Getenv("SKYNET_TEST_DB_USER"),
		os.Getenv("SKYNET_TEST_DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", conf)
	if err != nil {
		panic("failed to connect to database")
	}

	db.DropTableIfExists(
		&types.Unit{},
		&types.Prototype{},
	)
	db.AutoMigrate(
		&types.Unit{},
		&types.Prototype{},
	)

	return db
}
