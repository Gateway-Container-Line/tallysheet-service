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
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%CargoOut%'").Find(&tallysheet).Count(&count).Error; err != nil {
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

func CargoTemporaryOut() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Temporary%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return 0, err
		default:
			return 0, err
		}
	}
	return count, nil
}

func CargoShort() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Short%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return 0, err
		default:
			return 0, err
		}
	}
	return count, nil
}

func CargoOver() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Over%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return 0, err
		default:
			return 0, err
		}
	}
	return count, nil
}

func CargoCanceled() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Cancel%'").Find(&tallysheet).Count(&count).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return 0, err
		default:
			return 0, err
		}
	}
	return count, nil
}

func CountCargoLoadedInContainer() (int64, error) {
	var count int64
	var tallysheet []models.TallySheet
	if err := models.DB.Preload(clause.Associations).Where("`status_tally` LIKE '%Loaded%'").Find(&tallysheet).Count(&count).Error; err != nil {
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

type CountOutput struct {
	CountCargo struct {
		CountTotalTallysheet   int64
		CountCargoIn           int64
		CountCargoDamaged      int64
		CountCargoOut          int64
		CountCargoTemporaryOut int64
		CountCargoColoaded     int64
		CargoShort             int64
		CargoOver              int64
		CargoCanceled          int64
		CountLoadedCargo       int64
	}
	CountStuffing struct {
		CountUnloadedStuffing int64
		CountLoadedStuffing   int64
	}
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

	TallysheetTemporaryOutCountResult, errTemporaryOut := CargoTemporaryOut()
	if errTemporaryOut != nil {
		switch errTemporaryOut {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errTemporaryOut.Error())
		}
	}

	TallysheetShortCountResult, errShort := CargoShort()
	if errShort != nil {
		switch errShort {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errTemporaryOut.Error())
		}
	}

	TallysheetOverCountResult, errOver := CargoOver()
	if errOver != nil {
		switch errOver {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errTemporaryOut.Error())
		}
	}

	TallysheetCanceledCountResult, errCanceled := CargoCanceled()
	if errCanceled != nil {
		switch errCanceled {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errTemporaryOut.Error())
		}
	}

	TallysheetLoadedCountResult, errLoaded := CountCargoLoadedInContainer()
	if errLoaded != nil {
		switch errLoaded {
		case gorm.ErrRecordNotFound:
			helper.ResponseError(w, http.StatusNotFound, "There is no Tallysheet!")
			return
		default:
			helper.ResponseError(w, http.StatusInternalServerError, errTemporaryOut.Error())
		}
	}

	var countOutputResult CountOutput
	countOutputResult.CountCargo.CountTotalTallysheet = TallysheetCountResult
	countOutputResult.CountCargo.CountCargoIn = TallysheetCargoInCountResult
	countOutputResult.CountCargo.CountCargoOut = TallysheetCargoOutCountResult
	countOutputResult.CountCargo.CountCargoColoaded = TallysheetColoadCountResult
	countOutputResult.CountCargo.CountCargoDamaged = TallysheetDamagedCountResult
	countOutputResult.CountCargo.CountCargoTemporaryOut = TallysheetTemporaryOutCountResult
	countOutputResult.CountCargo.CargoShort = TallysheetShortCountResult
	countOutputResult.CountCargo.CargoOver = TallysheetOverCountResult
	countOutputResult.CountCargo.CargoCanceled = TallysheetCanceledCountResult
	countOutputResult.CountCargo.CountLoadedCargo = TallysheetLoadedCountResult

	helper.ResponseJSON(w, http.StatusOK, countOutputResult)
}
