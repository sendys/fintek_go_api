package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateProduct creates a new product with optional image upload
func CreateProduct(c *gin.Context) {
	var request models.ProductCreateRequest

	// Parse JSON data
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Create product instance
	product := models.Product{
		Uuid:        uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		Category:    request.Category,
		Brand:       request.Brand,
		SKU:         request.SKU,
		Status:      "active",
	}

	if request.Status != "" {
		product.Status = request.Status
	}

	// Get user ID from context (set by JWT middleware)
	if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(uuid.UUID); ok {
			product.CreatedBy = uid
		}
	}

	// Save product to database
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create product",
			"details": err.Error(),
		})
		return
	}

	// Convert to response format
	response := models.ProductResponse{
		ID:          product.Uuid,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Brand:       product.Brand,
		SKU:         product.SKU,
		ImageURL:    product.ImageURL,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    response,
	})
}

// UploadProductImage uploads image for a specific product
func UploadProductImage(c *gin.Context) {
	productID := c.Param("id")

	// Parse UUID
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	// Find product
	var product models.Product
	if err := config.DB.Where("uuid = ?", productUUID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Get uploaded file
	fileHeader, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No image file provided",
		})
		return
	}

	// Save uploaded image
	imageResponse, err := utils.SaveUploadedImage(c, fileHeader)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Delete old image if exists
	if product.ImagePath != "" {
		utils.DeleteImage(product.ImagePath)
	}

	// Update product with new image info
	product.ImagePath = imageResponse.ImagePath
	product.ImageURL = imageResponse.ImageURL

	if err := config.DB.Save(&product).Error; err != nil {
		// If database update fails, delete the uploaded file
		utils.DeleteImage(imageResponse.ImagePath)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update product with image info",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Image uploaded successfully",
		"image_url": imageResponse.ImageURL,
	})
}

// GetAllProducts retrieves all products with pagination and filtering
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	var total int64

	// Get query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	category := c.Query("category")
	status := c.DefaultQuery("status", "active")
	search := c.Query("search")

	// Calculate offset
	offset := (page - 1) * limit

	// Build query
	query := config.DB.Model(&models.Product{})

	// Apply filters
	if category != "" {
		query = query.Where("category = ?", category)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if search != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	// Get total count
	query.Count(&total)

	// Get products with pagination
	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	// Convert to response format
	var responses []models.ProductResponse
	for _, product := range products {
		responses = append(responses, models.ProductResponse{
			ID:          product.Uuid,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			Brand:       product.Brand,
			SKU:         product.SKU,
			ImageURL:    product.ImageURL,
			Status:      product.Status,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"data":    responses,
		"pagination": gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + int64(limit) - 1) / int64(limit),
		},
	})
}

// GetProductByID retrieves a single product by ID
func GetProductByID(c *gin.Context) {
	productID := c.Param("id")

	// Parse UUID
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	var product models.Product
	if err := config.DB.Where("uuid = ?", productUUID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Convert to response format
	response := models.ProductResponse{
		ID:          product.Uuid,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Brand:       product.Brand,
		SKU:         product.SKU,
		ImageURL:    product.ImageURL,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved successfully",
		"data":    response,
	})
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	productID := c.Param("id")

	// Parse UUID
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	// Find existing product
	var product models.Product
	if err := config.DB.Where("uuid = ?", productUUID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	var request models.ProductUpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request data",
			"details": err.Error(),
		})
		return
	}

	// Update fields if provided
	if request.Name != "" {
		product.Name = request.Name
	}
	if request.Description != "" {
		product.Description = request.Description
	}
	if request.Price != nil {
		product.Price = *request.Price
	}
	if request.Stock != nil {
		product.Stock = *request.Stock
	}
	if request.Category != "" {
		product.Category = request.Category
	}
	if request.Brand != "" {
		product.Brand = request.Brand
	}
	if request.SKU != "" {
		product.SKU = request.SKU
	}
	if request.Status != "" {
		product.Status = request.Status
	}

	// Set updated by user
	if userID, exists := c.Get("user_id"); exists {
		if uid, ok := userID.(uuid.UUID); ok {
			product.UpdatedBy = uid
		}
	}

	// Save changes
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update product",
			"details": err.Error(),
		})
		return
	}

	// Convert to response format
	response := models.ProductResponse{
		ID:          product.Uuid,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		Brand:       product.Brand,
		SKU:         product.SKU,
		ImageURL:    product.ImageURL,
		Status:      product.Status,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data":    response,
	})
}

// DeleteProduct soft deletes a product
func DeleteProduct(c *gin.Context) {
	productID := c.Param("id")

	// Parse UUID
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	var product models.Product
	if err := config.DB.Where("uuid = ?", productUUID).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Soft delete the product
	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete product",
		})
		return
	}

	// Optionally delete the image file
	if product.ImagePath != "" {
		utils.DeleteImage(product.ImagePath)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}

// GetProductCategories retrieves all unique product categories
func GetProductCategories(c *gin.Context) {
	var categories []string

	if err := config.DB.Model(&models.Product{}).
		Distinct("category").
		Where("category != ''").
		Pluck("category", &categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve categories",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Categories retrieved successfully",
		"data":    categories,
	})
}
