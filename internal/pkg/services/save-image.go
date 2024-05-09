package services

import (
	"fmt"
	"github.com/google/uuid"
	"oat-lab-module/internal/utils/config"
	"os"
	"path/filepath"
)

func SaveImage(cfg config.Config, picture []byte) (string, error) {
	name, _ := uuid.NewRandom()
	filename := fmt.Sprintf("%s.png", name)
	directory := "./media"

	filePath := filepath.Join(directory, filename)

	if err := os.MkdirAll(directory, 0755); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err := file.Write(picture); err != nil {
		return "", err
	}

	var link string
	if picture != nil {
		link = cfg.ServerNgrokHost + cfg.ServerPort + "/media/" + filename
	}
	return link, nil
}
