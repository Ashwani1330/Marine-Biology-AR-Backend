package controllers

import (
	"encoding/json"
	"marine-ar-backend/models"
	"net/http"
	"os"
	"strconv"
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

func GetModelByID(c *gin.Context) {
    idParam := c.Query("id")
    if idParam == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Model ID not provided"})
        return
    }

    // Convert string to int
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid model ID"})
        return
    }

    file, err := os.ReadFile("data/dummy_data.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read model data"})
        return
    }

    var modelsList []models.ModelMetadata
    if err := json.Unmarshal(file, &modelsList); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not parse model data"})
        return
    }

    for _, m := range modelsList {
        if m.ID == id {
            c.JSON(http.StatusOK, m)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
}
