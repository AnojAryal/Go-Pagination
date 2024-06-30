package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
)

type PaginationData struct {
	NextPage     int
	PreviousPage int
	CurrentPage  int
	TotalPages   int
}

func PeopleIndexGET(c *gin.Context) {
	// Get page number
	perPage := 10
	page := 1
	pageStr := c.Param("page")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if page < 1 {
		page = 1
	}

	// Calculate total pages
	var totalRows int64
	result := initializers.DB.Model(&models.Person{}).Count(&totalRows)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to count people"})
		return
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(perPage)))

	// Calculate offset
	offset := (page - 1) * perPage

	// Get people for the current page
	var people []models.Person
	result = initializers.DB.Limit(perPage).Offset(offset).Find(&people)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch people"})
		return
	}

	// Render the page
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"people": people,
		"pagination": PaginationData{
			NextPage:     page + 1,
			PreviousPage: page - 1,
			CurrentPage:  page,
			TotalPages:   totalPages,
		},
	})
}
