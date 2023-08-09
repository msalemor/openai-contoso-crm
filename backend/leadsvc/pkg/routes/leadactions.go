package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msalemor/contoso-crm-common/pkg/models"
	"github.com/msalemor/contoso-crm-common/pkg/utils"
)

const (
	SB_QUEUE = "lead-queue"
	MODEL    = "lead"
)

func getAll(c *gin.Context) {
	var leads []models.Lead
	db.Find(&leads)
	c.JSON(http.StatusOK, leads)
}

func getByID(c *gin.Context) {
	var leads []models.Lead
	id := c.Param("id")

	if err := db.First(&leads, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, leads)
}

func create(c *gin.Context) {
	var lead models.Lead
	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&lead)

	utils.SendEvent(SB_QUEUE, "create", MODEL, lead, sbClient)

	c.JSON(http.StatusCreated, lead)
}

func update(c *gin.Context) {
	var lead models.Lead
	id := c.Param("id")

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&lead)

	utils.SendEvent(SB_QUEUE, "update", MODEL, lead, sbClient)

	c.JSON(http.StatusOK, lead)
}

func patch(c *gin.Context) {
	var lead models.Lead
	id := c.Param("id")

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&lead); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&lead).Updates(lead)

	utils.SendEvent(SB_QUEUE, "patch", MODEL, lead, sbClient)

	c.JSON(http.StatusOK, lead)
}

func delete(c *gin.Context) {
	var lead models.Lead
	id := c.Param("id")

	if err := db.First(&lead, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	db.Delete(&lead)

	utils.SendEvent(SB_QUEUE, "delete", MODEL, lead, sbClient)

	c.Status(http.StatusNoContent)
}
