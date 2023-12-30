package core

import (
	"math"
)

const PagingLimitDefault = 20

type Paging struct {
	List  interface{} `json:"list"`
	Page  int64       `json:"page"`
	Limit int64       `json:"limit"`
	Count int64       `json:"count"`
	Total int64       `json:"total"`
	Start int64       `json:"start"`
	End   int64       `json:"end"`
}

func Offset(pageNo int64, limitNo int64) int64 {
	return (limitNo * pageNo) - limitNo
}

func Pagination(pageNo int64, limitNo int64, getCount func() int64, getData func(limit int64, offset int64) interface{}) Paging {
	var pageCount = math.Ceil(float64(getCount()) / float64(limitNo))
	pageCountInt := int64(pageCount)
	if pageNo <= 0 {
		pageNo = 1
	}
	offset := (limitNo * pageNo) - limitNo
	var startRow = (pageNo - 1) * limitNo
	var endRow = startRow + limitNo - 1

	data := getData(limitNo, offset)

	return Paging{
		List:  data,
		Page:  pageNo,
		Limit: limitNo,
		Count: pageCountInt,
		Total: getCount(),
		Start: startRow,
		End:   endRow,
	}
}
