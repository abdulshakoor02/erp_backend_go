package dbAdapter

import (
	"fmt"
	// "log"
	// "os"
	// "time"

	"github.com/abdul/erp_backend/config"
	"github.com/abdul/erp_backend/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// logging "gorm.io/gorm/logger"
)

var DB *gorm.DB

func DbConnect() *gorm.DB {
	// newLogger := logging.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // Output to standard output
	// 	logging.Config{
	// 		SlowThreshold:             time.Second,  // Log slow SQL queries over a second
	// 		LogLevel:                  logging.Info, // Log level set to Info to log all queries
	// 		IgnoreRecordNotFoundError: true,         // Ignore RecordNotFoundError logs
	// 		Colorful:                  false,        // Disable color output (optional)
	// 	},
	// )
	log := logger.Logger
	config.LoadEnv()

	log.Info().Msgf(config.DB_NAME)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:sh@k00oor@tcp(127.0.0.1:3306)/building_management?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v ",
		config.POSTGRES_USER,
		config.POSTGRES_PASSWORD,
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.DB_NAME,
	)
	// dsn := "abdul:sh@k00oor@tcp(172.17.0.1:5432)/erp_backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		fmt.Print(err)
		log.Fatal().Err(err).Msgf("failed to load env")
	}
	// db.Debug()
	DB = db
	pgdb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to initiate a database instance :%v", err)
	}
	pgdb.Ping()
	pgdb.SetMaxIdleConns(10)
	pgdb.SetMaxOpenConns(100)
	log.Info().Msg("Postgres DB connection established")
	return db
}
