package database

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gitlab.com/truemeds-dev-team/truemeds-dev-doctor/truemeds-dev-service/doctorportal-auth-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	instance *gorm.DB
	once     sync.Once
	initErr  error
)

// Get singleton database connection
func ConnectDb(cfg *config.Env) (*gorm.DB, error) {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			cfg.DB.Host, cfg.DB.User, cfg.DB.Password,
			cfg.DB.DBName, cfg.DB.Port,
		)

		var logLevel logger.LogLevel
		if cfg.DB.Logging {
			logLevel = logger.Info // all logs
		} else {
			logLevel = logger.Silent
		}

		gormLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second, // Log slower queries
				LogLevel:                  logLevel,    // Control log level
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: gormLogger,
		})
		if err != nil {
			initErr = fmt.Errorf("failed to connect to database: %w", err)
			return
		}

		sqlDB, err := db.DB()
		if err != nil {
			initErr = fmt.Errorf("failed to get underlying sql.DB: %w", err)
			return
		}

		sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConnection)
		sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConnection)
		sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.ConnectionLifeTimeMinute) * time.Minute)

		instance = db
	})

	return instance, initErr
}
