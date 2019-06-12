package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/loomnetwork/plasma-indexer/models"
)

const (
	maxTxPerPage    = 20
	maxStatsPerPage = 100
	TimeLayout      = "2006-01-02T15:04:05"
)

type PlasmaController struct {
	DB *gorm.DB
}

type ListLoomStoreEventsResponse struct {
	Data  []models.NewValueSet `json:"data"`
	Total int                  `json:"total"`
	Page  int                  `json:"page"`
	Limit int                  `json:"limit"`
}

// ListLoomStoreEvents returns events emitted from LoomStore contract
func (c *PlasmaController) ListLoomStoreEvents(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	pagen := 1
	page := r.URL.Query().Get("page")
	if page != "" {
		pagen, _ = strconv.Atoi(page)
		if pagen <= 0 {
			pagen = 1
		}
	}
	perPagen := maxStatsPerPage
	perPage := r.URL.Query().Get("per_page")
	if perPage != "" {
		perPagen, _ = strconv.Atoi(perPage)
		if perPagen <= 0 {
			perPagen = maxStatsPerPage
		}
	}

	var data []models.NewValueSet
	result := c.DB.
		Where(models.NewValueSet{Name: name}).
		Limit(perPagen).
		Offset((pagen - 1) * perPagen).
		Order("created_at DESC")

	err := result.Find(&data).Error
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var countTotal int
	c.DB.Model(models.NewValueSet{}).Count(&countTotal)

	resp := ListLoomStoreEventsResponse{
		Data:  data,
		Page:  pagen,
		Limit: perPagen,
		Total: countTotal,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&resp)
}
