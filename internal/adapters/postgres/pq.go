package postgres

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
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
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error fetching wd: %v", err)
	}
	jsonPath := filepath.Join(wd, "internal", "adapters", "postgres", "seeds")
	initDb(db, jsonPath)
	return db
}

func autoMigrate(db *gorm.DB) {
	db.Debug().AutoMigrate(&core.Comment{}, &core.Post{})
}

func initDb(db *gorm.DB, jsonPath string) {
	seedTags(db, filepath.Join(jsonPath, "tags.json"))
	seedPosts(db, filepath.Join(jsonPath, "posts.json"))
}

func seedTags(db *gorm.DB, tagPath string) {
	file, err := os.Open(tagPath)
	if err != nil {
		log.Fatalf("error opening tag's seed file: %v", err)
	}
	defer file.Close()

	var tags []core.Tag
	if err := json.NewDecoder(file).Decode(&tags); err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}

	if err := db.CreateInBatches(tags, 5).Error; err != nil {
		log.Fatalf("error saving tags to db: %v", err)
	}

	log.Println("Tags seeded!")
}

func seedPosts(db *gorm.DB, postPath string) {
	file, err := os.Open(postPath)
	if err != nil {
		log.Fatalf("error opening post's seed file: %v", err)
	}
	defer file.Close()

	var posts []core.Post
	if err := json.NewDecoder(file).Decode(&posts); err != nil {
		log.Fatalf("error decoding JSON: %v", err)
	}

	if err := db.CreateInBatches(posts, 5).Error; err != nil {
		log.Fatalf("error saving posts to db: %v", err)
	}

	log.Println("Posts seeded!")
}
