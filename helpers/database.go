package helpers

import (
	"backend-test/internal/models"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupPostgreSQL() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "5432"), GetEnv("DB_USER", ""), GetEnv("DB_PASSWORD", ""), GetEnv("DB_NAME", ""))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	logrus.Info("Successfully connect to database..")
	DB.AutoMigrate(models.User{}, models.UserSession{}, models.Inventory{}, models.Ingredient{}, models.Recipe{})
	seedInventoryData()
}

func seedInventoryData() {
	var count int64
	if err := DB.Model(&models.Inventory{}).Count(&count).Error; err != nil {
		log.Fatal("Failed to check inventory table: ", err)
	}

	if count == 0 {
		logrus.Info("Seeding inventory data...")

		inventoryData := []models.Inventory{
			{Item: "Aren Sugar", Qty: 1, Uom: "Kg", PricePerQty: 60000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Item: "Plastic Cup", Qty: 10, Uom: "Pcs", PricePerQty: 5000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Item: "Coffe Bean", Qty: 1, Uom: "Kg", PricePerQty: 100000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Item: "Mineral Water", Qty: 1, Uom: "Liter", PricePerQty: 5000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Item: "Ice Cub", Qty: 1, Uom: "Kg", PricePerQty: 15000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Item: "Milk", Qty: 1, Uom: "Liter", PricePerQty: 30000, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		if err := DB.Save(&inventoryData).Error; err != nil {
			log.Fatal("Failed to seed inventory data: ", err)
		}

		logrus.Info("Inventory data seeded successfully!")
	} else {
		logrus.Info("Inventory table already has data, skipping seeding.")
	}
}
