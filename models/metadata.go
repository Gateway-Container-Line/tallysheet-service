package models

import (
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
)

type MetaData struct {
	Token string `json:"token,omitempty"`
	Limit int    `json:"limit,omitempty;query:limit"`
	Page  int    `json:"page,omitempty;query:page"`
	//Sort       string `json:"sort,omitempty;query:sort"`
	TotalPages int   `json:"totalPages,omitempty"`
	TotalRows  int64 `json:"totalRows,omitempty"`
}

//
//func newPaginate(limit int, page int) *MetaData {
//	return&
//}

func (meta *MetaData) GetOffset(r *http.Request) int {
	return (meta.GetPage(r) - 1) * meta.GetLimit(r)
}

func (meta *MetaData) GetPage(r *http.Request) int {
	query := r.URL.Query()
	page, _ := strconv.Atoi(query.Get("page"))
	meta.Page = page
	if meta.Page <= 0 {
		meta.Page = 1
	}
	return meta.Page
}

func (meta *MetaData) GetLimit(r *http.Request) int {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	meta.Limit = limit
	switch {
	case meta.Limit > 100:
		meta.Limit = 100
	case meta.Limit <= 0:
		meta.Limit = 10
	}
	return meta.Limit
}

func (meta *MetaData) GetTotalRows(db *gorm.DB, model any) int64 {
	var totalRows int64
	db.Model(model).Count(&totalRows)
	return totalRows
}

func (meta *MetaData) GetTotalPages(db *gorm.DB, model any, r *http.Request) int {
	totalPages := int(math.Ceil(float64(meta.GetTotalRows(db, model)) / float64(meta.GetLimit(r))))
	return totalPages
}

//func (meta *MetaData) GetSort(r *) string {
//	query :=
//	if meta.Sort == "" {
//		meta.Sort = "Id desc"
//	}
//	return meta.Sort
//}

//	func paginate(r *http.Request, model any) func(db *gorm.DB) (*gorm.DB, *MetaData) {
//		return func(db *gorm.DB) (*gorm.DB, *MetaData) {
//			query := r.URL.Query()
//			page, _ := strconv.Atoi(query.Get("page"))
//			if page <= 0 {
//				page = 1
//			}
//
//			limit, _ := strconv.Atoi(query.Get("limit"))
//			switch {
//			case limit > 100:
//				limit = 100
//			case limit <= 0:
//				limit = 10
//			}
//
//			offset := (page - 1) * limit
//
//			var totalRows int64
//			db.Model(model).Count(&totalRows)
//			totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))
//
//			var Meta *MetaData
//			Meta.Page = page
//			Meta.Limit = limit
//			Meta.TotalRows = totalRows
//			Meta.TotalPages = totalPages
//
//			return db.Offset(offset).Limit(limit), Meta
//		}
//	}
func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := r.URL.Query()
		page, _ := strconv.Atoi(query.Get("page"))
		if page <= 0 {
			page = 1
		}

		limit, _ := strconv.Atoi(query.Get("limit"))
		switch {
		case limit > 100:
			limit = 100
		case limit <= 0:
			limit = 10
		}

		offset := (page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}
