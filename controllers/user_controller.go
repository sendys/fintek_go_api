package controllers

import (
	"backend/config"
	"backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	// Ambil user_id dari parameter (misal: /profile/:user_id)
	userIDParam := c.Param("user_id")
	var user models.User

	if userIDParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter user_id diperlukan"})
		return
	}

	// Konversi ke uint
	var userID uint
	if _, err := fmt.Sscanf(userIDParam, "%d", &userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id tidak valid"})
		return
	}

	if err := config.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func GetAllUsers(c *gin.Context) {
	// Ambil query parameter page dan pageSize, default: page=1, pageSize=10
	page := 1
	pageSize := 10
	if p := c.Query("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if ps := c.Query("pageSize"); ps != "" {
		fmt.Sscanf(ps, "%d", &pageSize)
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	var users []models.User
	offset := (page - 1) * pageSize
	if err := config.DB.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data user"})
		return
	}

	var total int64
	config.DB.Model(&models.User{}).Count(&total)

	var result []gin.H
	for _, user := range users {
		result = append(result, gin.H{
			"uuid":  user.Uuid,
			"name":  user.Name,
			"email": user.Email,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       result,
		"page":       page,
		"pageSize":   pageSize,
		"total":      total,
		"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}
