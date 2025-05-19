package postgres

import (
	"log"
	"os"
	"time"

	"github.com/victorsvart/vbi/internal/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func Connect() *gorm.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN var env is not set")
		return nil
	}

	config := &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,
		},
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return nil
	}
	// InitData(db)

	// setup verbose logger after data init because the seeds polute too much :/
	defer func() {
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)

		db.Config.Logger = newLogger
	}()

	autoMigrate(db)
	return db
}

func autoMigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&core.Comment{}, &core.Post{})
}
