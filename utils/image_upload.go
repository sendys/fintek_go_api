package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	MaxFileSize = 10 << 20 // 10MB
	UploadDir   = "uploads/products"
)

type ImageUploadResponse struct {
	ImagePath string `json:"image_path"`
	ImageURL  string `json:"image_url"`
}

// InitUploadDir creates upload directory if it doesn't exist
func InitUploadDir() error {
	if _, err := os.Stat(UploadDir); os.IsNotExist(err) {
		return os.MkdirAll(UploadDir, 0755)
	}
	return nil
}

// ValidateImageFile validates uploaded image file
func ValidateImageFile(fileHeader *multipart.FileHeader) error {
	// Check file size
	if fileHeader.Size > MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of 10MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".webp": true,
	}

	if !allowedExts[ext] {
		return fmt.Errorf("invalid file type. Only JPG, JPEG, PNG, GIF, and WEBP are allowed")
	}

	return nil
}

// SaveUploadedImage saves the uploaded image and returns file path and URL
func SaveUploadedImage(c *gin.Context, fileHeader *multipart.FileHeader) (*ImageUploadResponse, error) {
	// Validate file
	if err := ValidateImageFile(fileHeader); err != nil {
		return nil, err
	}

	// Generate unique filename
	ext := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String(), time.Now().Unix(), ext)
	filePath := filepath.Join(UploadDir, filename)

	// Open uploaded file
	src, err := fileHeader.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open uploaded file: %v", err)
	}
	defer src.Close()

	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create destination file: %v", err)
	}
	defer dst.Close()

	// Copy file content
	if _, err := io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("failed to save file: %v", err)
	}

	// Generate URL (assuming server runs on localhost:8081)
	imageURL := fmt.Sprintf("http://localhost:8081/uploads/products/%s", filename)

	return &ImageUploadResponse{
		ImagePath: filePath,
		ImageURL:  imageURL,
	}, nil
}

// DeleteImage deletes an image file from the filesystem
func DeleteImage(imagePath string) error {
	if imagePath == "" {
		return nil
	}

	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return nil // File doesn't exist, nothing to delete
	}

	return os.Remove(imagePath)
}

// UpdateImage handles image update - deletes old image and saves new one
func UpdateImage(c *gin.Context, fileHeader *multipart.FileHeader, oldImagePath string) (*ImageUploadResponse, error) {
	// Save new image
	imageResponse, err := SaveUploadedImage(c, fileHeader)
	if err != nil {
		return nil, err
	}

	// Delete old image if exists
	if oldImagePath != "" {
		DeleteImage(oldImagePath)
	}

	return imageResponse, nil
}