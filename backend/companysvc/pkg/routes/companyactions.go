package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msalemor/contoso-crm-common/pkg/models"
)

func getAll(c *gin.Context) {
	var companies []models.Company
	db.Find(&companies)
	c.JSON(http.StatusOK, companies)
}

func getByID(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func createContact(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&company)
	c.JSON(http.StatusCreated, company)
}

func update(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&company)
	c.JSON(http.StatusOK, company)
}

func patch(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&company).Updates(company)
	c.JSON(http.StatusOK, company)
}

func delete(c *gin.Context) {
	var company models.Company
	id := c.Param("id")

	if err := db.First(&company, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	db.Delete(&company)
	c.Status(http.StatusNoContent)
}
