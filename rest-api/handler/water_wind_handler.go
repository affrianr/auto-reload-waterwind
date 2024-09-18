package handler

import (
	"net/http"
	"strconv"

	"github.com/affrianr/auto-reload-waterwind/rest-api/dto"
	"github.com/affrianr/auto-reload-waterwind/rest-api/models"
	"github.com/affrianr/auto-reload-waterwind/rest-api/usecase"
	"github.com/gin-gonic/gin"
)

type WaterWindHandler struct {
	WaterWindUsecase usecase.WaterWindUsecase
}

func NewWaterWindHandler(uc usecase.WaterWindUsecase) *WaterWindHandler {
	return &WaterWindHandler{WaterWindUsecase: uc}
}

func (h *WaterWindHandler) GetAll(c *gin.Context) {
	data, err := h.WaterWindUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a slice of response DTOs
	var response []dto.WaterWindResponse
	for _, record := range data {
		res := dto.WaterWindResponse{
			ID:    record.ID,
			Water: record.Water,
			Wind:  record.Wind,
		}

		// Add status based on water and wind levels

		response = append(response, res)
	}

	c.JSON(http.StatusOK, response)
}

func (h *WaterWindHandler) Create(c *gin.Context) {
	var req dto.CreateWaterWindRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	waterWindModel := &models.WaterWind{
		Water: req.Water,
		Wind:  req.Wind,
	}

	newData, err := h.WaterWindUsecase.Create(waterWindModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "WaterWind data created successfully",
		"data":    newData,
	})
}

func (h *WaterWindHandler) Update(c *gin.Context) {
	var req dto.UpdateWaterWindRequest

	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	waterWindModel := &models.WaterWind{
		Water: req.Water,
		Wind:  req.Wind,
	}

	_, err := h.WaterWindUsecase.Update(uint(id), waterWindModel)

	response := dto.WaterWindResponse{
		ID:    uint(id),
		Water: req.Water,
		Wind:  req.Wind,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *WaterWindHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.WaterWindUsecase.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data deleted successfully"})
}
