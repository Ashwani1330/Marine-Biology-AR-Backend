package controllers

import (
	"encoding/json"
	"marine-ar-backend/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetModels(c *gin.Context) {
	file, err := os.ReadFile("data/dummy_data.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read model data"})
		return
	}

	var modelsList []models.ModelMetadata
	if err := json.Unmarshal(file, &modelsList); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read model data"})
		return;
	}

	filter := c.Query("type")
	if filter != "" {
		var filtered []models.ModelMetadata
		for _, m := range modelsList {
			if strings.EqualFold(m.Category, filter) {
				filtered = append(filtered, m);
			}
		}
		modelsList = filtered;
	}

	c.JSON(http.StatusOK, modelsList)
}
