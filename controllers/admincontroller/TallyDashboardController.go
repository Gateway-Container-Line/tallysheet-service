package admincontroller

import (
	"github.com/Gateway-Container-Line/tallysheet-service/helper"
	"github.com/Gateway-Container-Line/tallysheet-service/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func CountTallySheet() (int64, error) {
	//func CountTallySheet(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			//helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return 0, err
		default:
			//helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return 0, err
		}
	}
	//helper.ResponseJSON(w, http.StatusOK, count)
	return count, nil
}

func CountCargoIn() (int64, error) {
	//func CountCargoIn(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%In%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			//helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return 0, err
		default:
			//helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return 0, err
		}
	}
	//helper.ResponseJSON(w, http.StatusOK, count)
	return count, nil
}

func CountCargoOut() (int64, error) {
	//func CountCargoOut(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Out%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			//helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return 0, err
		default:
			//helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return 0, err
		}
	}
	//helper.ResponseJSON(w, http.StatusOK, count)
	return count, nil
}

func CargoCoload() (int64, error) {
	//func CargoCoload(w http.ResponseWriter, r *http.Request) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Coload%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			//helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return 0, err
		default:
			//helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return 0, err
		}
	}
	//helper.ResponseJSON(w, http.StatusOK, count)
	return count, nil
}

func CargoDamaged() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Damaged%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return 0, err
		default:
			return 0, err
		}
	}
	return count, nil
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

type CountRoleUser struct {
	CountTotalTallysheet   int64
	CountUnloadedStuffing  int
	CountCargoIn           int64
	CountCargoDamaged      int64
	CountCargoOut          int64
	CountCargoTemporaryOut int64
	CountCargoColoaded     int64
	CargoShort             int64
	CargoOver              int64
	CargoCanceled          int64
	CountCargoStuffing     int64
}

func CountAllContent(w http.ResponseWriter, r *http.Request) {
	//query := r.URL.Query()
	//role := query.Get("Role")

	logrus.Info("Getting Count API ...")

	TallysheetCountResult, errCountAll := CountTallySheet()
	if errCountAll != nil {
		switch errCountAll {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errCountAll.Error())
			return
		}
	}

	TallysheetCargoInCountResult, errCargoIn := CountCargoIn()
	if errCargoIn != nil {
		switch errCargoIn {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errCargoIn.Error())
			return
		}
	}

	TallysheetCargoOutCountResult, errCargoOut := CountCargoOut()
	if errCargoOut != nil {
		switch errCargoOut {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errCargoOut.Error())
			return
		}
	}

	TallysheetColoadCountResult, errCoload := CargoCoload()
	if errCoload != nil {
		switch errCoload {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errCoload.Error())
			return
		}
	}

	TallysheetDamagedCountResult, errDamaged := CargoDamaged()
	if errDamaged != nil {
		switch errDamaged {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errDamaged.Error())
			return
		}
	}

	var countRoleUserResult CountRoleUser
	countRoleUserResult.CountTotalTallysheet = TallysheetCountResult
	countRoleUserResult.CountCargoIn = TallysheetCargoInCountResult
	countRoleUserResult.CountCargoOut = TallysheetCargoOutCountResult
	countRoleUserResult.CountCargoColoaded = TallysheetColoadCountResult
	countRoleUserResult.CountCargoDamaged = TallysheetDamagedCountResult
	helper.ResponseJSON(w, http.StatusOK, countRoleUserResult)
}
