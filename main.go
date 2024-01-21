package main

import (
	"ecommerce/database"
	"ecommerce/models"
	"ecommerce/routes"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("System Error: ", err)
	}

	// MySQL database connection
	db, err := database.InitializesConnection()
	if err != nil {
		log.Fatal("Database Error: ", err)
	}

	// AutoMigrate will create tables based on the provided struct models
	if err := MigrateModels(db); err != nil {
		log.Fatal("Database Error: ", err)
	}
	log.Info("all migration is run")
	// Start the Echo server and define routes
	routes.StartServer(db)
}

// Function to auto-migrate models
func MigrateModels(db *gorm.DB) error {

	// db.Migrator().DropTable(
	// 	&models.User{},
	// 	&models.Session{},
	// 	&models.Product{},
	// 	&models.Category{},
	// 	&models.Order{},
	// 	&models.OrderDetail{},
	// 	&models.Review{},
	// 	&models.ShoppingCart{},
	// 	&models.Payment{},
	// 	&models.Address{},
	// 	&models.Coupon{},
	// )

	return db.AutoMigrate(
		&models.User{},
		&models.Session{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Review{},
		&models.ShoppingCart{},
		&models.Payment{},
		&models.Address{},
		&models.Coupon{},
	)
}
