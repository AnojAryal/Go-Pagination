package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/robbyklein/pages/helpers"
	"github.com/robbyklein/pages/initializers"
	"github.com/robbyklein/pages/models"
)

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

    pagination, err := helpers.GetPaginationData(page, perPage, models.Person{}, "/people")
    if err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to count people"})
        return
    }

    // Get people for the current page
    var people []models.Person
    result := initializers.DB.Limit(perPage).Offset(pagination.Offset).Find(&people)
    if result.Error != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch people"})
        return
    }

    // Render the page
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "people":     people,
        "pagination": pagination,
    })
}
