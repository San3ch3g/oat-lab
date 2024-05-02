package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"oat-lab-module/internal/pkg/models"
	"oat-lab-module/internal/utils/config"
	"strings"
)

type Storage struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func buildDSN(cfg *config.Config) string {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)
	return dsn
}

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MustNewPostgresDB(cfg *config.Config) *gorm.DB {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}

	err = createDataType(db, "sex_type", []string{string(models.MaleSex), string(models.WomanSex)})
	if err != nil {
		fmt.Println("sex_type is already exist")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("table 'users' is already exist")
	}
	err = db.AutoMigrate(&models.CodeForEmail{})
	if err != nil {
		fmt.Println("table 'code_for_email' is already exist")
	}
	err = db.AutoMigrate(&models.Item{})
	if err != nil {
		fmt.Println("table 'items' is already exist")
	}
	err = db.AutoMigrate(&models.News{})
	if err != nil {
		fmt.Println("table 'news' is already exist")
	}
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		fmt.Println("table 'orders' is already exist")
	}
	return db
}

func createDataType(db *gorm.DB, dataType string, values []string) error {
	valuesStr := strings.Join(values, "', '")
	query := fmt.Sprintf("CREATE TYPE %s AS ENUM ('%s')", dataType, valuesStr)
	err := db.Exec(query).Error
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			fmt.Printf("Data type %s already exists\n", dataType)
			return nil
		}
		return err
	}
	fmt.Printf("Data type %s created successfully\n", dataType)
	return nil
}

func (s *Storage) CheckUser(email string) error {
	var foundUser models.User
	err := s.db.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return nil
}
