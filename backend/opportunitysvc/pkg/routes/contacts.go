package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/msalemor/contoso-crm-common/pkg/models"
)

func getAllContacts(c *gin.Context) {
	var contacts []models.Contact
	db.Find(&contacts)
	c.JSON(http.StatusOK, contacts)
}

func getContactByID(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := db.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func createContact(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Create(&contact)
	c.JSON(http.StatusCreated, contact)
}

func updateContact(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := db.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&contact)
	c.JSON(http.StatusOK, contact)
}

func patchContact(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := db.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&contact).Updates(contact)
	c.JSON(http.StatusOK, contact)
}

func deleteContact(c *gin.Context) {
	var contact models.Contact
	id := c.Param("id")

	if err := db.First(&contact, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Contact not found"})
		return
	}

	db.Delete(&contact)
	c.Status(http.StatusNoContent)
}
