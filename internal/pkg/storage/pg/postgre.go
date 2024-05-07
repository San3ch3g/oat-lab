package pg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"oat-lab-module/internal/pkg/models"
	"oat-lab-module/internal/pkg/services"
	"oat-lab-module/internal/utils/config"
	"strings"
	"time"
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
	err = createDataType(db, "news_category_type", []string{string(models.Default), string(models.Covid), string(models.Complex), string(models.Popular)})
	if err != nil {
		fmt.Println("news_category_type is already exist")
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

func (s *Storage) isUserExist(email string) (bool, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *Storage) CheckUser(cfg config.Config, email string) (bool, error) {
	var foundUser models.User
	err := s.db.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			code := services.GenerateCodeForEmail()
			err := services.SendCodeToEmailService(cfg, code, email)
			if err != nil {
				return false, err
			}
			err = s.db.Create(&models.CodeForEmail{Email: email, Code: code, CreatedAt: time.Now().Format(time.RFC3339)}).Error
			if err != nil {
				return false, err
			}
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (s *Storage) SignUp(email string, password string) error {
	userExists, err := s.isUserExist(email)
	if err != nil {
		return fmt.Errorf("failed to check user: %w", err)
	}

	if userExists {
		var NewUser models.User

		err := s.db.Where("email = ?", email).First(&NewUser).Error
		if err != nil {
			return err
		}

		NewUser.Password = password
		err = s.db.Save(&NewUser).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) SignIn(email string, password string) error {
	var existingUser models.User
	err := s.db.Where("email = ?", email).First(&existingUser).Error
	if err != nil {
		return err
	}
	if existingUser.Password != password {
		return fmt.Errorf("invalid password")
	}
	return nil
}

func (s *Storage) CheckCode(email string, code string) error {
	var foundCode models.CodeForEmail
	err := s.db.Where("email = ? AND code = ?", email, code).First(&foundCode).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("invalid code")
		}
		return err
	}
	err = s.db.Delete(&foundCode).Error
	if err != nil {
		return err
	}

	var User = models.User{
		Email:     email,
		CreatedAt: time.Now().Format(time.RFC3339),
	}

	err = s.db.Create(&User).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) SendCodeAgain(cfg config.Config, email string) error {
	var foundCode models.CodeForEmail
	err := s.db.Where("email = ?", email).First(&foundCode).Error
	if err != nil {
		return err
	}
	s.db.Delete(&foundCode)
	code := services.GenerateCodeForEmail()

	err = services.SendCodeToEmailService(cfg, code, email)
	if err != nil {
		return err
	}

	if err := s.db.Create(&models.CodeForEmail{Email: email, Code: code, CreatedAt: time.Now().Format(time.RFC3339)}).Error; err != nil {
		return err
	}

	return nil
}

func (s *Storage) CreateNews(name, description, category string, price float32) error {
	var foundNews models.News
	err := s.db.Where("name = ?", name).First(&foundNews).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newNews := models.News{Name: name, Description: description, Price: price, Category: models.NewsCategory(category), CreatedAt: time.Now().Format(time.RFC3339)}
			err := s.db.Create(&newNews).Error
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return fmt.Errorf("news already exists")
}

func (s *Storage) GetNews(category string) ([]models.News, error) {
	var AllNews []models.News
	fmt.Println(category)
	switch category {
	case string(models.Covid):
		err := s.db.Where("category = ?", category).Find(&AllNews).Error
		if err != nil {
			return nil, err
		}
		return AllNews, nil
	case string(models.Popular):
		err := s.db.Where("category = ?", category).Find(&AllNews).Error
		if err != nil {
			return nil, err
		}
		return AllNews, nil
	case string(models.Complex):
		err := s.db.Where("category = ?", category).Find(&AllNews).Error
		if err != nil {
			return nil, err
		}
		return AllNews, nil
	case string(models.Default):
		err := s.db.Find(&AllNews).Error
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("we don't have this category")
	}

	return AllNews, nil
}

func (s *Storage) DeleteNews(id uint32) error {
	err := s.db.Where("id = ?", id).Delete(&models.News{}).Error
	if err != nil {
		return err
	}
	return nil
}
func (s *Storage) FillProfile(email, firstName, lastName, middleName, birthDate string, sex models.SexType) error {
	var foundUser models.User

	err := s.db.Where("email = ?", email).First(&foundUser).Error
	if err != nil {
		return err
	}
	foundUser.BirthDate = birthDate
	foundUser.FirstName = firstName
	foundUser.LastName = lastName
	foundUser.MiddleName = middleName
	foundUser.Sex = sex

	err = s.db.Save(&foundUser).Error
	if err != nil {
		return err
	}
	return nil
}
