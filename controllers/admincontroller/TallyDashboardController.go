package admincontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func CountTallySheet(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}

func CountCargoIn(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%In%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}

func CountCargoOut(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Out%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}

func CargoCoload(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Coload%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}

func CountCargoInRack(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Rack%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}

func CountCargoLoadedInContainer(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Loaded%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	helper.ResponseJSON(w, http.StatusOK, count)
}
